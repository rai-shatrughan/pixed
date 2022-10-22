package timeseries

import (
	"encoding/json"
	model "microsvc/pkg/model"

	mw "microsvc/pkg/mware"
)

type Service interface {
	PutTimeseries(string, *model.TimeseriesArray) *model.TimeseriesUploadStatus
	GetTimeseries(string) *model.TimeseriesArray
}

type service struct {
	kf      mw.KafkaWriter
	kv      mw.KV
	assetId string
}

func (s *service) PutTimeseries(assetId string, tsa *model.TimeseriesArray) *model.TimeseriesUploadStatus {
	kfmsg, err := json.Marshal(tsa)

	if err != nil {
		return &model.TimeseriesUploadStatus{TimeseriesUpload: "Failed"}
	}

	err = s.kf.Write([]byte(assetId), []byte(kfmsg))

	if err != nil {
		return &model.TimeseriesUploadStatus{TimeseriesUpload: "Failed"}
	} else {
		return &model.TimeseriesUploadStatus{TimeseriesUpload: "Ok"}
	}
}

func (s *service) GetTimeseries(assetId string) *model.TimeseriesArray {
	res := s.kv.GetFromKeyWithLimit("/"+assetId, 1000)
	tsa := model.TimeseriesArray{}
	_ = json.Unmarshal([]byte(res), &tsa)
	return &tsa
}

func NewService(conf *mw.Config, logger *mw.Logger) *service {
	kf := mw.KafkaWriter{}
	etc := mw.KV{}

	kf.New(conf, logger)
	etc.New(conf, logger)

	return &service{
		kf:      kf,
		kv:      etc,
		assetId: "",
	}
}
