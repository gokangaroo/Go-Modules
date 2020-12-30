package gorm

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
)

func initJaeger() (closer io.Closer, err error) {
	configuration := config.Configuration{
		ServiceName: "gormTracing",
		Disabled:    false,
		Sampler: &config.SamplerConfig{
			Type: jaeger.SamplerTypeConst,
			// param的值在0到1之间，设置为1则将所有的Operation输出到Reporter
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "localhost:6831",
		},
	}

	tracer, closer, err := (&configuration).NewTracer()
	if err != nil {
		return
	}

	// sets the [singleton] opentracing
	// call as early as possible
	opentracing.SetGlobalTracer(tracer)
	return closer, err
}
