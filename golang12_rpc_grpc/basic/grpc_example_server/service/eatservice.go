package service

import (
	"context"
	"go-learn-project/golang12_rpc_grpc/basic/greetercenter"
	"log"
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