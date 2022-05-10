package main

import (
	"github.com/husanmusa/pro-book-service-api_gateway/api"
	"github.com/husanmusa/pro-book-service-api_gateway/config"

	"github.com/saidamir98/udevs_pkg/logger"
)

func main() {
	var loggerLevel string
	cfg := config.Load()

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer func() {
		if err := logger.Cleanup(log); err != nil {
			log.Error("Failed to cleanup logger", logger.Error(err))
		}
	}()

	r := api.SetupRouter(cfg, log)

	if err := r.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}
