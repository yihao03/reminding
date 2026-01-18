package middleware

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func PgxPoolMiddleware(pool *pgxpool.Pool) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), DBPoolKey, pool)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetDBPoolFromContext(ctx context.Context) (*pgxpool.Pool, bool) {
	val := ctx.Value(DBPoolKey)
	if val == nil {
		return nil, false
	} else {
		pool := val.(*pgxpool.Pool)
		return pool, true
	}
}
