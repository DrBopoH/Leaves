package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
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
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Сервер работает!")
	})

	mux.HandleFunc("POST /signup", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var req AuthRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(AuthResponse{Status: "error", Message: "Неверный формат данных"})
			return
		}

		if req.Email == "" || req.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(AuthResponse{Status: "error", Message: "Email и пароль не могут быть пустыми!"})
			return
		}

		fmt.Printf("✅ Получены данные с фронта: Email=%s, Пароль=%s\n", req.Email, req.Password)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(AuthResponse{Status: "success", Message: "Бэкенд успешно прочитал твой JSON!"})
	})

	mux.HandleFunc("POST /signin", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Успешный коннект к signin!")
	})

	port := ":8080"
	fmt.Printf("🚀 Бэкенд запущен на http://localhost%s\n", port)
	
	if err := http.ListenAndServe(port, enableCORS(mux)); err != nil {
		log.Fatalf("Ошибка сервера: %v", err)
	}
}