package main

import (
	"context"
	"go-framework-01/grpc_example/proto"
	"log"
)

type Greeter  struct{}

func (greeter *Greeter) HelloService(ctx context.Context, request proto.Request, response *proto.Response) error {
	log.Print("received ")
	response.Msg="hello "+request.Name
	return nil
}
func main() {
	
}
