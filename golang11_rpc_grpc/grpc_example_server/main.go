package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	greetercenter "go-learn-project/golang11_rpc_grpc/greetercenter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Greeter struct{}

func (greeter *Greeter) Breakfast(ctx context.Context, req *greetercenter.BreakfastRequestMessage) (*greetercenter.BreakfastResponseMessage, error) {
	log.Print("received ")
	//conn, err :=dbServer.connect(ctx)
	//if err!=nil {
	//
	//}
	//conn.QueryContext(ctx,"SELECT `ID`, `Title`, `Description`, `Reminder` FROM ToDo WHERE `ID`=?",)
	responseMessage := greetercenter.BreakfastResponseMessage{Msg: "breakfast success", Code: 200, Data: "hello " + req.Name}
	return &responseMessage, nil
}
func (greeter *Greeter) Lunch(ctx context.Context, req *greetercenter.LunchRequestMessage) (*greetercenter.LunchResponseMessage, error) {
	log.Print("received ")
	//conn, err :=dbServer.connect(ctx)
	//if err!=nil {
	//
	//}
	//conn.QueryContext(ctx,"SELECT `ID`, `Title`, `Description`, `Reminder` FROM ToDo WHERE `ID`=?",)
	responseMessage := greetercenter.LunchResponseMessage{Msg: "breakfast success", Code: 200, Data: "hello " + req.Name}
	return &responseMessage, nil
}
func (greeter *Greeter) Dinner(ctx context.Context, req *greetercenter.DinnerRequestMessage) (*greetercenter.DinnerResponseMessage, error) {
	log.Print("received ")
	//conn, err :=dbServer.connect(ctx)
	//if err!=nil {
	//
	//}
	//conn.QueryContext(ctx,"SELECT `ID`, `Title`, `Description`, `Reminder` FROM ToDo WHERE `ID`=?",)
	responseMessage := greetercenter.DinnerResponseMessage{Msg: "dinner success", Code: 200, Data: "hello " + req.Name}
	return &responseMessage, nil
}
func main() {
	lis, err := net.Listen("tcp", "10.3.20.236:8083")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	greeter := Greeter{}
	greetercenter.RegisterEatServer(s, &greeter)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
