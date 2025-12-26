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
	return r
}

func SetupMiddleware(r *chi.Mux, app *firebase.App) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(appmiddleware.GetAuthMiddleware(app))
}

func SetupRoutes(r *chi.Mux, queries *sqlc.Queries, app *firebase.App) {
	r.Route("/api", func(r chi.Router) {
		SetupAdminRoutes(r, queries, app)

		r.Route("/user", routes.SetupUserRoutes(queries, app))
		r.Route("/events", routes.SetupEventRoutes(queries, app))
	})
}

func SetupAdminRoutes(r chi.Router, queries *sqlc.Queries, app *firebase.App) {
	r.Route("/admin", func(r chi.Router) {
		r.Route("/events", adminroutes.SetupEventRoutes(queries, app))
	})
}
