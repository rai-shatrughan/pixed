package test

import (
	// "encoding/json"
	// "fmt"

	"flag"
	"net/http"
	"testing"
	"time"

	"github.com/gavv/httpexpect/v2"
	"github.com/go-openapi/strfmt"

	model "gsvc/pkg/model"
)

var (
	baseEntity = "6fdae6af-226d-48bd-8b61-699758137eb3"
	tsBaseURL  string
	host       = flag.String("host", "gsvc", "Name of service to test")
)

func hostCheck() {
	if *host == "gsvc" {
		tsBaseURL = "http://172.18.0.21:8000/api/v1/exchange"
		// fmt.Println(*host)
	} else {
		tsBaseURL = "http://172.18.0.22:9000/api/v1/exchange"
		// fmt.Println(*host)
	}
}

func TestPostTimeseries(t *testing.T) {
	hostCheck()
	tsa := getTS()
	e := httpexpect.New(t, tsBaseURL)
	obj := e.POST(baseEntity).
		WithHeader("X-API-Key", "sr12345").
		WithHeader("Content-Type", "application/json").
		WithJSON(tsa).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	obj.Value("TimeseriesUpload").Equal("Ok")

}

func TestGetTimeSeries(t *testing.T) {
	hostCheck()
	e := httpexpect.New(t, tsBaseURL)
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
	e := httpexpect.New(b, tsBaseURL)

	for n := 0; n < b.N; n++ {
		// myid := uuid.New()
		myid := baseEntity
		tsa := getTS()
		obj := e.POST(myid).
			WithHeader("X-API-Key", "sr12345").
			WithHeader("Content-Type", "application/json").
			WithJSON(tsa).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object()

		obj.Value("TimeseriesUpload").Equal("Ok")
	}
}

func getTS() []model.Timeseries {

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

	// o1, _ := json.Marshal(ts1)
	// o, _ := json.Marshal(tsa)
	// fmt.Println("o1 - ", string(o1))
	// fmt.Println("tsa - ", string(o))

	return tsa
}
