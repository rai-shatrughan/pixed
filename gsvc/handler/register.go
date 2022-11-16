package handler

import (
	"gsvc/handler/mc"
	"gsvc/handler/stream"
	"gsvc/pkg/util"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type svcHandler struct {
	mux         *http.ServeMux
	conf        *util.Config
	logger      *util.Logger
	kf          *util.KafkaWriter
	kv          *util.KV
	serviceName string
}

func (sh svcHandler) RegisterHandlers() {
	streamPath := sh.conf.GetString("basepath.stream")
	exchange := sh.conf.GetString("basepath.exchange")
	timeseries := sh.conf.GetString("basepath.timeseries")

	sh.mux.Handle(exchange, mc.PostMixedTimeseries(sh.kf, sh.logger))
	sh.mux.Handle(timeseries, mc.GetTimeseries(sh.kv, sh.logger))

	sh.mux.Handle(streamPath, stream.StreamHandler(sh.conf))
	sh.mux.Handle("/metrics", promhttp.Handler())
}

func New(mux *http.ServeMux, conf *util.Config, logger *util.Logger, kf *util.KafkaWriter, kv *util.KV, serviceName string) svcHandler {
	return svcHandler{
		mux:         mux,
		conf:        conf,
		logger:      logger,
		kf:          kf,
		kv:          kv,
		serviceName: serviceName,
	}
}
