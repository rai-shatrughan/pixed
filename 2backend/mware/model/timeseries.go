package model

import "time"

type Timeseries struct {
	Timestamp time.Time `json:"timestamp"`
	Property  string    `json:"property"`
	Unit      string    `json:"unit"`
	Value     float64   `json:"value"`
}

type TimeseriesArray []Timeseries
