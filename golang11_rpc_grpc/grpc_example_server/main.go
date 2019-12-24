package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"go-learn-project/golang11_rpc_grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Greeter  struct{}

func (greeter *Greeter) Helloservice(ctx context.Context, req *protocol.HelloRequestMessage) (*protocol.HelloResponseMessage, error) {
	log.Print("received ")
	//conn, err :=dbServer.connect(ctx)
	//if err!=nil {
	//
	//}
	//conn.QueryContext(ctx,"SELECT `ID`, `Title`, `Description`, `Reminder` FROM ToDo WHERE `ID`=?",)
	responseMessage:=protocol.HelloResponseMessage{Msg:"success",Code:200,Data:"hello "+req.Name}
	return &responseMessage,nil
}
func main() {
	lis,err:=net.Listen("tcp","10.3.20.236:8083")
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
