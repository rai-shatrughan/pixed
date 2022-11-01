package main

import (
	// go modules
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	// git modules
	"github.com/segmentio/kafka-go"

	// websvc modules
	"gsvc/pkg/util"

	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
)

var (
	etc       = util.KV{}
	logger    = util.Logger{}
	client    gohbase.Client
	hbaseHost = "172.18.0.131"
)

func init() {
	conf := util.Config{}
	conf.New()

	logger.New()
	etc.New(&conf, &logger)
	client = gohbase.NewClient(hbaseHost)

}

func main() {
	// createTable()
	readMessage()
}

func createTable() {
	testTableName := "ts-table"

	var cFamilies = map[string]map[string]string{
		"cf":  nil,
		"cf2": nil,
	}

	ac := gohbase.NewAdminClient(hbaseHost)
	crt := hrpc.NewCreateTable(context.Background(), []byte(testTableName), cFamilies)

	if err := ac.CreateTable(crt); err != nil {
		fmt.Errorf("CreateTable returned an error: %v", err)
	}
}

func readMessage() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"172.18.0.41:9092"},
		GroupID:  "consumer-group-id",
		Topic:    "ts",
		MinBytes: 0,    // 10KB
		MaxBytes: 10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(context.TODO())
		if err != nil {
			log.Fatal("Error reading:", err)
		}
		// fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		fmt.Printf("message at topic/partition/offset %v/%v/%v: \n", m.Topic, m.Partition, m.Offset)

		// Values maps a ColumnFamily -> Qualifiers -> Values.
		values := map[string]map[string][]byte{"cf": map[string][]byte{"a": []byte{0}}}
		putRequest, err := hrpc.NewPutStr(context.Background(), "ts-table", "key", values)
		rsp, err := client.Put(putRequest)

		fmt.Println("HELLO", rsp)
		if err != nil {
			fmt.Printf("Error ETC - %s", err)
		}
		time.Sleep(10000)
	}

}

// func groupConsumer() {
// 	done := make(chan bool)
// 	for i := 0; i < kafkaReaderCount; i++ {
// 		go readKafkaMsg(kf.Readers[i], kafkaMsgChan)
// 	}

// 	for i := 0; i < tsParserCount; i++ {
// 		go parseKafkaMsg(kafkaMsgChan, tskvChan)
// 	}

// 	for i := 0; i < kvWriterCount; i++ {
// 		go writeKV(tskvChan, kvCounterChan)
// 	}

// 	go msgCounter(kvCounterChan)

// 	quitHandler()
// 	<-done
// }

// func msgCounter(kvCounterChan <-chan int8) {
// 	var kvCounter int64
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			logger.Info("stopping counter")
// 			return
// 		case <-kvCounterChan:
// 			kvCounter++
// 			logger.Info("Total messages written to etcd", zap.Int64("count", kvCounter))
// 		}
// 	}

// }

// func writeKV(tskvChan <-chan tskv, kvCounterChan chan<- int8) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			logger.Info("stopping parser")
// 			return
// 		default:
// 			kv := <-tskvChan
// 			etc.Put(kv.key, kv.value)
// 			kvCounterChan <- 1
// 		}
// 	}

// }

// func parseKafkaMsg(msgChan <-chan kafka.Message, tskvChan chan<- tskv) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			logger.Info("stopping parser")
// 			return
// 		default:
// 			start := time.Now()
// 			msg := <-msgChan
// 			tsa := md.TimeseriesArray{}
// 			json.Unmarshal(msg.Value, &tsa)
// 			for i := range tsa {
// 				date := strings.Split(tsa[i].Timestamp.String(), "T")

// 				key := "/" + string(msg.Key) + "/" + date[0] + "/" + tsa[i].Timestamp.String()
// 				logger.Debug(key)

// 				ts, _ := json.Marshal(tsa[i])
// 				val := string(ts)
// 				logger.Debug(val)
// 				tskvChan <- tskv{key: key, value: val}
// 			}
// 			logger.Debug("Time elapsed for parser",
// 				zap.String("duration", time.Since(start).String()),
// 			)
// 		}
// 	}
// }

// func readKafkaMsg(reader *kafka.Reader, kafkaMsgChan chan<- kafka.Message) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			logger.Info("stopping reader")
// 			return
// 		default:
// 			msg, err := kf.Read(reader)
// 			if err != nil {
// 				logger.Error(err.Error())
// 			} else {
// 				kafkaMsgChan <- msg
// 			}
// 		}

// 	}
// }

func quitHandler(r *kafka.Reader) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig
		logger.Info("\r- Ctrl+C pressed - stopping writer now")
		if err := r.Close(); err != nil {
			log.Fatal("failed to close reader:", err)
		}
	}()
}
