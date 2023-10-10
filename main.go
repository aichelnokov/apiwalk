package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aichelnokov/apiwalk/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
  r.Use(middleware.RequestID)
  r.Use(middleware.RealIP)
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	
	routes.Walk(r)
	
	fmt.Println("Server started at port 3000")
	http.ListenAndServe(":3000", r)
}