// Package router sets up the HTTP router with middleware and routes.
package router

import (
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yihao03/reminding/internal/database/sqlc"
	appmiddleware "github.com/yihao03/reminding/internal/router/middleware"
	"github.com/yihao03/reminding/internal/router/routes"
	"github.com/yihao03/reminding/internal/router/routes/adminroutes"
)

func Setup(queries *sqlc.Queries, app *firebase.App, pool *pgxpool.Pool) *chi.Mux {
	r := chi.NewRouter()

	SetupMiddleware(r, app, pool)
	SetupRoutes(r, queries, app)
	SetupAdminRoutes(r, queries, app)
	return r
}

func SetupMiddleware(r *chi.Mux, app *firebase.App, pool *pgxpool.Pool) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(appmiddleware.PgxPoolMiddleware(pool))
}

func SetupRoutes(r *chi.Mux, queries *sqlc.Queries, app *firebase.App) {
	r.Route("/api", func(r chi.Router) {
		// Unprotected routes
		r.Route("/auth", routes.SetupAuthRoutes(queries, app))

		// Protected routes
		r.Route("/", func(r chi.Router) {
			r.Use(appmiddleware.GetAuthMiddleware(app))
			r.Route("/event", routes.SetupEventRoutes(queries, app))
			r.Route("/journal", routes.SetupJournalRoutes(queries, app))
			r.Route("/quote", routes.SetupQuoteRoute(queries, app))
			r.Route("/mood", routes.SetupMoodRoutes(queries, app))
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
			r.Route("/event", adminroutes.SetupEventRoutes(queries, app))
		})
	})
}
