package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	firebase "firebase.google.com/go/v4"
	"github.com/yihao03/reminding/apperrors"
	"github.com/yihao03/reminding/internal/api"
)

func GetAuthMiddleware(app *firebase.App) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			client, err := app.Auth(r.Context())
			if err != nil {
				slog.Error("Error retrieving firebase client", "error", err)
				api.WriteError(http.StatusInternalServerError, &apperrors.InternalServerError{}, w, r.Context())
				return
			}

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				api.WriteError(http.StatusUnauthorized, &apperrors.UnauthorizedError{}, w, r.Context())
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader { // Prefix wasn't found
				api.WriteError(http.StatusUnauthorized, &apperrors.UnauthorizedError{}, w, r.Context())
				return
			}

			token, err := client.VerifyIDToken(r.Context(), tokenString)
			if err != nil {
				slog.Error("Error verifying Firebase ID token", "error", err)
				api.WriteError(http.StatusUnauthorized, &apperrors.UnauthorizedError{}, w, r.Context())
				return
			}

			// Add the UID to the context so handlers can access it
			ctx := context.WithValue(r.Context(), UserIDKey, token.UID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
