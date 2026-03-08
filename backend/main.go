package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./leaves.db")
	if err != nil {
		log.Fatalf("Ошибка открытия БД: %v", err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT UNIQUE NOT NULL,
		password_hash TEXT NOT NULL
	);`
	
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Ошибка создания таблицы: %v", err)
	}
	fmt.Println("📦 База данных SQLite подключена и готова!")
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	initDB()
	defer db.Close() 

	mux := http.NewServeMux()

	mux.HandleFunc("POST /signup", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var req AuthRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(AuthResponse{Status: "error", Message: "Неверный формат данных"})
			return
		}

		if req.Email == "" || req.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(AuthResponse{Status: "error", Message: "Email и пароль не могут быть пустыми!"})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(AuthResponse{Status: "error", Message: "Ошибка сервера при обработке пароля"})
			return
		}

		_, err = db.Exec("INSERT INTO users (email, password_hash) VALUES (?, ?)", req.Email, string(hashedPassword))
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(AuthResponse{Status: "error", Message: "Пользователь с таким email уже существует!"})
			return
		}

		fmt.Printf("✅ Зарегистрирован новый юзер: %s\n", req.Email)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(AuthResponse{Status: "success", Message: "Успешная регистрация!"})
	})

	mux.HandleFunc("POST /signin", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var req AuthRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(AuthResponse{Status: "error", Message: "Неверный формат данных"})
			return
		}

		var storedHash string
		err := db.QueryRow("SELECT password_hash FROM users WHERE email = ?", req.Email).Scan(&storedHash)
		if err != nil {
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(AuthResponse{Status: "error", Message: "Неверный email или пароль"})
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(req.Password))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(AuthResponse{Status: "error", Message: "Неверный email или пароль"})
			return
		}

		fmt.Printf("👋 Пользователь вошел: %s\n", req.Email)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(AuthResponse{Status: "success", Message: "Успешный вход в систему!"})
	})

	port := ":8080"
	fmt.Printf("🚀 Бэкенд запущен на http://localhost%s\n", port)
	
	if err := http.ListenAndServe(port, enableCORS(mux)); err != nil {
		log.Fatalf("Ошибка сервера: %v", err)
	}
}