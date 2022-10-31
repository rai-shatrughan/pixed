package test

import (
	// "encoding/json"
	// "fmt"

	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gavv/httpexpect/v2"
	"github.com/go-openapi/strfmt"

	model "gsvc/pkg/model"
)

var (
	baseEntity          = "6fdae6af-226d-48bd-8b61-699758137eb3"
	tsPostURL, tsGetURL string
	host                = flag.String("host", "gsvc", "Name of service to test")
)

func hostCheck() {
	if *host == "gsvc" {
		tsPostURL = "http://172.18.0.21:8000/api/mindconnect/v3/exchange/"
		tsGetURL = "http://172.18.0.21:8000/api/iottimeseries/v3/timeseries/"
		// fmt.Println(*host)
	} else {
		tsPostURL = "http://172.18.0.22:9000/api/v1/exchange"
		tsGetURL = "http://172.18.0.21:9000/api/iottimeseries/v3/timeseries/"
		// fmt.Println(*host)
	}
}

func TestPostTimeseries(t *testing.T) {
	hostCheck()
	tsBytes := getTSBytes()
	e := httpexpect.New(t, tsPostURL)
	obj := e.POST(baseEntity).
		WithHeader("X-API-Key", "sr12345").
		WithHeader("Content-Type", "application/json").
		WithBytes(tsBytes).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	obj.Value("TimeseriesUpload").Equal("Ok")

}

func TestGetTimeSeries(t *testing.T) {
	hostCheck()
	e := httpexpect.New(t, tsGetURL)
	obj := e.GET(baseEntity).
		WithHeader("X-API-Key", "sr12345").
		WithQuery("duration", "1m").
		Expect().
		Status(http.StatusOK).
		JSON().
		Array()

	obj.First().Object().Value("values")
}

func BenchmarkPostTimeSeries(b *testing.B) {
	hostCheck()
	e := httpexpect.New(b, tsPostURL)

	for n := 0; n < b.N; n++ {
		// myid := uuid.New()
		myid := baseEntity
		tsBytes := getTSBytes()
		obj := e.POST(myid).
			WithHeader("X-API-Key", "sr12345").
			WithHeader("Content-Type", "application/json").
			WithBytes(tsBytes).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object()

		obj.Value("TimeseriesUpload").Equal("Ok")
	}
}

func getTSBytes() []byte {

	dateTime := strfmt.DateTime(time.Now().UTC())

	dp1 := "dp1"
	dp2 := "dp2"
	v1 := 12.0
	v2 := 100.0
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
		Timestamp: &dateTime,
		Values:    []model.TimeseriesValue{val1, val2},
	}

	ts2 := model.Timeseries{
		Timestamp: &dateTime,
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

	out := []byte(fmt.Sprintf("%s %s %s ", mp1, string(o), mp2))
	fmt.Println("tsa - ", string(out))
	return out
}
