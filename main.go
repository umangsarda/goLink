package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/umangsarda/golink/handlers"
)

func main() {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/shorten", handlers.ShortenURL).Methods("POST")
	r.HandleFunc("/{code}", handlers.RedirectURL).Methods("GET")
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "GoLink is running!")
	}).Methods("GET")

	fmt.Println("🚀 GoLink server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}