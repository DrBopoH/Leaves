package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Leaves Backend! 🍃")
	})

	mux.HandleFunc("GET /signup", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Здесь скоро будет логика регистрации")
	})

	mux.HandleFunc("GET /signin", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Здесь скоро будет логика входа")
	})

	port := ":8080"
	fmt.Printf("🚀 Сервер Leaves запущен на http://localhost%s\n", port)
	fmt.Println("Для остановки нажми Ctrl+C")

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Критическая ошибка сервера: %v", err)
	}
}