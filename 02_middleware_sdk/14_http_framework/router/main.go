package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	_ "go-learn-project/02_middleware_sdk/14_http_framework/router/docs"
	"io"
	"net/http"
)
func NewJaegerTracer(serviceName string, jaegerHostPort string) (opentracing.Tracer, io.Closer) {
	cfg := &jaegerConfig.Configuration {
		Sampler: &jaegerConfig.SamplerConfig{
			Type  : "const", //固定采样
			Param : 1,       //1=全采样、0=不采样
		},

		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans           : true,
			LocalAgentHostPort : jaegerHostPort,
		},
		ServiceName: serviceName,
	}

	tracer, closer, err := cfg.NewTracer(jaegerConfig.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}
func SetUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var parentSpan opentracing.Span
		tracer, closer := NewJaegerTracer("gin server", "127.0.0.1:6831")
		defer closer.Close()
		spCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		if err != nil {
			parentSpan = tracer.StartSpan(c.Request.URL.Path)
			defer parentSpan.Finish()
		} else {
			parentSpan = opentracing.StartSpan(
				c.Request.URL.Path,
				opentracing.ChildOf(spCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
				ext.SpanKindRPCServer,
			)
			defer parentSpan.Finish()
		}
		c.Set("Tracer", tracer)
		c.Set("ParentSpanContext", parentSpan.Context())
		c.Next()
	}
}
func main() {
	engine:=gin.Default()
	engine.Use(SetUp())
	engine.GET("/someGet",getting)
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.Group("/api/v1")
	engine.POST("/postSome",posting)
	engine.Run(":8082")
}

func posting(context *gin.Context) {
	buf:=bytes.Buffer{}
	buf.ReadFrom(context.Request.Body)
	context.Request.Body.Close();
	s:=buf.String();
	context.String(http.StatusOK,"request is:"+s);
}

// @Summary 获取查询的信息
// @Produce  json
// @Param one path string true "查询的信息"
// @Success 200 {string} string "{"status":"OK","message":"ok"}"
// @Router /someGet?one={one} [get]
func getting(context *gin.Context) {
	msg:=context.Query("one");
	context.JSON(http.StatusOK,gin.H{"status":"OK","message":msg})
}
