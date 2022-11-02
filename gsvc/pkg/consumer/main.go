package main

import (
	// go modules
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	// git modules
	"github.com/segmentio/kafka-go"

	// websvc modules
	md "gsvc/pkg/model"
	"gsvc/pkg/util"

	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
)

var (
	logger     util.Logger
	conf       util.Config
	client     gohbase.Client
	hbaseHost  string
	hbaseTable string
)

func init() {

	conf.New()

	hbaseHost = conf.GetString("hbase.zookeeper")
	hbaseTable = conf.GetString("kafka.topic")

	logger.New()
	client = gohbase.NewClient(hbaseHost)
}

func main() {
	createTable()
	readMessage()
}

func createTable() {

	var created bool
	var retryCount = 0

	for !created && retryCount < 10 {
		time.Sleep(3 * time.Second)
		var cFamilies = map[string]map[string]string{
			"cf": nil,
		}

		ac := gohbase.NewAdminClient(hbaseHost)
		crt := hrpc.NewCreateTable(context.Background(), []byte(hbaseTable), cFamilies)

		if err := ac.CreateTable(crt); err != nil {
			logger.Sugar().Error("createTable returned an error: %v", err)
			retryCount++
			logger.Sugar().Error("retry count : ", retryCount)
		} else {
			logger.Sugar().Info("table created in attempt : ", retryCount)
			created = true
		}

	}
}

func readMessage() {
	done := make(chan bool)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  conf.GetStringSlice("kafka.brokers"),
		GroupID:  conf.GetString("kafka.groupId"),
		Topic:    conf.GetString("kafka.topic"),
		MinBytes: 0,
		MaxBytes: 10e6, // 10MB
	})
	setupCloseReaderHandler(r)

	for i := 0; i < 2000; i++ {
		go consume(r)
	}
	done <- true
}

func consume(r *kafka.Reader) {
	for {
		m, err := r.ReadMessage(context.TODO())
		if err != nil {
			log.Fatal("Error reading:", err)
		}
		// fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		logger.Sugar().Infof("message at topic/partition/offset %v/%v/%v", m.Topic, m.Partition, m.Offset)
		parseKafkaMsg(m)
	}
}

func writeHbase(key string, value string) {
	logger.Sugar().Info("Writing : ", key)
	values := map[string]map[string][]byte{"cf": {"a": []byte(value)}}
	putRequest, _ := hrpc.NewPutStr(context.Background(), hbaseTable, key, values)
	_, err := client.Put(putRequest)

	// logger.Sugar().Info("Insert status : ", rsp)
	if err != nil {
		logger.Sugar().Error("error in writing to hbase - ", err)
	}
}

func parseKafkaMsg(msg kafka.Message) {
	tsa := md.TimeseriesArray{}
	json.Unmarshal(msg.Value, &tsa)
	for i := range tsa {
		// date := strings.Split(tsa[i].Timestamp.String(), "T")

		key := string(msg.Key) + "_" + tsa[i].Timestamp.String()
		logger.Debug(key)

		ts, _ := json.Marshal(tsa[i])
		val := string(ts)
		writeHbase(key, val)
	}
}

func setupCloseReaderHandler(r *kafka.Reader) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		logger.Info("\r- Ctrl+C pressed in Terminal")
		logger.Info("Closing reader...")
		if err := r.Close(); err != nil {
			logger.Fatal("Failed to close reader...")
		}
		os.Exit(0)
	}()
}
