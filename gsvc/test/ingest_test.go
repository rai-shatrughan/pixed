package test

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/go-openapi/strfmt"

	model "gsvc/pkg/model"
	"gsvc/pkg/util"
	"gsvc/test/data"
)

var (
	tsPostURL, tsGetURL string
	host                = flag.String("host", "gsvc", "Name of service to test")
	agents              = data.Agents
	kv                  util.KV
	logger              util.Logger
	conf                util.Config
)

func init() {
	hostCheck()
	logger.New()
	conf.New()

	kv.New(&conf, &logger)
}

func hostCheck() {
	if *host == "gsvc" {
		tsPostURL = "http://172.18.0.21:8000/api/mindconnect/v3/exchange/"
		tsGetURL = "http://172.18.0.21:8000/api/iottimeseries/v3/timeseries/"
	} else {
		tsPostURL = "http://172.18.0.22:9000/api/v1/exchange/"
		tsGetURL = "http://172.18.0.21:9000/api/iottimeseries/v3/timeseries/"
	}
}

func TestPostTimeseries(t *testing.T) {
	tsBytes := getTSBytes()
	for _, agent := range agents {
		// time.Sleep(time.Millisecond * 500)
		logger.Sugar().Infof("agent : %s", agent)
		res, err := http.Post(tsPostURL+agent, "application/json", strings.NewReader(tsBytes))
		if err != nil && res.StatusCode != 200 {
			t.Failed()
		}
	}
}

func TestGetTimeseries(t *testing.T) {
	for _, agent := range agents {
		logger.Sugar().Infof("agent : %s", agent)
		res, err := http.Get(tsGetURL + agent)
		if err != nil && res.StatusCode != 200 {
			t.Failed()
		}
		resp_body, _ := io.ReadAll(res.Body)
		logger.Sugar().Infof("agent - %s , response : %s", agent, resp_body)
	}
}

// func TestGetKV(t *testing.T) {
// 	for _, agent := range agents {
// 		res, _ := kv.Get(agent + "_")
// 		logger.Sugar().Infof("agent - %s , response : %s", agent, res)
// 	}
// }

func getTSBytes() string {
	nw := time.Now().UTC()
	bf := nw.Add(-time.Second * 10)
	dt1 := strfmt.DateTime(nw)
	dt2 := strfmt.DateTime(bf)

	dp1 := "dp1"
	dp2 := "dp2"
	v1 := float64(nw.Second())
	v2 := float64(bf.Second())
	q1 := 0.0
	q2 := 1.0

	val1 := model.TimeseriesValue{
		DataPointId: &dp1,
		Value:       &v1,
		QualityCode: &q1,
	}

	val2 := model.TimeseriesValue{
		DataPointId: &dp2,
		Value:       &v2,
		QualityCode: &q2,
	}

	ts1 := model.Timeseries{
		Timestamp: &dt1,
		Values:    []model.TimeseriesValue{val1, val2},
	}

	ts2 := model.Timeseries{
		Timestamp: &dt2,
		Values:    []model.TimeseriesValue{val1, val2},
	}

	tsa := []model.Timeseries{
		ts1,
		ts2,
	}

	o, _ := json.Marshal(tsa)

	mp1 := `--f0Ve5iPP2ySppIcDSR6Bak
	Content-Type: multipart/related;boundary=penFL6sBQHJJUN3HA4ftqC
	
	--penFL6sBQHJJUN3HA4ftqC
	Content-Type: application/vnd.siemens.mindsphere.meta+json
	
	{
		"type": "item",
		"version": "1.0",
		"payload": {
			"type": "standardTimeSeries",
			"version": "1.0",
			"details": {
				"configurationId": ""
			}
		}
	}
	--penFL6sBQHJJUN3HA4ftqC
	Content-Type: application/json`

	mp2 := `--penFL6sBQHJJUN3HA4ftqC--
	--f0Ve5iPP2ySppIcDSR6Bak--`

	// out := []byte(fmt.Sprintf("%s %s %s ", mp1, string(o), mp2))
	// fmt.Println("tsa - ", string(out))
	out := fmt.Sprintf("%s %s %s ", mp1, string(o), mp2)
	return out
}
