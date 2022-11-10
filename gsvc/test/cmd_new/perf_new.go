package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	cm "gsvc/test/common"
	"gsvc/test/perf"
)

type Result struct {
	ReqCount int64
	TotalDur time.Duration
	TotalErr int64
}

func init() {
	cm.HostCheck()
	cm.Setup()
}

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	var mt perf.Metric
	poolsize := 8000
	duration := 10 // seconds

	client := &http.Client{}
	//overriding the default parameters
	client.Transport = &http.Transport{
		DisableCompression:    false,
		DisableKeepAlives:     true,
		ResponseHeaderTimeout: time.Millisecond * time.Duration(30000),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(duration)*time.Second)
	defer cancel()

	resultChan := make(chan *Result, poolsize)
	agentChan := make(chan string, poolsize)
	maxGoroutine := make(chan struct{}, poolsize)

	mt.SetStartTime(time.Now())
	counter := 0

	aggResult := Result{}
	for {
		select {
		case <-sigChan:
			cm.Logger.Sugar().Infof("Result %v ", aggResult)
			return
		case <-ctx.Done():
			cm.Logger.Sugar().Infof("Result %v ", aggResult)
			return
		case agent := <-agentChan:
			maxGoroutine <- struct{}{}
			go ingest(*client, agent, resultChan, maxGoroutine)
		case result := <-resultChan:
			aggResult.ReqCount += result.ReqCount
			aggResult.TotalDur += result.TotalDur
			aggResult.TotalErr += result.TotalErr
		default:
			agentChan <- cm.Agents[counter%len(cm.Agents)]
			counter++
		}
	}
}

func ingest(client http.Client, agent string, resultChan chan<- *Result, maxGoroutine <-chan struct{}) {
	result := &Result{}
	tsBytes := cm.GetTSBytes()
	st := time.Now()
	req, _ := http.NewRequest("POST", cm.TSPostURL+agent, strings.NewReader(tsBytes))
	resp, err := client.Do(req)
	dur := time.Since(st)
	if resp != nil && resp.Body != nil {
		cm.Logger.Sugar().Infof("res", resp.StatusCode)
		resp.Body.Close()
	}
	if err != nil {
		cm.Logger.Sugar().Errorf("err", err)
		result.TotalErr++
	}
	result.TotalDur += dur
	result.ReqCount++
	resultChan <- result
	<-maxGoroutine
}
