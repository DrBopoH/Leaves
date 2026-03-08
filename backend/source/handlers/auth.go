// source/handlers/auth.go
package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"leaves/source/auth"
	"golang.org/x/crypto/bcrypt"
)

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

func Signup(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var req AuthRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(AuthResponse{Status: "error", Message: "Invalid request format"})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("INSERT INTO users (username, email, password_hash) VALUES (?, ?, ?)", req.Username, req.Email, string(hashedPassword))
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(AuthResponse{Status: "error", Message: "User with this email already exists"})
			return
		}

		fmt.Printf("Sign up request successfully by \"%s (%s)\" user\n", req.Username, req.Email)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(AuthResponse{Status: "success", Message: "Account created successfully!"})
	}
}

func Signin(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var req AuthRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(AuthResponse{Status: "error", Message: "Invalid request format"})
			return
		}

		var user struct {
			ID           int
			Username     string
			PasswordHash string
		}

		err := db.QueryRow("SELECT id, username, password_hash FROM users WHERE email = ?", req.Email).Scan(&user.ID, &user.Username, &user.PasswordHash)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(AuthResponse{Status: "error", Message: "Invalid email or password"})
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(AuthResponse{Status: "error", Message: "Invalid email or password"})
			return
		}

		tokenString, expirationTime, err := auth.GenerateToken(user.ID, user.Username, req.RememberMe)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "auth_token",
			Value:    tokenString,
			Expires:  expirationTime,
			HttpOnly: true,
			Secure:   false,
			Path:     "/",
			SameSite: http.SameSiteLaxMode,
		})

		fmt.Printf("Sign in request successfully by \"%s\" user.\n", user.Username)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(AuthResponse{Status: "success", Message: "Successfully signed in!"})
	}
}

func Me() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		
		cookie, err := r.Cookie("auth_token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(AuthResponse{Status: "error", Message: "Unauthorized"})
			return
		}

		claims, err := auth.ParseToken(cookie.Value)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(AuthResponse{Status: "error", Message: "Invalid token"})
			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "success",
			"user": map[string]interface{}{
				"id": claims.UserID,
				"username": claims.Username,
			},
		})
	}
}