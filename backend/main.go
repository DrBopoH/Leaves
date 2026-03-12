/*
	Leaves - A decentralized hybrid messenger.
    Copyright (C) 2026 MorangTong Creative Studio

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as published
    by the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

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

const defaultExternalApiUrl string = "http://localhost:5173"
const defaultInternalPort string = "8080"

const warnEnvNotFound string = "[WW] File .env not found at ./, using default"

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
		fmt.Println(warnEnvNotFound) // DEBUG PRINT
		log.Println(warnEnvNotFound) // DEBUG LOG
	}

	if os.Getenv("JWT_SECRET_KEY") == "" || len(os.Getenv("JWT_SECRET_KEY")) < 16 {
		log.Fatal("[FF] Fatal: JWT_SECRET_KEY is missing or too short in environment variables. Aborting.")
	}

	externalApiUrl = os.Getenv("EXTERNAL_API_URL")
	if externalApiUrl == "" {
		externalApiUrl = defaultExternalApiUrl
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultInternalPort
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
	fmt.Printf("[II] Server ready, and listens {\n\t\thttp://localhost:%s,\n\t\t%s,\n}\n\n", port, externalApiUrl) // DEBUG PRINT
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
