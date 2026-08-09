package handler

import (
	"log"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write([]byte(`{"status": "healthy"}`)); err != nil {
		log.Printf("write health check response: %v", err)
	}
}
