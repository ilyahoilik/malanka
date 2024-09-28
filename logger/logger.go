package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	logEncoding   = "json"
	LogLevel      = zapcore.InfoLevel
	LogOutputPath = "malanka.log"
)

func NewLogger() (*zap.Logger, error) {
	loggerCfg := zap.Config{
		Encoding:    logEncoding,
		Level:       zap.NewAtomicLevelAt(LogLevel),
		OutputPaths: []string{LogOutputPath},
	}
	return loggerCfg.Build()
}

func SyncLogger(logger *zap.Logger) {
	if err := logger.Sync(); err != nil {
		logger.Error("failed to sync zap logger", zap.Error(err))
		panic(fmt.Sprintf("failed to sync zap logger (%v)", err))
	}
}
