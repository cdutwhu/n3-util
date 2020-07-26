package n3tracing

import (
	"context"

	"github.com/opentracing/opentracing-go"
	tags "github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go/config"
)

// ITrace :
type ITrace interface {
	SetTracer(opentracing.Tracer)
	Tracer() opentracing.Tracer
	SetContext(context.Context)
	Context() context.Context
}

// InitTracer :
func InitTracer(service string) opentracing.Tracer {
	cfg, err := config.FromEnv()
	failOnErr("%v: ", err)
	cfg.ServiceName = service
	cfg.Sampler.Type = "const"
	cfg.Sampler.Param = 1
	tracer, _, err := cfg.NewTracer()
	failOnErr("%v: ", err)
	return tracer
}

// DoTracing :
func DoTracing(it ITrace, operName, spanValue, tagKey, tagValue, event string) {
	if ctx := it.Context(); ctx != nil {
		if span := opentracing.SpanFromContext(ctx); span != nil {
			if tracer := it.Tracer(); tracer != nil {
				// fPln(" ------- DoTracing ------- ")
				logger("DoTracing: %s", event)
				span := tracer.StartSpan(operName, opentracing.ChildOf(span.Context()))
				defer span.Finish()
				span.LogEvent(event)
				span.SetTag(tagKey, tagValue)
				tags.SpanKindRPCClient.Set(span)
				tags.PeerService.Set(span, spanValue)
				ctx = opentracing.ContextWithSpan(ctx, span)
			}
		}
	}
}
