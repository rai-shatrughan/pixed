package test

import (
	"math/rand"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/gavv/httpexpect/v2"
)

var (
	tsBaseURL     = "http://172.18.0.21:8002/api/iottimeseries/v3"
	entityID28Len = "550a5500b55c55d5e12345678x"
	max           = 99999
	min           = 10000
)

type TimeSeriesDataItem struct {
	Time        string      `json:"_time"`
	Temperature interface{} `json:"temperature"`
	Pressure    interface{} `json:"pressure"`
}

type TimeSeriesItem struct {
	Data            []TimeSeriesDataItem `json:"data"`
	EntityID        string               `json:"entityId"`
	PropertySetName string               `json:"propertySetName"`
}

type TimeSeries struct {
	Timeseries []TimeSeriesItem `json:"timeseries"`
}

func TestPutTimeSeries(t *testing.T) {
	tsa := getTS("550e8400e29b41d4a716446655440000")
	e := httpexpect.New(t, tsBaseURL)
	obj := e.PUT("/timeseries").
		WithHeader("Authorization", "Bearer sr12345").
		WithHeader("Content-Type", "application/json").
		WithJSON(tsa).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	obj.Value("TimeseriesUpload").String().Equal("ok")
}

func TestGetTimeSeries(t *testing.T) {
	e := httpexpect.New(t, tsBaseURL)
	obj := e.GET("/timeseries/550e8400e29b41d4a716446655440000/p1").
		WithHeader("Authorization", "Bearer sr12345").
		WithHeader("Content-Type", "application/json").
		Expect().
		Status(http.StatusOK).
		JSON().
		Array()

	obj.First().Object().Value("pressure").Number().Equal(20)
}

func BenchmarkPostTimeSeries(b *testing.B) {

	e := httpexpect.New(b, tsBaseURL)

	for n := 0; n < b.N; n++ {
		i := rand.Intn(max-min) + min
		tsa := getTS(entityID28Len + strconv.Itoa(i))
		obj := e.PUT("/timeseries").
			WithHeader("Authorization", "Bearer sr12345").
			WithHeader("Content-Type", "application/json").
			WithJSON(tsa).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object()

		obj.Value("TimeseriesUpload").String().Equal("ok")
	}
}

func getTS(entityID string) TimeSeries {
	loc, _ := time.LoadLocation("UTC")

	tsi := TimeSeriesDataItem{
		Time:        time.Now().In(loc).Format("2006-01-02T15:04:05.000Z"),
		Temperature: 50,
		Pressure:    20,
	}

	ts1 := TimeSeriesItem{EntityID: entityID, PropertySetName: "p1", Data: []TimeSeriesDataItem{tsi}}
	ts2 := TimeSeriesItem{EntityID: entityID, PropertySetName: "p2", Data: []TimeSeriesDataItem{tsi}}
	ts := []TimeSeriesItem{
		ts1,
		ts2,
	}

	tsa := TimeSeries{ts}
	// o1, _ := json.Marshal(ts1)
	// o, _ := json.Marshal(tsa)
	// fmt.Println("o1 - ", string(o1))
	// fmt.Println("tsa - ", string(o))

	return tsa
}
