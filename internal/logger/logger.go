package logger

import (
	"log/slog"
	"os"
)

type EnvLevel string

const (
	EnvLocal EnvLevel = "local"
	EnvDev   EnvLevel = "dev"
	EnvProd  EnvLevel = "prod"
)

type LoggerConfig struct {
	Format string
	Level  slog.Level
}

var envLoggerConfigs = map[EnvLevel]LoggerConfig{
	EnvLocal: {Format: "text", Level: slog.LevelDebug},
	EnvDev:   {Format: "json", Level: slog.LevelDebug},
	EnvProd:  {Format: "json", Level: slog.LevelInfo},
}

func NewLogger(env EnvLevel) *slog.Logger {
	var log *slog.Logger

	config, exists := envLoggerConfigs[env]
	if !exists {
		config = envLoggerConfigs[EnvLocal]
	}

	var handler slog.Handler
	if config.Format == "text" {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: config.Level})
	} else {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: config.Level})
	}

	log = slog.New(handler)

	return log
}
