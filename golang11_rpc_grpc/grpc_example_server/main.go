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

type EatService struct{}

func (greeter *EatService) BreakfastEat(ctx context.Context, req *greetercenter.BreakfastEatRequestMessage) (*greetercenter.BreakfastEatResponseMessage, error) {
	log.Print("BreakfastEat received ")
	responseMessage := greetercenter.BreakfastEatResponseMessage{Msg: "breakfast success", Code: 200, Data: "hello " + req.Name}
	return &responseMessage, nil
}
func (greeter *EatService) LunchEat(ctx context.Context, req *greetercenter.LunchEatRequestMessage) (*greetercenter.LunchEatResponseMessage, error) {
	log.Print("LunchEat received ")
	responseMessage := greetercenter.LunchEatResponseMessage{Msg: "breakfast success", Code: 200, Data: "hello " + req.Name}
	return &responseMessage, nil
}
func (greeter *EatService) DinnerEat(ctx context.Context, req *greetercenter.DinnerEatRequestMessage) (*greetercenter.DinnerEatResponseMessage, error) {
	log.Print("DinnerEat received ")
	responseMessage := greetercenter.DinnerEatResponseMessage{Msg: "dinner success", Code: 200, Data: "hello " + req.Name}
	return &responseMessage, nil
}
type DrinkService struct{}

func (greeter *DrinkService) BreakfastDrink(ctx context.Context, req *greetercenter.BreakfastDrinkRequestMessage) (*greetercenter.BreakfastDrinkResponseMessage, error) {
	log.Print("BreakfastDrink received ")
	responseMessage := greetercenter.BreakfastDrinkResponseMessage{Msg: "breakfast success", Code: 200, Data: "hello " + req.Name}
	return &responseMessage, nil
}
func (greeter *DrinkService) LunchDrink(ctx context.Context, req *greetercenter.LunchDrinkRequestMessage) (*greetercenter.LunchDrinkResponseMessage, error) {
	log.Print("LunchDrink received ")
	responseMessage := greetercenter.LunchDrinkResponseMessage{Msg: "breakfast success", Code: 200, Data: "hello " + req.Name}
	return &responseMessage, nil
}
func (greeter *DrinkService) DinnerDrink(ctx context.Context, req *greetercenter.DinnerDrinkRequestMessage) (*greetercenter.DinnerDrinkResponseMessage, error) {
	log.Print("DinnerDrink received ")
	responseMessage := greetercenter.DinnerDrinkResponseMessage{Msg: "dinner success", Code: 200, Data: "hello " + req.Name}
	return &responseMessage, nil
}
func main() {
	lis, err := net.Listen("tcp", "10.3.20.236:8083")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	eatService := EatService{}
	drinkService :=DrinkService{}
	greetercenter.RegisterEatServer(s, &eatService)
	greetercenter.RegisterDrinkServer(s,&drinkService)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
