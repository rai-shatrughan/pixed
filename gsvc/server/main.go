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

	"go.uber.org/zap"
)

var (
	conf     util.Config
	logger   util.Logger
	httpAddr string
	serveMux *http.ServeMux
	kf       util.KafkaWriter
	kv       util.KV
)

func init() {
	conf.New()
	logger.New()

	kv.New(&conf, &logger)
	kf.New(&conf, &logger)

	httpAddr = conf.GetString("http.address")
	serveMux = http.NewServeMux()
}

func main() {

	serveMux.Handle(conf.GetString("basepath.stream"), stream.StreamHandler(&conf))
	serveMux.Handle(conf.GetString("basepath.exchange"), mc.ExchangeHandler(&kf, &conf, &logger))

	http.Handle("/", accessControl(serveMux))

	startServer()

}

func startServer() {
	errs := make(chan error, 2)
	go func() {
		logger.Info("Server Started on", zap.String("Address", httpAddr))
		errs <- http.ListenAndServe(httpAddr, nil)
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
