package main

import ("context"
greetterService "go-framework-01/grpc_example/grpc_example_proto/protocol"
	"log"
)

type Greeter  struct{}

func (greeter *Greeter) HelloService(ctx context.Context, request greetterService.Request, response *greetterService.Response) error {
	log.Print("received ")
	response.Msg="hello "+request.Name
	return nil
}
func main() {
	
}
