package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"gsvc/handler"
	"gsvc/mware"
	"gsvc/pkg/util"

	"github.com/gorilla/mux"

	"go.uber.org/zap"

	"github.com/slok/go-http-metrics/middleware"
)

var (
	conf        util.Config
	logger      util.Logger
	httpAddr    string
	muxRouter   *mux.Router
	kf          util.KafkaWriter
	kv          util.KV
	serviceName = os.Getenv("SERVICE_NAME")
	jaegerIP    = os.Getenv("JAEGER_IP")
	mdlw        middleware.Middleware
)

func init() {
	conf.New()
	logger.New()

	kv.New(&conf, &logger)
	kf.New(&conf, &logger)

	httpAddr = conf.GetString("http.address")
	muxRouter = mux.NewRouter()

	mdlw = mware.InitMetricMiddleware()
}

func main() {
	tp := mware.InitTracer(jaegerIP, &logger)
	defer mware.TracerShutDown(tp, &logger)

	hndlr := handler.SVCHandlers{
		MuxRouter: muxRouter,
		Conf:      &conf,
		Logger:    &logger,
		Kf:        &kf,
		Kv:        &kv,
	}

	hndlr.RegisterHandlers()

	muxRouter.Use(mware.TracingMiddleware(serviceName))
	muxRouter.Use(mware.LoggingMiddleware(&logger))
	muxRouter.Use(mware.PrometheusMiddleware(mdlw))
	muxRouter.Use(mware.SetAccessControl)

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
