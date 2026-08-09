package main

import (
	"log"
	"net/http"
	"os"
)

var listen_port = ":8080"

func init() {
	if port := os.Getenv("PORT"); port != "" {
		listen_port = ":" + port
	}
}

func main() {
	mux := http.NewServeMux()

	// register handlers
	mux.HandleFunc("/health", healthHandler)

	log.Printf("Starting server on port %s", listen_port)
	log.Fatal(http.ListenAndServe(listen_port, mux))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status": "healthy"}`))
}
