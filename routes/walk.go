package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Walk(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Walk v0.1.0"))
	})
}