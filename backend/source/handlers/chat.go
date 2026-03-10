// source/handlers/chat.go
package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"log"
	"os"
	"sync"
	"time"

	"leaves/source/auth"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		expectedOrigin := os.Getenv("EXTERNAL_API_URL")
		if expectedOrigin == "" {
			expectedOrigin = "http://localhost:5173"
		}
		return origin == expectedOrigin
	},
}

type ChatMessage struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Username  string    `json:"username"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

var clients = make(map[*websocket.Conn]bool)
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

		clientsMutex.Lock()
		clients[ws] = true
		clientsMutex.Unlock()

		log.Printf("[II] User '%s' connected to WS", claims.Username)

		for {
			var msgPayload struct {
				Text string `json:"text"`
			}

			err := ws.ReadJSON(&msgPayload)
			if err != nil {
				log.Printf("[II] User '%s' disconnected", claims.Username)
				clientsMutex.Lock()
				delete(clients, ws)
				clientsMutex.Unlock()
				break
			}

			res, err := db.Exec("INSERT INTO messages (user_id, content) VALUES (?, ?)", claims.UserID, msgPayload.Text)
			if err != nil {
				log.Printf("[EE] DB insert error: %v", err)
				continue
			}
			msgID, _ := res.LastInsertId()

			chatMsg := ChatMessage{
				ID:        int(msgID),
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
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("[EE] WS Write error: %v", err)
				client.Close()
				clientsMutex.RUnlock()
				clientsMutex.Lock()
				delete(clients, client)
				clientsMutex.Unlock()
				clientsMutex.RLock()
			}
		}
		clientsMutex.RUnlock()
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

		rows, err := db.Query(`
			SELECT sub.id, sub.user_id, u.username, sub.content, sub.created_at
			FROM (
				SELECT id, user_id, content, created_at
				FROM messages
				ORDER BY id DESC
				LIMIT 100
			) sub
			JOIN users u ON sub.user_id = u.id
			ORDER BY sub.id ASC
		`)
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
