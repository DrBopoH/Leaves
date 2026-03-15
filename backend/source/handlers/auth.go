// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// Package handlers source/handlers/auth.go
package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"fmt"
	"log"
	"regexp"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"

	"leaves/source/auth"
)

var (
	emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	// Simple In-Memory Rate Limiter for Auth endpoints
	rlMutex sync.Mutex
	rlMap   = make(map[string]int)
	rlLast  = time.Now()
)

func isRateLimited(ip string) bool {
	rlMutex.Lock()
	defer rlMutex.Unlock()

	// Reset limits every minute
	if time.Since(rlLast) > time.Minute {
		rlMap = make(map[string]int)
		rlLast = time.Now()
	}

	rlMap[ip]++
	if rlMap[ip] > 10 { // Max 10 attempts per minute per IP
		return true
	}
	return false
}

func getIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.RemoteAddr
	}
	return ip
}

type AuthRequest struct {
	Username   string `json:"username,omitempty"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	RememberMe bool   `json:"rememberMe,omitempty"`
}

type AuthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type User struct {
	ID           int
	Username     string
	PasswordHash string
}

func Signup(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if isRateLimited(getIP(r)) {
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(AuthResponse{Status: "error", Message: "Too many requests. Please try again later."})
			return
		}

		var req AuthRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(AuthResponse{Status: "error", Message: "Invalid request format"})
			return
		}

		req.Email = strings.TrimSpace(req.Email)
		req.Username = strings.TrimSpace(req.Username)

		if len(req.Password) < 8 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(AuthResponse{Status: "error", Message: "Password must be at least 8 characters"})
			return
		}
		if len(req.Username) < 3 || len(req.Username) > 30 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(AuthResponse{Status: "error", Message: "Username must be between 3 and 30 characters"})
			return
		}
		if !emailRegex.MatchString(strings.ToLower(req.Email)) || len(req.Email) > 254 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(AuthResponse{Status: "error", Message: "Invalid email format"})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("[EE] Signup: Error hashing password: %v", err)
			return
		}

		_, err = db.Exec("INSERT INTO users (username, email, password_hash) VALUES (?, ?, ?)", req.Username, req.Email, string(hashedPassword))
		if err != nil {
			log.Printf("[WW] Signup: Failed to create user %s (email: %s), error: %v", req.Username, req.Email, err)
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(AuthResponse{Status: "error", Message: "User with this email already exists"})
			return
		}

		fmt.Printf("\n[II] Sign up request successfully by \"%s (%s)\" user\n", req.Username, req.Email) // DEBUG PRINT
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(AuthResponse{Status: "success", Message: "Account created successfully!"})
	}
}

func Signin(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if isRateLimited(getIP(r)) {
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(AuthResponse{Status: "error", Message: "Too many requests. Please try again later."})
			return
		}

		var req AuthRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(AuthResponse{Status: "error", Message: "Invalid request format"})
			return
		}

		var user User

		err := db.QueryRow("SELECT id, username, password_hash FROM users WHERE email = ?", req.Email).Scan(&user.ID, &user.Username, &user.PasswordHash)
		if err != nil {
			log.Printf("[WW] Signin: User with email %s not found or DB error: %v", req.Email, err)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(AuthResponse{Status: "error", Message: "Invalid email or password"})
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
			log.Printf("[WW] Signin: Password mismatch for user %s (email: %s)", user.Username, req.Email)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(AuthResponse{Status: "error", Message: "Invalid email or password"})
			return
		}

		tokenString, expirationTime, err := auth.GenerateToken(user.ID, user.Username, req.RememberMe)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("[EE] Signin: Error generating token for user %d: %v", user.ID, err)
			return
		}

		cookie := &http.Cookie{
			Name:     "auth_token",
			Value:    tokenString,
			Expires:  expirationTime,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteNoneMode,
			Path:     "/",
		}

		cookieString := cookie.String() + "; Partitioned"
		w.Header().Add("Set-Cookie", cookieString)

		fmt.Printf("\n[II] Sign in request successfully by \"%s\" user.\n", user.Username) // DEBUG PRINT
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(AuthResponse{Status: "success", Message: "Successfully signed in!"})
	}
}

func Me() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		cookie, err := r.Cookie("auth_token")
		if err != nil {
			log.Printf("[WW] Me: Auth token cookie not found: %v", err)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(AuthResponse{Status: "error", Message: "Unauthorized"})
			return
		}
		log.Printf("[II] Me: Auth token cookie found. Value length: %d", len(cookie.Value))

		claims, err := auth.ParseToken(cookie.Value)
		if err != nil {
			log.Printf("[WW] Me: Invalid token received: %v", err)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(AuthResponse{Status: "error", Message: "Invalid token"})
			return
		}
		log.Printf("[II] Me: Token successfully parsed for UserID: %d, Username: %s", claims.UserID, claims.Username)

		json.NewEncoder(w).Encode(map[string]any{
			"status": "success",
			"user": map[string]any{
				"id":       claims.UserID,
				"username": claims.Username,
			},
		})
	}
}
