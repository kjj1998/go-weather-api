package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

const (
	version     = "1.0.0"
	environment = "development"
)

type application struct {
	logger *slog.Logger
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", 4000),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "addr", srv.Addr, "env", "development")

	err := srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
