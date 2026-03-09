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



func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		externalApiUrl := os.Getenv("EXTERNAL_API_URL")
		if externalApiUrl == "" {
			externalApiUrl = default_externalApiUrl
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
	var err error

	err = godotenv.Load();
	if err != nil {
		fmt.Println(warn_envNotFound)		// DEBUG PRINT
		log.Println(warn_envNotFound)		// DEBUG LOG
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


	port := os.Getenv("PORT")
	if port == "" {
		port = default_internalPort
	}

	handler := enableCORS(mux)
	fmt.Printf("[II] Server ready, and lisens {\n\t\thttp://localhost:%s,\n\t\t%s,\n}\n\n", port, default_externalApiUrl)		// DEBUG PRINT
	log.Fatal(http.ListenAndServe(":"+port, handler))
}