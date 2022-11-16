package util

import "net/http"

type AppState struct {
	Mux         *http.ServeMux
	Conf        *Config
	Logger      *Logger
	Kfw         *KafkaWriter
	Kv          *KV
	ServiceName string
}
