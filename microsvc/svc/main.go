package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"microsvc/pkg/mware"
	"microsvc/svc/stream"
	"microsvc/svc/timeseries"

	"github.com/go-kit/kit/log"
)

var (
	conf       = mware.Config{}
	logger     log.Logger
	httpAddr   string
	mux        *http.ServeMux
	httpLogger log.Logger
)

func init() {
	conf.New()
	httpAddr = conf.GetString("http.address")
	mux = http.NewServeMux()

	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	httpLogger = log.With(logger, "component", "http")
}

func main() {

	ts := timeseries.NewService()

	mux.Handle(conf.GetString("basepath.timeseries"), timeseries.MakeHandler(ts, httpLogger))

	mux.Handle(conf.GetString("basepath.stream"), stream.StreamHandler(conf))

	http.Handle("/", accessControl(mux))

	startServer()

}

func startServer() {
	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(httpAddr, nil)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
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
