package test

import (
	// "encoding/json"
	// "fmt"

	"net/http"
	"testing"
	"time"

	"github.com/gavv/httpexpect/v2"
	"github.com/go-openapi/strfmt"

	model "microsvc/pkg/model"
)

var (
	tsBaseURL  = "http://172.18.0.21:8000/api/v1/timeseries"
	baseEntity = "6fdae6af-226d-48bd-8b61-699758137eb3"
)

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

	obj.Value("TimeseriesUpload").Equal("Ok")

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

	obj.First().Object().Value("values")
}

func BenchmarkPostTimeSeries(b *testing.B) {

	e := httpexpect.New(b, tsBaseURL)

	for n := 0; n < b.N; n++ {
		// myid := uuid.New()
		myid := baseEntity
		tsa := getTS()
		obj := e.PUT(myid).
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

	val1 := model.TimeseriesValue{
		DataPointId: "dp1",
		Value:       "12",
		QualityCode: "0",
	}

	val2 := model.TimeseriesValue{
		DataPointId: "dp2",
		Value:       "15",
		QualityCode: "1",
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
