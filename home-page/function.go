package main

import (
	_ "embed"
	"net/http"
)

//go:embed index.html
var homePage string

// Handler for the "/home-page" endpoint.
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	// You can access query params via r.URL.Query().
	// For example:
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(homePage))
}
