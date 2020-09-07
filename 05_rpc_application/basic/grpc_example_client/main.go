package main

import (
	"context"
	"go-learn-project/05_rpc_application/basic/greetercenter"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("10.3.20.236:12345", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := greetercenter.NewEatClient(conn)

	ctx,cancel:=context.WithTimeout(context.Background(),time.Second*200)
	defer cancel()

	response,err:=c.BreakfastEat(ctx,&greetercenter.BreakfastEatRequestMessage{Name:"xiaogang"})
	if err!=nil {
		log.Fatalln("could not  greet :%v",err)
	}
	log.Printf("greeting: %s",response.Msg)
}
