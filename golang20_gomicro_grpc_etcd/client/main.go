package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	protocol "go-learn-project/golang20_gomicro_grpc_etcd/proto"
	"log"
)

func main() {
	registry:=etcdv3.NewRegistry(func(options *registry.Options) {
		options.Addrs=[]string{"192.168.172.5:2379","192.168.172.6:2379","192.168.172.7:2379"}
	})
	service:=micro.NewService(
		micro.Registry(registry),
	)
	service.Init()
	serviceClient:=protocol.NewGreeterService("greeter.service",service.Client())
	responseMessage,err:= serviceClient.Helloservice(context.Background(),&protocol.RequestMessage{Name:"xiaogang",Age:18})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(responseMessage)
}
