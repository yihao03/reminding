package router

import (
	"net/http"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/yihao03/reminding/internal/api"
)

func Setup(app *firebase.App) *chi.Mux {
	r := chi.NewRouter()

	SetupMiddleware(r)
	SetupRoutes(r, app)
	return r
}

func SetupMiddleware(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
}

func SetupRoutes(r *chi.Mux, app *firebase.App) {
	r.Get("/", api.HTTPHandler(app,
		func(w http.ResponseWriter, r *http.Request, app *firebase.App) error {
			api.WriteResponse("Didn't forget to run", w)
			return nil
		}))
}
