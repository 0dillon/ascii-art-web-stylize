package main

import (
	"log"
	"net/http"

	"ascii-art-web-stylize/internal/handlers"
)

func main() {
	// Declare a new serve mux for routing.
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", handlers.Home)
	mux.HandleFunc("POST /ascii-art", handlers.AsciiArt)

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	log.Print("Your server is running on port http://localhost:8080")
	// Start our server on port 8080
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
