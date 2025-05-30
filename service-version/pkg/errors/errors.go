package errors

import (
	"context"
	"fmt"
	"runtime/debug"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type internalError struct {
	Code            int    `json:"-"`
	Message         string `json:"message"`
	InternalMessage string `json:"-"`
	Err             error  `json:"-"`
}

func NewInternalError(ctx context.Context, Code int, Message string, InternalMessage string, Err error) *internalError {
	err := &internalError{
		Code:            Code,
		Message:         Message,
		InternalMessage: InternalMessage,
		Err:             Err,
	}
	return err.appendCurrentSpan(ctx)
}

func (i *internalError) appendCurrentSpan(ctx context.Context) *internalError {
	span := trace.SpanFromContext(ctx)
	if span == nil || !span.IsRecording() {
		return i
	}

	attrs := []attribute.KeyValue{
		attribute.String("error.message", i.Message),
		attribute.String("exception.type", fmt.Sprintf("%T", i.Err)),
		attribute.String("error.internal_message", i.InternalMessage),
		attribute.String("exception.stacktrace", string(debug.Stack())),
	}

	span.SetStatus(codes.Error, i.InternalMessage)
	span.AddEvent("error", trace.WithAttributes(attrs...))
	return i
}
