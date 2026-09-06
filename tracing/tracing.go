package tracing

import (
	"context"
)

const (
	serviceName = "distribution"

	defaultSamplingRatio = 1

	AttributePrefix = "io.cncf.distribution."
)

func InitOpenTelemetry(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
