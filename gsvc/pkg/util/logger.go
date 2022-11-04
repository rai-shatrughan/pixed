package util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is wrapper for zap
type Logger struct {
	*zap.Logger
	err error
}

// New returns a instance of logger
func (l *Logger) New() {
	// this or zap.NewProduction()
	config := zap.NewProductionConfig()
	// config := zap.NewDevelopmentConfig()
	config.Level.Enabled(zap.DebugLevel)
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	l.Logger, l.err = config.Build()

	// l.Logger, l.err = zap.NewProduction()

	if l.err != nil {
		l.Error("Can't initialize zap logger: %v", zap.Error(l.err))
	}
	defer l.Logger.Sync()
}
