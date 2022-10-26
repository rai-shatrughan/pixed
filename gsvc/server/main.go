package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"gsvc/domain/mc"
	"gsvc/domain/mware"
	"gsvc/domain/stream"
	"gsvc/pkg/util"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var (
	conf      util.Config
	logger    util.Logger
	httpAddr  string
	muxRouter *mux.Router
	kf        util.KafkaWriter
	kv        util.KV
)

func init() {
	conf.New()
	logger.New()

	kv.New(&conf, &logger)
	kf.New(&conf, &logger)

	httpAddr = conf.GetString("http.address")
	muxRouter = mux.NewRouter()
}

func main() {
	streamPath := conf.GetString("basepath.stream")
	exPath := conf.GetString("basepath.exchange") + "{assetId}"

	muxRouter.
		Path(exPath).
		Handler(mc.PostTimeseries(&kf, &logger)).
		Methods("POST")

	muxRouter.
		Path(exPath).
		Handler(mc.GetTimeseries(&kv, &logger)).
		Methods("GET")

	muxRouter.
		Path(streamPath).
		Handler(stream.StreamHandler(&conf))

	muxRouter.Use(mware.Logging)
	muxRouter.Use(mware.AccessControl)
	muxRouter.Use(mware.ResponseContentType)

	startServer()

}

func startServer() {
	errs := make(chan error, 2)
	go func() {
		logger.Info("Server Started on", zap.String("Address", httpAddr))
		errs <- http.ListenAndServe(httpAddr, muxRouter)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()
	<-errs
	logger.Info("terminated server")
}
