package main

import (
	"log"
	"net/http"
	"time"

	"github.com/yihao03/reminding/internal/router"
)

const (
	READ_HEADER_TIMEOUT_SEC = 5
)

func main() {
	r := router.Setup()
	server := &http.Server{
		Addr:              "8080",
		Handler:           r,
		ReadHeaderTimeout: READ_HEADER_TIMEOUT_SEC * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
