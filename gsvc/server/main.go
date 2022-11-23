package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	// "gsvc/handler/mc"
	// "gsvc/handler/stream"
	"gsvc/handler"
	"gsvc/mware"
	"gsvc/pkg/util"

	"go.uber.org/zap"

	// "github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/slok/go-http-metrics/middleware"
)

var (
	logger      *zap.Logger
	mux         *http.ServeMux
	wrappedMux  http.Handler
	httpAddr    string
	serviceName = os.Getenv("SERVICE_NAME")
	jaegerIP    = os.Getenv("JAEGER_IP")
	mdlw        middleware.Middleware
)

func main() {

	conf := util.NewConfig()
	logger = util.NewLogger()
	kv := util.NewKV(conf, logger)
	kfw := util.NewKafkaWriter(conf, logger)
	httpAddr = conf.GetString("http.address")
	mux = http.NewServeMux()

	mdlw = mware.InitMetricMiddleware()

	tp := mware.InitTracer(jaegerIP, logger)
	defer mware.TracerShutDown(tp, logger)

	st := &util.AppState{
		Conf:        conf,
		Logger:      logger,
		Kfw:         kfw,
		Kv:          kv,
		ServiceName: &serviceName,
	}
	api := &handler.ApiHandler{}
	api.New(st)

	mux.Handle("/", api)

	wrappedMux = mware.LoggingMiddleware(logger, mux)
	wrappedMux = mware.TracingMiddleware(serviceName, wrappedMux)
	wrappedMux = mware.PrometheusMiddleware(mdlw, wrappedMux)
	wrappedMux = mware.SetAccessControl(wrappedMux)

	startServer()

}

func startServer() {
	errs := make(chan error, 2)
	go func() {
		logger.Info("Server Started on", zap.String("Address", httpAddr))
		errs <- http.ListenAndServe(httpAddr, wrappedMux)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()
	<-errs
	logger.Info("terminated server")
}
