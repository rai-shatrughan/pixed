package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"gsvc/domain/mc"
	"gsvc/domain/mware"
	"gsvc/domain/stream"
	"gsvc/pkg/util"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.uber.org/zap"
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
	url := fmt.Sprintf("http://%s:14268/api/traces", jaegerIP)
	tp, err := mware.InitTracer(url)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

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

	muxRouter.Use(otelmux.Middleware(serviceName))
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
