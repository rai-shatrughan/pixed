package util

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	kafka "github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// KafkaReaders wraps kafka.Reader
type kafkaReaders struct {
	conf    *viper.Viper
	logger  *zap.Logger
	Readers []*kafka.Reader
}

// KafkaWriter wraps kafka.Writer
type kafkaWriter struct {
	conf   *viper.Viper
	logger *zap.Logger
	writer *kafka.Writer
}

// New initializes new instance of readers
func NewKafkaReader(readerCount int, conf *viper.Viper, logger *zap.Logger) *kafkaReaders {
	kf := new(kafkaReaders)
	kf.conf = conf
	kf.logger = logger

	for i := 0; i < readerCount; i++ {
		reader := kafka.NewReader(kafka.ReaderConfig{
			Brokers:     conf.GetStringSlice("kafka.brokers"),
			Topic:       conf.GetString("kafka.topic"),
			GroupID:     conf.GetString("kafka.groupId"),
			MinBytes:    conf.GetInt("kafka.minBytes"),
			MaxBytes:    conf.GetInt("kafka.maxBytes"),
			StartOffset: conf.GetInt64("kafka.startOffSet"),
		})
		kf.Readers = append(kf.Readers, reader)
		kf.setupCloseReaderHandler()
	}
	return kf
}

// New initializes new instance of writer
func NewKafkaWriter(conf *viper.Viper, logger *zap.Logger) *kafkaWriter {
	kf := new(kafkaWriter)
	kf.conf = conf
	kf.logger = logger

	if kf.writer == nil {
		kf.writer = &kafka.Writer{
			Addr:     kafka.TCP(conf.GetStringSlice("kafka.brokers")...),
			Topic:    conf.GetString("kafka.topic"),
			Balancer: &kafka.RoundRobin{},
			Async:    true,
		}
	}
	kf.setupCloseWriterHandler()
	return kf
}

// Read returns kafka Message
func (kf *kafkaReaders) Read(reader *kafka.Reader) (kafka.Message, error) {
	kf.logger.Info("Reading message from group",
		zap.String("groupId", reader.Config().GroupID),
	)

	var start time.Time
	var elapsed time.Duration

	start = time.Now()

	msg, err := reader.ReadMessage(context.Background())
	if err != nil {
		kf.logger.Error("Error in Reading msg from Kafka ", zap.Error(err))
		return msg, err
	}
	elapsed = time.Since(start)

	kf.logger.Debug("Kafka Read", zap.Duration("duration", elapsed))

	return msg, nil
}

// Write pushes data to kafka
func (kf *kafkaWriter) Write(key, value []byte) error {
	start := time.Now()
	kf.logger.Info("Kafka Write", zap.String("start", start.String()))
	err := kf.writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   key,
			Value: value,
		},
	)
	elapsed := time.Since(start)
	kf.logger.Debug("Kafka Write", zap.Duration("duration", elapsed))

	if err != nil {
		return err
	} else {
		return nil
	}
}

// BatchWrite puts data into batches
func (kf *kafkaWriter) BatchWrite(key, val [][]byte) error {

	msgs := make([]kafka.Message, len(key))
	for k := range key {
		msgs[k] = kafka.Message{Key: key[k], Value: val[k]}
	}

	start := time.Now()

	err := kf.writer.WriteMessages(context.Background(), msgs...)
	elapsed := time.Since(start)
	kf.logger.Debug("Kafka Write", zap.Duration("duration", elapsed))

	if err != nil {
		kf.logger.Error("Error Writing msg to Kafka ", zap.Error(err))
		return err
	}

	return nil
}

func (kf *kafkaWriter) closeWriter() {
	kf.logger.Info("Closing writer")
	if err := kf.writer.Close(); err != nil {
		kf.logger.Error("Error in Reading msg from Kafka ", zap.Error(err))
	}
}

func (kf *kafkaReaders) closeReaders() {
	for i := range kf.Readers {
		kf.logger.Info("Closing reader : ", zap.Int("nos.", i))
		if err := kf.Readers[i].Close(); err != nil {
			kf.logger.Error("Failed to close kafka reader", zap.Error(err))
		}
	}
}

func (kf *kafkaReaders) setupCloseReaderHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		kf.logger.Info("\r- Ctrl+C pressed in Terminal")
		kf.closeReaders()
		os.Exit(0)
	}()
}

func (kf *kafkaWriter) setupCloseWriterHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		kf.logger.Info("\r- Ctrl+C pressed in Terminal")
		kf.closeWriter()
		os.Exit(0)
	}()
}
