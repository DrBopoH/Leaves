// main.go
package main

import (
	"net/http"

	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"leaves/source/database"
	"leaves/source/handlers"
)

const default_externalApiUrl string = "http://localhost:5173"
const default_internalPort string = "8080"

const warn_envNotFound string = "[WW] File .env not found at ./, using default"

var externalApiUrl string

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		if origin != externalApiUrl {
			log.Printf("[DEBUG] CORS mismatch. Request Origin: '%s', Expected: '%s'\n", origin, externalApiUrl)
			http.Error(w, "Forbidden: Invalid Origin. Only frontend is trusted.", http.StatusForbidden)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", externalApiUrl)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, ngrok-skip-browser-warning")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	var err error

	err = godotenv.Load()
	if err != nil {
		fmt.Println(warn_envNotFound) // DEBUG PRINT
		log.Println(warn_envNotFound) // DEBUG LOG
	}

	if os.Getenv("JWT_SECRET_KEY") == "" || len(os.Getenv("JWT_SECRET_KEY")) < 16 {
		log.Fatal("[FF] Fatal: JWT_SECRET_KEY is missing or too short in environment variables. Aborting.")
	}

	externalApiUrl = os.Getenv("EXTERNAL_API_URL")
	if externalApiUrl == "" {
		externalApiUrl = default_externalApiUrl
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = default_internalPort
	}

	err = database.InitDB()
	if err != nil {
		return
	}

	defer database.DB.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /signup", handlers.Signup(database.DB))
	mux.HandleFunc("POST /signin", handlers.Signin(database.DB))

	mux.HandleFunc("GET /me", handlers.Me())

	mux.HandleFunc("GET /messages", handlers.GetHistory(database.DB))
	mux.HandleFunc("/ws", handlers.HandleWebSocket(database.DB))

	go handlers.BroadcastMessages()

	handler := enableCORS(mux)
	fmt.Printf("[II] Server ready, and lisens {\n\t\thttp://localhost:%s,\n\t\t%s,\n}\n\n", port, externalApiUrl) // DEBUG PRINT
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
