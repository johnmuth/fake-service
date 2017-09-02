package main

import (
	"fmt"
	"net/http"
	"time"
	"math/rand"
)

// Handler handles requests
type Handler struct {
	MaxRandomDelayMS int
}

// ServeHTTP serves HTTP
func (handler Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	randi := rand.Intn(handler.MaxRandomDelayMS)
	time.Sleep(time.Duration(randi) * time.Millisecond)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"qux": "zzz"}`)
}
