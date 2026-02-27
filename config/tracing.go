// Copyright © 2025 Ping Identity Corporation

package config

import (
	"context"

	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

// tracerInstrumentationScope is the OpenTelemetry instrumentation scope name used when
// creating spans in the config package. It follows the convention of using the module path.
const tracerInstrumentationScope = "github.com/pingidentity/pingone-go-client"

// noopConfigSpan is a reusable no-op span returned by startConfigSpan when no
// TracerProvider has been configured. Calling End, SetStatus, RecordError, or AddEvent
// on this span is always safe and has no observable effect.
var noopConfigSpan trace.Span = noop.Span{}

// startConfigSpan starts a new tracing span for the given config-layer operation.
// The ctx parameter provides the parent context and any active parent span for
// trace propagation. The name parameter should describe the operation in the form
// "pingone.config.OperationName". Additional span start options such as span kind
// or initial attributes may be supplied via opts.
//
// If no TracerProvider has been configured on the Configuration (i.e. TracerProvider
// is nil), this function returns the original context together with a no-op span so
// that callers require no nil-checks. When tracing is enabled the returned context
// carries the new span as the active span.
func (c *Configuration) startConfigSpan(ctx context.Context, name string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	if c.TracerProvider == nil {
		return ctx, noopConfigSpan
	}
	return c.TracerProvider.Tracer(tracerInstrumentationScope).Start(ctx, name, opts...)
}

// recordConfigSpanError marks span as failed when err is non-nil. It is a no-op
// when err is nil or when the span is a no-op span (i.e. tracing is disabled).
func recordConfigSpanError(span trace.Span, err error) {
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)
	}
}
