// main.go
package main
import (
	"flag"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	otlog "github.com/opentracing/opentracing-go/log"
	"golang.org/x/net/context"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	jaegerClientConfig "github.com/uber/jaeger-client-go/config"
)

var (
	serverPort = flag.String("port", "8000", "server port")
	// 默认为服务模式
	actorKind  = flag.String("actor", "server", "server or client")
)

const (
	server = "server"
	client = "client"
)

func main() {
	flag.Parse()

	if *actorKind != server && *actorKind != client {
		log.Fatal("Please specify '-actor server' or '-actor client'")
	}

	cfg := jaegerClientConfig.Configuration{
		Sampler: &jaegerClientConfig.SamplerConfig{
			Type:  "const",
			Param: 1.0, // sample all traces
		},
	}
	// jaeger.NewRemoteReporter(transport)
	tracer, closer, _ := cfg.New(*actorKind)
	defer closer.Close()

	if *actorKind == server {
		runServer(tracer)
		return
	}

	runClient(tracer)

	// Close the tracer to guarantee that all spans that could
	// be still buffered in memory are sent to the tracing backend
	closer.Close()
}

func getTime(w http.ResponseWriter, r *http.Request) {
	log.Print("Received getTime request")
	t := time.Now()
	ts := t.Format("Mon Jan _2 15:04:05 2006")
	io.WriteString(w, fmt.Sprintf("The time is %s", ts))
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r,
		fmt.Sprintf("http://localhost:%s/gettime", *serverPort), 301)
}

func runServer(tracer opentracing.Tracer) {
	http.HandleFunc("/gettime", getTime)
	http.HandleFunc("/", redirect)
	log.Printf("Starting server on port %s", *serverPort)
	err := http.ListenAndServe(
		fmt.Sprintf(":%s", *serverPort),
		// use nethttp.Middleware to enable OpenTracing for server
		nethttp.Middleware(tracer, http.DefaultServeMux))
	if err != nil {
		log.Fatalf("Cannot start server: %s", err)
	}
}
func runClient(tracer opentracing.Tracer) {
	// nethttp.Transport from go-stdlib will do the tracing
	c := &http.Client{Transport: &nethttp.Transport{}}

	// create a top-level span to represent full work of the client
	span := tracer.StartSpan(client)
	span.SetTag(string(ext.Component), client)
	defer span.Finish()
	ctx := opentracing.ContextWithSpan(context.Background(), span)

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("http://localhost:%s/", *serverPort),
		nil,
	)
	if err != nil {
		onError(span, err)
		return
	}

	req = req.WithContext(ctx)
	// wrap the request in nethttp.TraceRequest
	req, ht := nethttp.TraceRequest(tracer, req)
	defer ht.Finish()

	res, err := c.Do(req)
	if err != nil {
		onError(span, err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		onError(span, err)
		return
	}
	fmt.Printf("Received result: %s\n", string(body))
}

func onError(span opentracing.Span, err error) {
	// handle errors by recording them in the span
	span.SetTag(string(ext.Error), true)
	span.LogKV(otlog.Error(err))
	log.Print(err)
}