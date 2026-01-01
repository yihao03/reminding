// Package router sets up the HTTP router with middleware and routes.
package router

import (
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/yihao03/reminding/internal/database/sqlc"
	appmiddleware "github.com/yihao03/reminding/internal/router/middleware"
	"github.com/yihao03/reminding/internal/router/routes"
	"github.com/yihao03/reminding/internal/router/routes/adminroutes"
)

func Setup(queries *sqlc.Queries, app *firebase.App) *chi.Mux {
	r := chi.NewRouter()

	SetupMiddleware(r, app)
	SetupRoutes(r, queries, app)
	SetupAdminRoutes(r, queries, app)
	return r
}

func SetupMiddleware(r *chi.Mux, app *firebase.App) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
}

func SetupRoutes(r *chi.Mux, queries *sqlc.Queries, app *firebase.App) {
	r.Route("/api", func(r chi.Router) {
		// Unprotected routes
		r.Route("/auth", routes.SetupAuthRoutes(queries, app))

		// Protected routes
		r.Route("/", func(r chi.Router) {
			r.Use(appmiddleware.GetAuthMiddleware(app))
			r.Route("/events", routes.SetupEventRoutes(queries, app))
		})
	})
}

func SetupAdminRoutes(r chi.Router, queries *sqlc.Queries, app *firebase.App) {
	r.Route("/api/admin", func(r chi.Router) {
		// Unprotected routes
		r.Route("/auth", adminroutes.SetupAuthRoutes(queries, app))

		// Protected routes
		r.Route("/", func(r chi.Router) {
			r.Use(appmiddleware.GetAuthMiddleware(app))
			r.Route("/events", adminroutes.SetupEventRoutes(queries, app))
		})
	})
}
