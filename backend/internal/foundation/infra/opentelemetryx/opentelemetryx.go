package opentelemetryx

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.40.0"
	"google.golang.org/grpc"
	oteltracing "google.golang.org/grpc/experimental/opentelemetry"
	"google.golang.org/grpc/stats/opentelemetry"
)

func NewOpenTelemetry() grpc.ServerOption {

	meterExporter, err := prometheus.New()

	if err != nil {
		panic(err)
	}

	meterProvider := metric.NewMeterProvider(metric.WithReader(meterExporter))

	traceExporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())

	if err != nil {
		panic(err)
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter),
		trace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String("ozma-foundation"),
			),
		),
	)

	textMapPropagator := propagation.TraceContext{}

	so := opentelemetry.ServerOption(
		opentelemetry.Options{
			MetricsOptions: opentelemetry.MetricsOptions{
				MeterProvider: meterProvider,
				Metrics: opentelemetry.DefaultMetrics().Add(
					"grpc.lb.pick_first.connection_attempts_succeeded",
					"grpc.lb.pick_first.connection_attempts_failed",
				),
			},
			TraceOptions: oteltracing.TraceOptions{
				TracerProvider:    traceProvider,
				TextMapPropagator: textMapPropagator,
			},
		},
	)
	go http.ListenAndServe(":9090", promhttp.Handler())

	return so

}
