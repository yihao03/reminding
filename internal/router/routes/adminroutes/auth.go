package adminroutes

import (
	firebase "firebase.google.com/go/v4"
	"github.com/go-chi/chi/v5"
	"github.com/yihao03/reminding/internal/api"
	"github.com/yihao03/reminding/internal/database/sqlc"
	"github.com/yihao03/reminding/internal/handlers/user"
)

func SetupAuthRoutes(queries *sqlc.Queries, app *firebase.App) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/login", api.HTTPHandler(queries, app, user.HandleAuthorizeUser))
	}
}
