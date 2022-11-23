package util

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New returns a instance of logger
func NewLogger() *zap.Logger {
	// this or zap.NewProduction()
	config := zap.NewProductionConfig()
	// config := zap.NewDevelopmentConfig()
	config.Level.Enabled(zap.DebugLevel)
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := config.Build()

	// l.Logger, l.err = zap.NewProduction()

	if err != nil {
		fmt.Printf("Can't initialize zap logger: %v", zap.Error(err))
	}
	defer logger.Sync()

	return logger
}
