package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	// look up code in store
	link, exists := linkStore[code]
	if !exists {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	// increment hit counter
	link.Hits++
	linkStore[code] = link

	// redirect to original URL
	http.Redirect(w, r, link.LongURL, http.StatusMovedPermanently)
}