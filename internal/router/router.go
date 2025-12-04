package router

import (
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func Setup() *chi.Mux {
	r := chi.NewRouter()

	SetupMiddleware(r)
	SetupRoutes(r)
	return r
}

func SetupMiddleware(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
}

func SetupRoutes(r *chi.Mux) {
}
