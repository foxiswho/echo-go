package model

import (
	"github.com/opentracing/opentracing-go"

	"github.com/foxiswho/echo-go/module/log"
)

type Model struct {
	Context context
}

func (m Model) Trace() opentracing.Span {
	if m.Context != nil {
		if span := m.Context.OpenTracingSpan(); span != nil {
			comp := "orm"
			s := opentracing.StartSpan(comp+":GetUserById", opentracing.ChildOf(span.Context()))
			s.SetTag("component", comp)
			s.SetTag("span.kind", "server")

			return s
		} else {
			log.Debugf("trace faile context span nil")
		}
	} else {
		log.Debugf("trace faile orm context nil")
	}
	return nil
}
