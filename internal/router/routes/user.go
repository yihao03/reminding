package routes

import (
	firebase "firebase.google.com/go/v4"
	"github.com/go-chi/chi/v5"
	"github.com/yihao03/reminding/internal/api"
	"github.com/yihao03/reminding/internal/database/sqlc"
	"github.com/yihao03/reminding/internal/handlers/user"
)

func SetupUserRoutes(queries *sqlc.Queries, app *firebase.App) func(chi.Router) {
	return func(r chi.Router) {
		r.Post("/create", api.HTTPHandler(queries, app, user.CreateUser))
	}
}
