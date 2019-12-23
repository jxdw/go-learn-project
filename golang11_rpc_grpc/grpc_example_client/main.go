package main

import (
	"context"
	protocol "go-learn-code/golang10_rpc_grpc/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("10.3.20.223:8083", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := protocol.NewGreeterClient(conn)

	ctx,cancel:=context.WithTimeout(context.Background(),time.Second*200)
	defer cancel()

	response,err:=c.Helloservice(ctx,&protocol.RequestMessage{Name:"xiaogang"})
	if err!=nil {
		log.Fatalln("could not  greet :%v",err)
	}
	log.Printf("greeting: %s",response.Msg)
}
