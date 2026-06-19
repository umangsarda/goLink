package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/umangsarda/golink/handlers"
	"github.com/umangsarda/golink/store"
)

func main() {
	// init DynamoDB
	if err := store.InitDynamo(); err != nil {
		log.Fatal("Failed to init DynamoDB:", err)
	}
	fmt.Println("✅ DynamoDB connected")

	r := mux.NewRouter()
	r.HandleFunc("/shorten", handlers.ShortenURL).Methods("POST")
	r.HandleFunc("/{code}", handlers.RedirectURL).Methods("GET")
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "GoLink is running!")
	}).Methods("GET")

	fmt.Println("🚀 GoLink server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}