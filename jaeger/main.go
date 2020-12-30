package main

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"log"
	"time"
)

func main() {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "localhost:6831",
		},
	}

	closer, err := cfg.InitGlobalTracer(
		"serviceName",
	)
	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s\n", err.Error())
		return
	}
	defer closer.Close()

	var ctx = context.TODO()
	span1, ctx := opentracing.StartSpanFromContext(ctx, "span_1")
	time.Sleep(time.Second / 2)

	span11, _ := opentracing.StartSpanFromContext(ctx, "span_1-1")
	time.Sleep(time.Second / 2)
	span11.Finish()

	span1.Finish()
}
