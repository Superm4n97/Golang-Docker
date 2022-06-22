package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ping"))
}

func main() {
	r := chi.NewRouter()
	r.Get("/ping", pong)
	fmt.Println("Starting server on port 8080:")
	http.ListenAndServe(":8080", r)
}
