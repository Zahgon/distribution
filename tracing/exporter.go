package tracing

import (
	"context"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type compositeExporter struct {
	exporters []sdktrace.SpanExporter
}

func newCompositeExporter(exporters ...sdktrace.SpanExporter) *compositeExporter {
	_ = "STUB: not implemented"
	return nil
}

func (ce *compositeExporter) ExportSpans(ctx context.Context, spans []sdktrace.ReadOnlySpan) error {
	_ = "STUB: not implemented"
	return nil
}

func (ce *compositeExporter) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
