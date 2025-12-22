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

var (
	ErrRetrieveFirebaseClient = "Error retrieving firebase client"
	ErrUnauthorized           = "Unauthorized access"
	ErrInvalidToken           = "Invalid Firebase ID token"
)

func GetAuthMiddleware(app *firebase.App) func(http.Handler) http.Handler {
	client, err := app.Auth(context.Background())
	if err != nil {
		slog.Error("Error getting firebase auth client", "error", err)
		panic(err)
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				api.WriteError(http.StatusUnauthorized, apperrors.Wrap(nil, ErrUnauthorized+": Missing Authorization token"), w, r.Context())
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader { // Prefix wasn't found
				api.WriteError(http.StatusUnauthorized, apperrors.Wrap(nil, ErrUnauthorized), w, r.Context())
				return
			}

			token, err := client.VerifyIDToken(r.Context(), tokenString)
			if err != nil {
				api.WriteError(http.StatusUnauthorized, apperrors.Wrap(err, ErrInvalidToken), w, r.Context())
				return
			}

			// Add the UID to the context so handlers can access it
			ctx := context.WithValue(r.Context(), UserUIDKey, token.UID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserIDFromContext(ctx context.Context) string {
	return ctx.Value(UserUIDKey).(string)
}
