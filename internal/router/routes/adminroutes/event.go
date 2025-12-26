package adminroutes

import (
	firebase "firebase.google.com/go/v4"
	"github.com/go-chi/chi/v5"
	"github.com/yihao03/reminding/internal/api"
	"github.com/yihao03/reminding/internal/database/sqlc"
	adminevents "github.com/yihao03/reminding/internal/handlers/admin/events"
)

func SetupEventRoutes(queries *sqlc.Queries, app *firebase.App) func(chi.Router) {
	return func(r chi.Router) {
		r.Get("/list", api.HTTPHandler(queries, app, adminevents.HandleGetEvents))
		r.Get("/{id}", api.HTTPHandler(queries, app, adminevents.HandleReadEvents))
	}
}
