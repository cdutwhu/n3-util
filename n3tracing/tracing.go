package n3tracing

import (
	"context"

	"github.com/opentracing/opentracing-go"
	tags "github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go/config"
)

// ITrace :
type ITrace interface {
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
	ctx, tracer := it.Context(), it.Tracer()
	failP1OnErrWhen(ctx == nil, "%v", fEf("Need 'SetContext'"))
	failP1OnErrWhen(tracer == nil, "%v", fEf("Need 'SetTracer'"))
	span := opentracing.SpanFromContext(ctx)
	failP1OnErrWhen(span == nil, "%v", fEf("Need 'jaegertracing.New(e, nil)'"))
	// fPln(" ------- DoTracing ------- ")
	logger("DoTracing: %s", event)
	span = tracer.StartSpan(operName, opentracing.ChildOf(span.Context()))
	defer span.Finish()
	span.LogEvent(event)
	span.SetTag(tagKey, tagValue)
	tags.SpanKindRPCClient.Set(span)
	tags.PeerService.Set(span, spanValue)
	opentracing.ContextWithSpan(ctx, span)
}
