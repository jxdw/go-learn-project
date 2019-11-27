package main

import (
	"context"
	"fmt"
	protocol "go-framework-01/google_grpc_basic_example/proto"
	"go-framework-01/google_grpc_etcd_example/registercenter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"log"
	"time"
)

func main() {
	EtcdAddr:="192.168.172.5:2379;192.168.172.6:2379;192.168.172.7:2379"
	r := registercenter.NewResolver(EtcdAddr)
	resolver.Register(r)
	conn, err := grpc.Dial(r.Scheme()+"://author/GreeterService", grpc.WithBalancerName("round_robin"),grpc.WithInsecure())
	//conn, err := grpc.Dial(r.Scheme()+"/GreeterService", grpc.WithBalancerName("round_robin"),grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := protocol.NewGreeterClient(conn)

	ctx,cancel:=context.WithTimeout(context.Background(),time.Second*30)
	defer cancel()

	response,err:=c.Helloservice(ctx,&protocol.RequestMessage{Name:"xiaogang"})
	if err!=nil {
		log.Fatalln("could not  greet :%v",err)
	}
	fmt.Printf("greeting: %s",response.Msg)
}
