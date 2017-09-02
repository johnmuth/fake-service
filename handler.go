package main

import (
	"fmt"
	"net/http"
)

// Handler handles requests
type Handler struct {}

// ServeHTTP serves HTTP
func (handler Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"qux": "zzz"}`)
}
