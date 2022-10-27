package model

import (
	"errors"

	"github.com/go-openapi/strfmt"
)

type Timeseries struct {
	Timestamp *strfmt.DateTime  `json:"timestamp"`
	Values    []TimeseriesValue `json:"values"`
}

type TimeseriesValue struct {
	DataPointId *string  `json:"dataPointId"`
	Value       *float64 `json:"value"`
	QualityCode *float64 `json:"qualityCode"`
}

type TimeseriesArray []Timeseries

func (tsa *TimeseriesArray) Validate() error {
	err := errors.New("invalid timeseries")

	if len(*tsa) == 0 {
		return err
	}

	for _, ts := range *tsa {
		if ts.Timestamp == nil {
			return err
		} else if len(ts.Values) == 0 {
			return err
		}
	}

	for _, ts := range *tsa {
		for _, tsv := range ts.Values {
			if tsv.DataPointId == nil {
				return err
			}
			if tsv.QualityCode == nil {
				return err
			}
			if tsv.Value == nil {
				return err
			}
		}
	}

	return nil
}

type TimeseriesUploadStatus struct {
	TimeseriesUpload string `json:"TimeseriesUpload"`
}
