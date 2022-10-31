package handler

import (
	"gsvc/handler/mc"
	"gsvc/handler/stream"
	"gsvc/pkg/util"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type SVCHandlers struct {
	MuxRouter *mux.Router
	Conf      *util.Config
	Logger    *util.Logger
	Kf        *util.KafkaWriter
	Kv        *util.KV
}

func (sh SVCHandlers) RegisterHandlers() {
	streamPath := sh.Conf.GetString("basepath.stream")
	exchange := sh.Conf.GetString("basepath.exchange") + "{assetId}"
	timeseries := sh.Conf.GetString("basepath.timeseries") + "{assetId}"

	sh.MuxRouter.
		Path(exchange).
		Handler(mc.PostMixedTimeseries(sh.Kf, sh.Logger)).
		Methods("POST")

	sh.MuxRouter.
		Path(timeseries).
		Handler(mc.GetTimeseries(sh.Kv, sh.Logger)).
		Methods("GET")

	sh.MuxRouter.
		Path(streamPath).
		Handler(stream.StreamHandler(sh.Conf))

	sh.MuxRouter.Path("/metrics").Handler(promhttp.Handler())
}
