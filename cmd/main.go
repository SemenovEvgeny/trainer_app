package main

import (
	"treners_app/internal/config"
	"treners_app/internal/logger"

	"go.uber.org/zap"
)

func main() {
	//TODO: init config
	cfg := config.NewConfig()
	_ = cfg
	//TODO: init logger
	//log := logger.NewLogger(cfg.Env)

	logger := logger.NewLogger(cfg.Env)
	defer logger.Sync()

	logger.Info("This is an info message",
		zap.String("key", "value"),
		zap.Int("number", 42),
	)

	logger.Debug("This is a debug message",
		zap.String("key", "value"),
		zap.Int("number", 42),
	)

	logger.Error("This is an error message",
		zap.String("key", "value"),
		zap.Int("number", 42),
	)

	//TODO: init db

	//TODO: init middleware

	//TODO: star server
}
