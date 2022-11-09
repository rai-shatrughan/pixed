package main

import (
	"context"
	"net/http"
	"strings"
	"time"

	cm "gsvc/test/common"
	"gsvc/test/perf"
)

type result struct {
	resp *http.Response
	dur  time.Duration
	err  error
}

func init() {
	cm.HostCheck()
	cm.Setup()
}

func main() {
	var mt perf.Metric
	poolsize := 1000

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resultChan := make(chan result, poolsize)
	agentChan := make(chan string, poolsize)

	mt.SetStartTime(time.Now())
	go updateAgentChan(poolsize, agentChan)
	for {
		select {
		case <-ctx.Done():
			mt.SetEndTime(time.Now())
			cm.Logger.Sugar().Infof("count - %v : duration - %v : start - %v : end - %v ", mt.Count(), mt.Duration(), mt.StartTime(), mt.EndTime())
			return
		case agent := <-agentChan:
			go ingest(agent, resultChan)
		case result := <-resultChan:
			updateMetric(result, &mt)
		}
	}
}

func ingest(agent string, resultChan chan<- result) {
	tsBytes := cm.GetTSBytes()
	st := time.Now()
	response, err := http.Post(cm.TSPostURL+agent, "application/json", strings.NewReader(tsBytes))
	d := time.Since(st)
	cm.Logger.Sugar().Infof("res - %v : dur - %v : err - %v ", response.StatusCode, d, err)
	resultChan <- result{response, d, err}
}

func updateMetric(res result, mt *perf.Metric) {
	mt.IncrCount(1)
	if res.err != nil {
		mt.IncrFailed(1)
		cm.Logger.Sugar().Errorf("error in call", res.err)
	}
	mt.IncrDuration(res.dur)
	cm.Logger.Sugar().Infof("count - %v : failed - %v ", mt.Count(), mt.Failed())
}

func updateAgentChan(poolsize int, agentChan chan<- string) {
	for {
		for i := 0; i < poolsize; i++ {
			agentChan <- cm.Agents[i%400]
		}
		time.Sleep(1 * time.Second)
	}
}
