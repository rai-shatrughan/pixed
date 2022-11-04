package main

import (
	// go modules
	"context"
	"encoding/json"
	"os"
	"os/signal"
	"strings"
	"sync"
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
	reader     *kafka.Reader
)

func init() {

	conf.New()

	hbaseHost = conf.GetString("hbase.zookeeper")
	hbaseTable = conf.GetString("kafka.topic")

	logger.New()
	client = gohbase.NewClient(hbaseHost)

	reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  conf.GetStringSlice("kafka.brokers"),
		GroupID:  conf.GetString("kafka.groupId"),
		Topic:    conf.GetString("kafka.topic"),
		MinBytes: 0,
		MaxBytes: 10e6, // 10MB
	})

}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	quitHandler(cancel)
	createTable()
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go consume(&wg, ctx, i)
	}
	wg.Wait()
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
			if strings.Contains(err.Error(), "TableExists") {
				created = true
				logger.Sugar().Info("table exists")
				break
			}
			logger.Sugar().Error("createTable returned an error: %v", err)
			retryCount++
			logger.Sugar().Error("retry count : ", retryCount)
		} else {
			logger.Sugar().Info("table created in attempt : ", retryCount)
			created = true
		}

	}
}

func consume(wg *sync.WaitGroup, ctx context.Context, num int) {
	for {
		select {
		case <-ctx.Done():
			logger.Sugar().Info("exiting consumer - ", num)
			wg.Done()
			return
		default:
			logger.Sugar().Infof("consumer - %v", num)
			ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			m, err := reader.ReadMessage(ctxTimeout)
			if err != nil {
				if strings.Contains(err.Error(), "context deadline exceeded") {
					logger.Sugar().Infof("no new kafka msg in last 10 seconds")
					break
				}
				logger.Sugar().Infof("error reading kafka:", err)
				break
			}
			// fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
			logger.Sugar().Infof("consumer - %v : message at topic/partition/offset %v/%v/%v", num, m.Topic, m.Partition, m.Offset)

			tsa := md.TimeseriesArray{}
			json.Unmarshal(m.Value, &tsa)
			for i := range tsa {
				// date := strings.Split(tsa[i].Timestamp.String(), "T")

				key := string(m.Key) + "_" + tsa[i].Timestamp.String()
				logger.Debug(key)

				ts, _ := json.Marshal(tsa[i])
				value := string(ts)

				logger.Sugar().Info("writing : ", key)
				values := map[string]map[string][]byte{"cf": {"a": []byte(value)}}
				putRequest, _ := hrpc.NewPutStr(context.Background(), hbaseTable, key, values)
				rsp, err := client.Put(putRequest)

				logger.Sugar().Info("Insert status : ", rsp)
				if err != nil {
					logger.Sugar().Error("error in writing to hbase - ", err)
				}
			}
		}
	}

}

func quitHandler(cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, os.Interrupt)
	go func() {
		<-c
		logger.Info("\r- Ctrl+C pressed - stopping consumer now")
		cancel()
	}()
}
