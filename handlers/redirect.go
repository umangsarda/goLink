package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/umangsarda/golink/store"
)

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	link, err := store.GetLink(code)
	if err != nil || link == nil {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, link.LongURL, http.StatusMovedPermanently)
}