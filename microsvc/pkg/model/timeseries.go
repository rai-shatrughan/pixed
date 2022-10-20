package model

import (
	"github.com/go-openapi/strfmt"
)

type Timeseries struct {
	Timestamp *strfmt.DateTime `json:"timestamp"`
	Property  string           `json:"property"`
	Unit      string           `json:"unit"`
	Value     float64          `json:"value"`
}

type TimeseriesArray []Timeseries

type TimeseriesUploadStatus struct {
	TimeseriesUpload string `json:"TimeseriesUpload"`
}
