package n3tracing

import (
	"context"
	"net/http"
	"testing"

	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
)

type TrObj struct {
	tracer opentracing.Tracer
	ctx    context.Context
}

func (to *TrObj) SetTracer(t opentracing.Tracer) {
	to.tracer = t
}

func (to *TrObj) Tracer() opentracing.Tracer {
	return to.tracer
}

func (to *TrObj) SetContext(ctx context.Context) {
	to.ctx = ctx
}

func (to *TrObj) Context() context.Context {
	return to.ctx
}

// docker run -d --name jaeger -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p 5775:5775/udp -p 6831:6831/udp -p 6832:6832/udp -p 5778:5778 -p 16686:16686 -p 14268:14268 -p 14250:14250 -p 9411:9411 jaegertracing/all-in-one:1.18

func TestInitTracer(t *testing.T) {
	// os.Setenv("JAEGER_SERVICE_NAME", "n3tracing-test")
	// os.Setenv("JAEGER_SAMPLER_TYPE", "const")
	// os.Setenv("JAEGER_SAMPLER_PARAM", "1")

	e := echo.New()
	defer e.Close()
	defer e.Start(":1500")

	// Add Jaeger Tracer into Middleware
	c := jaegertracing.New(e, nil)
	defer c.Close()

	obj := TrObj{InitTracer("n3tracing-test"), nil}

	// e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	// 	return func(c echo.Context) error {
	// 		obj.SetContext(c.Request().Context())
	// 		return next(c)
	// 	}
	// })

	e.GET("/", func(c echo.Context) error {
		obj.SetContext(c.Request().Context())
		DoTracing(&obj, "testOperName", "testSpanValue", "testTAG", "testTAGValue", "testEvent")
		return c.String(http.StatusOK, "test")
	})
}
