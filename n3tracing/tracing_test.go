package n3tracing

// import (
// 	"context"
// 	"net/http"
// 	"testing"

// 	"github.com/labstack/echo/v4"
// 	"github.com/opentracing/opentracing-go"
// )

// type TrObj struct {
// 	msg    string
// 	tracer opentracing.Tracer
// 	ctx    context.Context
// }

// func (to *TrObj) SetTracer(tracer opentracing.Tracer) {
// 	to.tracer = tracer
// }

// func (to *TrObj) GetTracer() opentracing.Tracer {
// 	return to.tracer
// }

// func (to *TrObj) SetContext(ctx context.Context) {
// 	to.ctx = ctx
// }

// func (to *TrObj) GetContext() context.Context {
// 	return to.ctx
// }

// // docker run -d --name jaeger -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p 5775:5775/udp -p 6831:6831/udp -p 6832:6832/udp -p 5778:5778 -p 16686:16686 -p 14268:14268 -p 14250:14250 -p 9411:9411 jaegertracing/all-in-one:1.18

// func TestInitTracer(t *testing.T) {
// 	e := echo.New()
// 	defer e.Close()
// 	defer e.Start(":1500")

// 	e.GET("/", func(c echo.Context) error {
// 		obj := &TrObj{"TrObj", InitTracer("n3tracing-test"), c.Request().Context()}
// 		DoTracing(obj, "testOperName", "testSpanValue", "testTAG", "testTAGValue", "testEvent")
// 		return c.String(http.StatusOK, "test")
// 	})
// }
