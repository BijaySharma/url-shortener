package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/BijaySharma/url-shortener/write-service/internal/service"
)

type URLServiceHandlers struct {
	svc service.URLWriteService
}

func NewURLServiceHandlers(svc service.URLWriteService) *URLServiceHandlers {
	return &URLServiceHandlers{svc: svc}
}

func (h *URLServiceHandlers) SaveURLHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the original URL.
	var req struct {
		OriginalURL    string     `json:"url"`
		CustomAlias    string     `json:"custom_alias,omitempty"`
		ExpirationTime *time.Time `json:"expiration_time,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call the service to save the URL.
	shortCode, err := h.svc.SaveURL(r.Context(), req.OriginalURL)
	if err != nil {
		http.Error(w, "Failed to save URL", http.StatusInternalServerError)
		return
	}

	// Respond with the generated short URL.
	resp := struct {
		ShortCode string `json:"short_url"`
	}{
		ShortCode: shortCode,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("encode save url response: %v", err)
	}
}

// This handler is not required but is added just for the sake of completeness. It allows retrieving the original URL from a short URL.
func (h *URLServiceHandlers) GetOriginalURLHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the short URL from the query parameters.
	shortCode := r.URL.Query().Get("short_code")
	if shortCode == "" {
		http.Error(w, "Missing short_code query param", http.StatusBadRequest)
		return
	}

	// Call the service to get the original URL.
	originalURL, err := h.svc.GetOriginalURL(r.Context(), shortCode)
	if err != nil {
		http.Error(w, "Failed to retrieve original URL", http.StatusInternalServerError)
		return
	}

	// Respond with the original URL.
	resp := struct {
		OriginalURL string `json:"original_url"`
	}{
		OriginalURL: originalURL,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("encode get original url response: %v", err)
	}
}
