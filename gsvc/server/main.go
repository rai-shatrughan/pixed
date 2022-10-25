package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"gsvc/domain/mc"
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

	muxRouter.Handle(streamPath, stream.StreamHandler(&conf))
	muxRouter.Handle(exPath, mc.PostTimeseries(&kf, &logger)).Methods("POST")
	muxRouter.Handle(exPath, mc.GetTimeseries(&kv, &logger)).Methods("GET")

	http.Handle("/", accessControl(muxRouter))

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

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
