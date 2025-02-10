package mtl

import (
	"github.com/kitex-contrib/obs-opentelemetry/provider"
)

// InitTracing 初始化链路追踪
func InitTracing(serviceName string) provider.OtelProvider {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithInsecure(),
		provider.WithEnableMetrics(false),
	)
	return p
}
