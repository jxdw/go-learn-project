package main

import (
	"context"
	"go-learn-project/01_middleware_sdk/15_rpc_framework/basic/greetercenter"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("10.3.20.215:8083", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := greetercenter.NewEatClient(conn)

	ctx,cancel:=context.WithTimeout(context.Background(),time.Second*200)
	defer cancel()

	response,err:=c.BreakfastEat(ctx,&greetercenter.BreakfastEatRequestMessage{Name: "xiaogang"})
	if err!=nil {
		log.Fatalln("could not  greet :%v",err)
	}
	log.Printf("greeting: %s,data is %s",response.Msg,response.Data)
}
