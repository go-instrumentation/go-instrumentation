package jaeger

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"os"
	"strconv"
)

var Context context.Context
var IsGlobalRegistered bool
var MaxTagLength int

func init() {
	env := os.Getenv("JAEGER_MAX_TAG_LENGTH")
	if env != "" {
		var err error
		MaxTagLength, err = strconv.Atoi(env)
		awesome_error.CheckFatal(err)
	}
}

func initTracer(service string) (tracer opentracing.Tracer, closer io.Closer) {
	cfg, err := config.FromEnv()
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	if cfg.ServiceName == "" {
		cfg.ServiceName = service
	}
	if cfg.Sampler != nil {
		if cfg.Sampler.Param == 0 {
			cfg.Sampler.Param = 1
		}
	}
	var options []config.Option
	if cfg.Reporter.LogSpans {
		options = append(options, config.Logger(jaeger.StdLogger))
	}
	if MaxTagLength > 0 {
		options = append(options, config.MaxTagValueLength(MaxTagLength))
	}
	tracer, closer, err = cfg.NewTracer(options...)
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)
	IsGlobalRegistered = true
	return
}

func Span(operationName string, serviceName ...string) (span opentracing.Span, closer io.Closer) {
	if IsGlobalRegistered {
		span, _ = opentracing.StartSpanFromContext(Context, operationName)
		return
	}
	service := "trace"
	if len(serviceName) > 0 {
		if len(serviceName[0]) > 0 {
			service = serviceName[0]
		}
	}
	tracer, closer := initTracer(service)
	span = tracer.StartSpan(operationName)
	Context = opentracing.ContextWithSpan(context.Background(), span)
	return
}
