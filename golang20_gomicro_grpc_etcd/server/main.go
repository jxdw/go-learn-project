package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"go-learn-project/golang20_gomicro_grpc_etcd/proto"
	"log"
	"strconv"
)
//一个微服务的定义
type Greeter struct {}

//服务的接口实现
func (g *Greeter) Helloservice(ctx context.Context, requestMessage *protocol.RequestMessage,
	                          responseMessage *protocol.ResponseMessage) error {
	responseMessage.Message="hello "+requestMessage.Name+",your age is "+strconv.FormatInt(requestMessage.Age,10)
	return nil
}
func main() {
	registry:=etcdv3.NewRegistry(func(options *registry.Options) {
		options.Addrs=[]string{"192.168.172.5:2379","192.168.172.6:2379","192.168.172.7:2379"}
	})
	service:=micro.NewService(
		micro.Name("greeter.service"),
		micro.Version("1.0.0"),
		micro.Registry(registry),
		)
	service.Init()
	protocol.RegisterGreeterHandler(service.Server(),new(Greeter))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
