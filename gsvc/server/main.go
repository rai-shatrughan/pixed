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
	conf        = &util.Config{}
	logger      = &util.Logger{}
	mux         *http.ServeMux
	wrappedMux  http.Handler
	httpAddr    string
	kfw         = &util.KafkaWriter{}
	kv          = &util.KV{}
	serviceName = os.Getenv("SERVICE_NAME")
	jaegerIP    = os.Getenv("JAEGER_IP")
	mdlw        middleware.Middleware
)

func main() {

	conf.New()
	logger.New()
	kv.New(conf, logger)
	kfw.New(conf, logger)
	httpAddr = conf.GetString("http.address")
	mux = http.NewServeMux()

	mdlw = mware.InitMetricMiddleware()

	tp := mware.InitTracer(jaegerIP, logger)
	defer mware.TracerShutDown(tp, logger)

	appState := util.AppState{
		Mux:         mux,
		Conf:        conf,
		Logger:      logger,
		Kfw:         kfw,
		Kv:          kv,
		ServiceName: serviceName,
	}

	handler.RegisterHandlers(appState)

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
