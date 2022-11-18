package util

type AppState struct {
	Conf        *Config
	Logger      *Logger
	Kfw         *KafkaWriter
	Kv          *KV
	ServiceName *string
}
