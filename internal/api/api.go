// Package api provides HTTP handler utilities for the Reminding application.
package api

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/yihao03/reminding/internal/database/sqlc"
)

type Response struct {
	Messages []string `json:"messages"`
	Data     any      `json:"data"`
}

type Handler = func(http.ResponseWriter, *http.Request, *sqlc.Queries, *firebase.App) error

// HTTPHandler converts the internal Handler type into a standard http.HandlerFunc.
func HTTPHandler(queries *sqlc.Queries, app *firebase.App, handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if err := handler(w, r, queries, app); err != nil {
				WriteError(http.StatusInternalServerError, err, w, r.Context())
			}
		})

		http.TimeoutHandler(h, 15*time.Second, `{"messages":["request timed out"]}`).ServeHTTP(w, r)
	}
}

func WriteResponse(payload any, w http.ResponseWriter, message ...string) {
	w.Header().Set("Content-Type", "application/json")

	res := Response{
		Data: payload,
	}

	if len(message) > 0 {
		res.Messages = message
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		slog.Error("Error encoding json", "error", res)
	}
}

func WriteError(code int, err error, w http.ResponseWriter, ctx context.Context) {
	slog.ErrorContext(ctx, "Error: "+err.Error())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	res := Response{
		Messages: []string{err.Error()},
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		slog.Error("Error encoding json", "error", res)
	}
}
