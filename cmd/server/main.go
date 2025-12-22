package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/yihao03/reminding/internal/database"
	"github.com/yihao03/reminding/internal/firebase"
	"github.com/yihao03/reminding/internal/router"
)

const (
	READ_HEADER_TIMEOUT_SEC = 5 //nolint:gosec
)

func main() {
	slog.Info("Starting server...")
	if err := godotenv.Load(".env"); err != nil {
		slog.Error("Error loading .env file", "error", err)
		panic(err)
	}

	app, err := firebase.InitFirebase()
	if err != nil {
		slog.Error("Error initializing firebase", "error", err)
	}

	queries, pgxPool := database.Connect()
	defer pgxPool.Close()

	r := router.Setup(queries, app)
	cors := getCorsConfig().Handler(r)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           cors,
		ReadHeaderTimeout: READ_HEADER_TIMEOUT_SEC * time.Second,
	}

	slog.Info("Listening on :8080")
	if err := server.ListenAndServe(); err != nil {
		slog.Error("Server failed to start: %v", "error", err)
		panic(err)
	}
}

func getCorsConfig() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})
}
