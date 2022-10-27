package model

import (
	"github.com/go-openapi/strfmt"
)

type Timeseries struct {
	Timestamp *strfmt.DateTime  `json:"timestamp" validate:"required"`
	Values    []TimeseriesValue `json:"values" validate:"required,dive"`
}

type TimeseriesValue struct {
	DataPointId string  `json:"dataPointId" validate:"required"`
	Value       float64 `json:"value" validate:"required,min=0"`
	QualityCode float64 `json:"qualityCode" validate:"required,min=0"`
}

type TimeseriesArray []Timeseries

type TimeseriesUploadStatus struct {
	TimeseriesUpload string `json:"TimeseriesUpload"`
}
