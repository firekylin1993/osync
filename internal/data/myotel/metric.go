package myotel

import (
	"context"
	"osync/internal/conf"
	"time"

	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/sdk/metric/export"
	"google.golang.org/grpc"
)

func NewMetricClient(c *conf.Server) otlpmetric.Client {
	if c == nil || c.Otel == nil || c.Otel.Addr == "" {
		return nil
	}
	return otlpmetricgrpc.NewClient(
		otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithEndpoint(c.Otel.Addr),
		otlpmetricgrpc.WithDialOption(grpc.WithBlock()))
}

func NewMetricExporter(ctx context.Context, client otlpmetric.Client) export.Exporter {
	if client == nil {
		return nil
	}
	c, cancelFunc := context.WithTimeout(ctx, time.Second)
	defer cancelFunc()
	metricExp, err := otlpmetric.New(c, client)
	if err != nil {
		panic(errors.WithMessage(err, "初始化 metric exporter 失败"))
	}
	return metricExp
}
