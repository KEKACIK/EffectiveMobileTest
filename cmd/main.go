package main

import (
	"TestTask/config"
	"TestTask/internal/api"
	"TestTask/pkg/logging"
	"TestTask/pkg/postgresql"
	"context"
	"net/http"
)

func main() {
	cfg := config.NewConfig()
	logger := logging.NewLogger(cfg.Debug)

	client, err := postgresql.NewClient(context.TODO(), logger, 5, cfg.GetPostgresDsn())
	if err != nil {
		logger.Fatal(err.Error())
	}

	mux := http.NewServeMux()
	handler := api.NewHandler(client, logger)

	api.RegisterRoutes(mux, handler)

	logger.Info("API listen :80")
	http.ListenAndServe(":80", mux)
}
