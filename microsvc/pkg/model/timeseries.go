package model

import (
	"github.com/go-openapi/strfmt"
)

type Timeseries struct {
	Timestamp *strfmt.DateTime  `json:"timestamp"`
	Values    []TimeseriesValue `json:"values"`
}

type TimeseriesValue struct {
	DataPointId string      `json:"dataPointId"`
	Value       interface{} `json:"value"`
	QualityCode interface{} `json:"qualityCode"`
}

type TimeseriesArray []Timeseries

type TimeseriesUploadStatus struct {
	TimeseriesUpload string `json:"TimeseriesUpload"`
}
