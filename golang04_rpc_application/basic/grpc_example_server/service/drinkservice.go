package service

import (
	"context"
	"go-learn-project/golang04_rpc_application/basic/greetercenter"
	"log"
)

/**
实现饮品服务
*/
type DrinkService struct{}
/**
实现饮品服务的早餐饮品接口
*/
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
