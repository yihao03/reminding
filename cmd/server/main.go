package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/yihao03/reminding/internal/firebase"
	"github.com/yihao03/reminding/internal/router"
)

const (
	READ_HEADER_TIMEOUT_SEC = 5 //nolint:gosec
)

func main() {
	slog.Info("Starting server...")

	app, err := firebase.InitFirebase()
	if err != nil {
		slog.Error("Error initializing firebase", "error", err)
	}

	r := router.Setup(app)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadHeaderTimeout: READ_HEADER_TIMEOUT_SEC * time.Second,
	}

	slog.Info("Listening on :8080")
	if err := server.ListenAndServe(); err != nil {
		slog.Error("Server failed to start: %v", "error", err)
		panic(err)
	}
}
