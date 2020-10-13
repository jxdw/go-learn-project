package main

import (
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go-learn-project/01_middleware_sdk/11_trace_jaeger/jaeger_gin_example/util"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.New()
	r.Use(util.SetUp("api gateway"))
	r.GET("/testjaeger",httpRequest)
	r.Run(":8080")
}

func httpRequest(context *gin.Context) {
	result, _ :=httpClient("http://10.3.20.215:8082/someGet?one=123456",context)
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
