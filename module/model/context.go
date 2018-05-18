package model

import (
	"github.com/opentracing/opentracing-go"
)

type context interface {
	OpenTracingSpan() opentracing.Span
}
