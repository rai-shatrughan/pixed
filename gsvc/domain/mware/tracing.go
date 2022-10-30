package mware

import (
	"context"
	"fmt"
	"gsvc/pkg/util"
	"net/http"
	"os"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/otel"

	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

var (
	serviceName = os.Getenv("SERVICE_NAME")
)

func InitTracer(JAEGER_IP string, logger *util.Logger) *sdktrace.TracerProvider {
	// exporter, err := stdout.New(stdout.WithPrettyPrint())
	url := fmt.Sprintf("http://%s:14268/api/traces", JAEGER_IP)
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		logger.Sugar().Error(err)
		return nil
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.TraceIDRatioBased(0.001)),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp
}

func TracerShutDown(tp *sdktrace.TracerProvider, logger *util.Logger) func() {
	return func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			logger.Sugar().Error(err)
		}
	}
}

func TracingMiddleware(serviceName string) func(http.Handler) http.Handler {
	return otelmux.Middleware(serviceName)

}
