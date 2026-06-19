package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/umangsarda/golink/models"
	"github.com/umangsarda/golink/store"
)

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var body struct {
		LongURL string `json:"long_url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.LongURL == "" {
		http.Error(w, "Invalid request. Provide long_url", http.StatusBadRequest)
		return
	}

	code := uuid.New().String()[:6]

	link := models.Link{
		Code:      code,
		LongURL:   body.LongURL,
		ShortURL:  "http://localhost:8080/" + code,
		CreatedAt: time.Now().Format(time.RFC3339),
		Hits:      0,
	}

	// save to DynamoDB
	if err := store.SaveLink(link); err != nil {
		http.Error(w, "Failed to save link: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(link)
}