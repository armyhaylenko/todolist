package main

import (
	"fmt"
	"net/http"

	"github.com/armyhaylenko/todolist/internal/config"
	"github.com/armyhaylenko/todolist/internal/db"
	httpHandlers "github.com/armyhaylenko/todolist/internal/http"
	"github.com/armyhaylenko/todolist/internal/logging"
	"github.com/joho/godotenv"
)

func main() {
	logging.SetupLogger(true) // TODO switch between dev & prod logger
	// logger setup
	logger := logging.Logger
	logger.Info("Starting...")

	// load .env
	if err := godotenv.Load(".env"); err != nil {
		logger.Fatalw("Failed to load .env", "error", err)
	}
	cfg, err := config.FromEnv()

	if err != nil {
		logger.Fatalw("Error intializing config: %v", "error", err)
	}

	logger.Infow("Using config", "host", cfg.Host, "port", cfg.Port, "redis", cfg.RedisUrl)
	logger.Debug("Connecting to redis")
	redis := db.NewClient(cfg.RedisUrl)
	logger.Debug("Creating router")
	router := httpHandlers.RootRouter(redis)
	logger.Info("Starting server")
	if err := http.ListenAndServe(fmt.Sprintf("%v:%v", cfg.Host, cfg.Port), router); err != nil {
		logger.Fatalw("Failed to start sever", "error", err)
	}
}
