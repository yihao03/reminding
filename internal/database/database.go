package database

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	database "github.com/yihao03/reminding/internal/database/sqlc"
)

func Connect() (*database.Queries, *pgxpool.Pool) {
	connStr := os.Getenv("DATABASE_URL")
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		slog.Error("Failed to connect to the database", "error", err)
		panic(err)
	}

	return database.New(pool), pool
}
