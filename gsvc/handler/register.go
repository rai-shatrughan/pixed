package handler

import (
	"gsvc/handler/mc"
	"gsvc/handler/stream"
	"gsvc/pkg/util"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterHandlers(st *util.AppState) {
	streamPath := st.Conf.GetString("basepath.stream")
	exchange := st.Conf.GetString("basepath.exchange")
	timeseries := st.Conf.GetString("basepath.timeseries")

	st.Mux.Handle(exchange, mc.PostMixedTimeseries(st))
	st.Mux.Handle(timeseries, mc.GetTimeseries(st))

	st.Mux.Handle(streamPath, stream.StreamHandler(st))
	st.Mux.Handle("/metrics", promhttp.Handler())
}
