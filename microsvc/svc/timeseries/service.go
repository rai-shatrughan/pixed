package timeseries

import (
	"encoding/json"
	"fmt"
	model "microsvc/pkg/model"

	mw "microsvc/pkg/mware"
)

type Service interface {
	PutTimeseries(string, *model.TimeseriesArray) *model.TimeseriesUploadStatus
	GetTimeseries(string) string
}

type service struct {
	kf      mw.KafkaWriter
	kv      mw.KV
	assetId string
}

func (s *service) PutTimeseries(assetId string, tsa *model.TimeseriesArray) *model.TimeseriesUploadStatus {
	kfmsg, _ := json.Marshal(tsa)
	err := s.kf.Write([]byte(assetId), []byte(kfmsg))
	if err != nil {
		return &model.TimeseriesUploadStatus{TimeseriesUpload: "Failed"}
	} else {
		return &model.TimeseriesUploadStatus{TimeseriesUpload: "Ok"}
	}
}

func (s *service) GetTimeseries(assetId string) string {
	resp := s.kv.GetFromKeyWithLimit("/"+assetId, 1000)
	return resp
}

func NewService() *service {
	var (
		topic   = "ts"
		groupID = "ts-consumer-group"
		kf      = mw.KafkaWriter{}
		brokers []string

		etc = mw.KV{}
	)

	conf := mw.Config{}
	conf.New()

	brokers = conf.GetStringSlice("kafka.brokers")
	kf.GroupID = &groupID
	kf.Topic = &topic
	kf.Brokers = brokers

	kf.New()
	etc.New()

	fmt.Println("kf", kf)
	fmt.Println("kv", etc)

	return &service{
		kf:      kf,
		kv:      etc,
		assetId: "",
	}
}
