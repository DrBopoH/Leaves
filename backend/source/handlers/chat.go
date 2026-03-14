// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// Package handlers source/handlers/chat.go
package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"log"
	"os"
	"strings"
	"sync"
	"time"

	"leaves/source/auth"

	"github.com/gorilla/websocket"
)

const sendMessageSQL string = `
INSERT INTO messages (user_id, content) VALUES (?, ?)
`
const getHistorySQL string = `
SELECT sub.id, sub.user_id, u.username, sub.content, sub.created_at
FROM (
	SELECT id, user_id, content, created_at
	FROM messages
	ORDER BY id DESC
	LIMIT 100
) sub
JOIN users u ON sub.user_id = u.id
ORDER BY sub.id ASC
`

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		expectedOrigins := os.Getenv("EXTERNAL_API_URL")
		if expectedOrigins == "" {
			expectedOrigins = "http://localhost:5173"
		}
		for _, o := range strings.Split(expectedOrigins, ",") {
			if origin == strings.TrimSpace(o) {
				return true
			}
		}
		return false
	},
}

type ChatMessage struct {
	ID        int       `json:"id,omitempty"`
	Type      string    `json:"type,omitempty"`
	UserID    int       `json:"user_id,omitempty"`
	Username  string    `json:"username"`
	Text      string    `json:"text,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Online    bool      `json:"online"`
	LastSeen  time.Time `json:"last_seen,omitempty"`
}

type Client struct {
	Conn     *websocket.Conn
	UserID   int
	Username string
}

var clients = make(map[*Client]bool)
var activeUsers = make(map[string]int)
var clientsMutex sync.RWMutex
var broadcast = make(chan ChatMessage)

func HandleWebSocket(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth_token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims, err := auth.ParseToken(cookie.Value)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("[EE] WS Upgrade error: %v", err)
			return
		}
		defer ws.Close()

		client := &Client{Conn: ws, UserID: claims.UserID, Username: claims.Username}

		clientsMutex.Lock()
		clients[client] = true
		activeUsers[claims.Username]++
		isFirstConnection := activeUsers[claims.Username] == 1
		clientsMutex.Unlock()

		log.Printf("[II] User '%s' connected to WS", claims.Username)

		if isFirstConnection {
			db.Exec("UPDATE users SET last_seen = CURRENT_TIMESTAMP WHERE id = ?", claims.UserID)
			broadcast <- ChatMessage{
				Type:     "status",
				Username: claims.Username,
				Online:   true,
				LastSeen: time.Now(),
			}
		}

		for {
			var msgPayload struct {
				Type string `json:"type"`
				Text string `json:"text"`
			}

			err := client.Conn.ReadJSON(&msgPayload)
			if err != nil {
				log.Printf("[II] User '%s' disconnected", claims.Username)
				clientsMutex.Lock()
				delete(clients, client)
				activeUsers[claims.Username]--
				isLastDisconnection := activeUsers[claims.Username] == 0
				clientsMutex.Unlock()

				if isLastDisconnection {
					db.Exec("UPDATE users SET last_seen = CURRENT_TIMESTAMP WHERE id = ?", claims.UserID)
					broadcast <- ChatMessage{
						Type:     "status",
						Username: claims.Username,
						Online:   false,
						LastSeen: time.Now(),
					}
				}
				break
			}

			if msgPayload.Type == "typing" {
				chatMsg := ChatMessage{
					Type:     "typing",
					Username: claims.Username,
				}
				broadcast <- chatMsg
				continue
			}
			
			msgPayload.Type = "message"

			if len(msgPayload.Text) == 0 || len(msgPayload.Text) > 2000 {
				log.Printf("[WW] User '%s' sent an invalid message length: %d", claims.Username, len(msgPayload.Text))
				continue
			}

			res, err := db.Exec(sendMessageSQL, claims.UserID, msgPayload.Text)
			if err != nil {
				log.Printf("[EE] DB insert error: %v", err)
				continue
			}
			msgID, _ := res.LastInsertId()

			chatMsg := ChatMessage{
				ID:        int(msgID),
				Type:      "message",
				UserID:    claims.UserID,
				Username:  claims.Username,
				Text:      msgPayload.Text,
				Timestamp: time.Now(),
			}
			broadcast <- chatMsg
		}
	}
}

func BroadcastMessages() {
	for {
		msg := <-broadcast

		clientsMutex.RLock()
		for client := range clients {
			err := client.Conn.WriteJSON(msg)
			if err != nil {
				log.Printf("[EE] WS Write error: %v", err)
				client.Conn.Close()
			}
		}
		clientsMutex.RUnlock()
	}
}

func GetUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		cookie, err := r.Cookie("auth_token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		_, err = auth.ParseToken(cookie.Value)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		rows, err := db.Query("SELECT DISTINCT username, last_seen FROM users ORDER BY username ASC")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		clientsMutex.RLock()
		defer clientsMutex.RUnlock()

		var users []map[string]interface{}
		for rows.Next() {
			var username string
			var lastSeenStr sql.NullString
			if err := rows.Scan(&username, &lastSeenStr); err == nil {
				var lastSeen time.Time
				if lastSeenStr.Valid {
					lastSeen, _ = time.Parse(time.RFC3339, lastSeenStr.String)
				}
				
				isOnline := activeUsers[username] > 0
				
				users = append(users, map[string]interface{}{
					"username":  username,
					"online":    isOnline,
					"last_seen": lastSeen,
				})
			}
		}

		if users == nil {
			users = []map[string]interface{}{}
		}

		json.NewEncoder(w).Encode(users)
	}
}

func GetHistory(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		cookie, err := r.Cookie("auth_token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		_, err = auth.ParseToken(cookie.Value)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		rows, err := db.Query(getHistorySQL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var msgs []ChatMessage
		for rows.Next() {
			var msg ChatMessage
			if err := rows.Scan(&msg.ID, &msg.UserID, &msg.Username, &msg.Text, &msg.Timestamp); err == nil {
				msgs = append(msgs, msg)
			}
		}

		if msgs == nil {
			msgs = []ChatMessage{}
		}

		json.NewEncoder(w).Encode(msgs)
	}
}
