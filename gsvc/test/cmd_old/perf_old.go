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
	// conGoroutines := make(chan struct{}, poolsize)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(duration+2)*time.Second)
	defer cancel()

	resultChan := make(chan *Result, poolsize)
	agentChan := make(chan string, poolsize)

	mt.SetStartTime(time.Now())
	// counter := 0
	go fillAgentChan(ctx, agentChan)

	for i := 0; i < poolsize; i++ {
		agent := cm.Agents[i%len(cm.Agents)]
		go ingest(duration, agent, resultChan)
	}

	responders := 0
	aggResult := Result{}
	for responders < poolsize {
		select {
		case <-sigChan:
			cm.Logger.Sugar().Infof("Result %v ", aggResult)
			return
		case <-ctx.Done():
			cm.Logger.Sugar().Infof("Result %v ", aggResult)
			return
		case result := <-resultChan:
			aggResult.ReqCount += result.ReqCount
			aggResult.TotalDur += result.TotalDur
			aggResult.TotalErr += result.TotalErr
			responders++
		}
	}
	cm.Logger.Sugar().Infof("Result %v ", aggResult)
}

func ingest(duration int, agent string, resultChan chan<- *Result) {
	client := &http.Client{}
	//overriding the default parameters
	client.Transport = &http.Transport{
		DisableCompression:    false,
		DisableKeepAlives:     true,
		ResponseHeaderTimeout: time.Millisecond * time.Duration(30000),
	}

	result := &Result{}
	start := time.Now()
	for time.Since(start).Seconds() <= float64(duration) {
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
		// cm.Logger.Sugar().Infof("count %v", result.ReqCount)
		// time.Sleep(1 * time.Second)
	}
	resultChan <- result
}

func fillAgentChan(ctx context.Context, agentChan chan<- string) {
	counter := 0
	for {
		select {
		case <-ctx.Done():
			return
		default:
			agentChan <- cm.Agents[counter%len(cm.Agents)]
			counter++
		}
	}
}
