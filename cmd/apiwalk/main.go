package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aichelnokov/apiwalk/internal/config"
	"github.com/aichelnokov/apiwalk/internal/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.MustLoad()

	r := chi.NewRouter()
  r.Use(middleware.RequestID)
  r.Use(middleware.RealIP)
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Walk " + cfg.ApiConfig.Version))
	})
	routes.Walk(r)
	
	fmt.Println("Server started at " + cfg.HTTPServer.Host + ":" + cfg.HTTPServer.Port)
	http.ListenAndServe(":" + cfg.HTTPServer.Port, r)
}