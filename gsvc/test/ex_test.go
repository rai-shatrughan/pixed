package test

import (
	// "encoding/json"
	// "fmt"

	"net/http"
	"testing"
	"time"

	"github.com/gavv/httpexpect/v2"
	"github.com/go-openapi/strfmt"

	model "gsvc/pkg/model"
)

var (
	exBaseURL    = "http://localhost:9000/api/v1/exchange"
	exBaseEntity = "6fdae6af-226d-48bd-8b61-699758137eb3"
)

func TestPostExTimeSeries(t *testing.T) {
	tsa := getExTS()
	e := httpexpect.New(t, exBaseURL)
	obj := e.POST(exBaseEntity).
		WithHeader("X-API-Key", "sr12345").
		WithHeader("Content-Type", "application/json").
		WithJSON(tsa).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	obj.Value("TimeseriesUpload").Equal("Ok")

}

func getExTS() []model.Timeseries {

	dateTime := strfmt.DateTime(time.Now().UTC())

	val1 := model.TimeseriesValue{
		DataPointId: "dp1",
		Value:       12,
		QualityCode: 0,
	}

	val2 := model.TimeseriesValue{
		DataPointId: "dp2",
		Value:       15,
		QualityCode: 1,
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
