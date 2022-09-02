package test

import (
	// "encoding/json"
	// "fmt"

	"net/http"
	"testing"
	"time"

	"github.com/gavv/httpexpect/v2"
	"github.com/google/uuid"
)

var (
	tsBaseURL  = "http://172.18.0.21:8002/api/v1/timeseries"
	baseEntity = "6fdae6af-226d-48bd-8b61-699758137eb3"
)

type TimeSeries struct {
	Timestamp string `json:"timestamp"`
	Property  string `json:"property"`
	Unit      string `json:"unit"`
	Value     int    `json:"value"`
}

func TestPutTimeSeries(t *testing.T) {
	tsa := getTS()
	e := httpexpect.New(t, tsBaseURL)
	obj := e.PUT(baseEntity).
		WithHeader("X-API-Key", "sr12345").
		WithHeader("Content-Type", "application/json").
		WithJSON(tsa).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	obj.Value("TimeseriesUpload").Equal("ok")

}

func TestGetTimeSeries(t *testing.T) {
	e := httpexpect.New(t, tsBaseURL)
	obj := e.GET(baseEntity).
		WithHeader("X-API-Key", "sr12345").
		WithQuery("duration", "1m").
		Expect().
		Status(http.StatusOK).
		JSON().
		Array()

	obj.First().Object().Value("property")
}

func BenchmarkPostTimeSeries(b *testing.B) {

	e := httpexpect.New(b, tsBaseURL)

	for n := 0; n < b.N; n++ {
		myid := uuid.New()
		tsa := getTS()
		obj := e.PUT(myid.String()).
			WithHeader("X-API-Key", "sr12345").
			WithHeader("Content-Type", "application/json").
			WithJSON(tsa).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object()

		obj.Value("TimeseriesUpload").Equal("ok")
	}
}

func getTS() []TimeSeries {
	loc, _ := time.LoadLocation("UTC")

	ts1 := TimeSeries{
		Timestamp: time.Now().In(loc).Format("2006-01-02T15:04:05.012Z"),
		Property:  "temperature",
		Unit:      "celcius",
		Value:     10,
	}

	ts2 := TimeSeries{
		Timestamp: time.Now().In(loc).Format("2006-01-02T15:04:05.012Z"),
		Property:  "pressure",
		Unit:      "psi",
		Value:     20,
	}

	tsa := []TimeSeries{
		ts1,
		ts2,
	}

	// o1, _ := json.Marshal(ts1)
	// o, _ := json.Marshal(tsa)
	// fmt.Println("o1 - ", string(o1))
	// fmt.Println("tsa - ", string(o))

	return tsa
}
