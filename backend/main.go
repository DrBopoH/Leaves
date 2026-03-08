// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	
	"leaves/source/database"
	"leaves/source/handlers"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		externalApiUrl := os.Getenv("EXTERNAL_API_URL")
		if externalApiUrl == "" {
			externalApiUrl = "http://localhost:5173"
		}

		w.Header().Set("Access-Control-Allow-Origin", externalApiUrl) 
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("[WW] File .env not found at /, use system properties")
	}

	database.InitDB()
	defer database.DB.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /signup", handlers.Signup(database.DB))
	mux.HandleFunc("POST /signin", handlers.Signin(database.DB))

	mux.HandleFunc("GET /me", handlers.Me())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	handler := enableCORS(mux)
	fmt.Printf("🚀 Сервер запущен на http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}