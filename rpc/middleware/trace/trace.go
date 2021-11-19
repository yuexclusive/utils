package trace

import (
	"context"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"github.com/yuexclusive/utils/config"
)

func Tracer() (opentracing.Tracer, io.Closer, error) {
	cfg := config.MustGet()
	jcfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		ServiceName: cfg.Name,
	}

	report := jaegercfg.ReporterConfig{
		LogSpans:           true,
		LocalAgentHostPort: cfg.JaegerAddress,
		QueueSize:          1000,
	}

	reporter, _ := report.NewReporter(cfg.Name, jaeger.NewNullMetrics(), jaeger.NullLogger)
	return jcfg.NewTracer(
		jaegercfg.Reporter(reporter),
	)
}

type SpanOption func(span opentracing.Span)

func SpanWithError(err error) SpanOption {
	return func(span opentracing.Span) {
		if err != nil {
			ext.Error.Set(span, true)
			span.LogFields(log.String("event", "error"), log.String("msg", err.Error()))
		}
	}
}

// example:
// SpanWithLog(
//    "event", "soft error",
//    "type", "cache timeout",
//    "waited.millis", 1500)
func SpanWithLog(arg ...interface{}) SpanOption {
	return func(span opentracing.Span) {
		span.LogKV(arg...)
	}
}

func Start(tracer opentracing.Tracer, spanName string, ctx context.Context) (newCtx context.Context, finish func(...SpanOption)) {
	if ctx == nil {
		ctx = context.TODO()
	}
	span, newCtx := opentracing.StartSpanFromContextWithTracer(ctx, tracer, spanName,
		opentracing.Tag{Key: string(ext.Component), Value: "func"},
	)

	finish = func(ops ...SpanOption) {
		for _, o := range ops {
			o(span)
		}
		span.Finish()
	}

	return
}
