package main

import (
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	jaeger "github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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
			tracer, closer := NewJaegerTracer("apigateway", "127.0.0.1:6831")
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
	r := gin.New()
	r.Use(SetUp())
	r.GET("/testjaeger",httpRequest)
	r.Run(":8080")
}

func httpRequest(context *gin.Context) {
	result, _ :=httpClient("http://10.3.20.236:8082/someGet?one=123456",context)
	context.JSON(http.StatusOK,gin.H{"data":result})
}
func httpClient( url string,context  *gin.Context) (string,error) {
	tr := &http.Transport{
		TLSClientConfig : &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Timeout   : time.Second * 5, //默认5秒超时时间
		Transport : tr,
	}

	req, err := http.NewRequest("GET", url,nil)
	if err != nil {
		return "", err
	}
	tracer, _            := context.Get("Tracer")
	parentSpanContext, _ := context.Get("ParentSpanContext")

	span := opentracing.StartSpan(
		"call Http Get",
		opentracing.ChildOf(parentSpanContext.(opentracing.SpanContext)),
		opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
		ext.SpanKindRPCClient,
	)

	span.Finish()

	injectErr := tracer.(opentracing.Tracer).Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
	if injectErr != nil {
		log.Fatalf("%s: Couldn't inject headers", err)
	}

	resp ,err :=  client.Do(req)
	if err != nil {
		return "", err
	}
	span.Finish();
	content, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return "", err
	}
	return string(content), err
}
