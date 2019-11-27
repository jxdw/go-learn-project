package main

import (
	"context"
	"go-framework-01/google_grpc_example/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Greeter  struct{}

func (greeter *Greeter) Helloservice(ctx context.Context, req *protocol.RequestMessage) (*protocol.ResponseMessage, error) {
	log.Print("received ")
	responseMessage:=protocol.ResponseMessage{Msg:"hello "+req.Name}
	return &responseMessage,nil
}
func main() {
	lis,err:=net.Listen("tcp","10.3.20.57:8083")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s:=grpc.NewServer()
	greeter:=Greeter{}
	protocol.RegisterGreeterServer(s,&greeter)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
