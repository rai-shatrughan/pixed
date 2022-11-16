package test

import (
	"net/http"
	"strings"
	"testing"
	"time"

	cm "gsvc/test/common"
	"gsvc/test/perf"
)

func init() {
	cm.HostCheck()
	cm.Setup()
}

func TestPostTimeseries(t *testing.T) {
	tsBytes := cm.GetTSBytes()
	mt := perf.Metric{}
	mt.SetStartTime(time.Now())

	for _, agent := range cm.Agents {
		// time.Sleep(time.Millisecond * 500)
		// logger.Sugar().Infof("agent : %s", agent)
		st := time.Now()
		res, err := http.Post(cm.TSPostURL+agent, "application/json", strings.NewReader(tsBytes))
		dur := time.Since(st)
		if err != nil && res.StatusCode != 200 {
			t.Failed()
		}
		mt.IncrCount(1)
		mt.IncrDuration(dur)
	}
	mt.SetEndTime(time.Now())
	cm.Logger.Sugar().Infof("count - %v : duration - %v : start - %v : end - %v ", mt.Count(), mt.Duration(), mt.StartTime(), mt.EndTime())
}

func TestGetTimeseries(t *testing.T) {
	for _, agent := range cm.Agents {
		// logger.Sugar().Infof("agent : %s", agent)
		res, err := http.Get(cm.TSGetURL + agent)
		if err != nil && res.StatusCode != 200 {
			t.Failed()
		}
		// resp_body, _ := io.ReadAll(res.Body)
		// fmt.Printf("agent - %s , response : %s", agent, resp_body)
	}
}
