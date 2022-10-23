package timeseries

import (
	"encoding/json"
	model "microsvc/pkg/model"

	"microsvc/pkg/util"
)

type Service interface {
	PutTimeseries(string, *model.TimeseriesArray) (*model.TimeseriesUploadStatus, error)
	GetTimeseries(string) (*model.TimeseriesArray, error)
}

type service struct {
	kf      util.KafkaWriter
	kv      util.KV
	assetId string
}

func (s *service) PutTimeseries(assetId string, tsa *model.TimeseriesArray) (*model.TimeseriesUploadStatus, error) {
	kfmsg, _ := json.Marshal(tsa)

	err := s.kf.Write([]byte(assetId), []byte(kfmsg))

	if err != nil {
		return &model.TimeseriesUploadStatus{TimeseriesUpload: "Failed"}, err
	} else {
		return &model.TimeseriesUploadStatus{TimeseriesUpload: "Ok"}, nil
	}
}

func (s *service) GetTimeseries(assetId string) (*model.TimeseriesArray, error) {
	tsa := model.TimeseriesArray{}

	res, err := s.kv.GetFromKeyWithLimit("/"+assetId, 1000)
	if err != nil {
		return &tsa, err
	}

	err = json.Unmarshal([]byte(res), &tsa)
	if err != nil {
		return &tsa, err
	}

	return &tsa, err
}

func NewService(conf *util.Config, logger *util.Logger) *service {
	kf := util.KafkaWriter{}
	etc := util.KV{}

	kf.New(conf, logger)
	etc.New(conf, logger)

	return &service{
		kf:      kf,
		kv:      etc,
		assetId: "",
	}
}
