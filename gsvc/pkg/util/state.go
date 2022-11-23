package util

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type AppState struct {
	Conf        *viper.Viper
	Logger      *zap.Logger
	Kfw         *kafkaWriter
	Kv          *kvStrore
	ServiceName *string
}
