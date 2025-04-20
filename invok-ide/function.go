package main

import (
	_ "embed"
	"net/http"
)

//go:embed index.html
var idePage string

// Handler for the "/invok-ide" endpoint.
func InvokIdeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(idePage))
}
