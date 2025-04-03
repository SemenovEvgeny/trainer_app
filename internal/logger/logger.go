package logger

import (
	"log/slog"
	"os"
	"sync"

	"treners_app/internal/config"
)

type EnvLevel string

const (
	EnvLocal EnvLevel = "local"
	EnvDev   EnvLevel = "dev"
	EnvProd  EnvLevel = "prod"
)

type LogConfig struct {
	Format string
	Level  slog.Level
}

var envLoggerConfigs = map[EnvLevel]LogConfig{
	EnvLocal: {Format: "text", Level: slog.LevelDebug},
	EnvDev:   {Format: "json", Level: slog.LevelDebug},
	EnvProd:  {Format: "json", Level: slog.LevelInfo},
}

var (
	instance *slog.Logger
	once     sync.Once
)

func GetLogger() *slog.Logger {
	config := config.GetConfig()

	once.Do(func() {

		logConfig, exists := envLoggerConfigs[EnvLevel(config.Env)]
		if !exists {
			logConfig = envLoggerConfigs[EnvLocal]
		}

		var handler slog.Handler
		if logConfig.Format == "text" {
			handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logConfig.Level})
		} else {
			handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logConfig.Level})
		}

		instance = slog.New(handler)
	})

	return instance
}
