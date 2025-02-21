package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func NewLogger(env string) (*zap.Logger, error) {

	var logger *zap.Logger
	var err error

	switch env {
	case envLocal:
		cfg := zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logger, err = cfg.Build()
	case envDev:
		cfg := zap.NewProductionConfig()
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		logger, err = cfg.Build()
	case envProd:
		config := zap.NewProductionConfig()
		config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
		logger, err = config.Build()
	}

	if err != nil {
		// В случае ошибки инициализации логгера, выводим ошибку и завершаем программу
		_, _ = os.Stderr.WriteString("Failed to initialize logger: " + err.Error() + "\n")
		os.Exit(1)
	}

	return logger, nil
}
