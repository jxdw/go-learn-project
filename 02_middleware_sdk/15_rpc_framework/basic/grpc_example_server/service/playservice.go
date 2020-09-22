package service

import (
	"context"
	"go-learn-project/02_middleware_sdk/15_rpc_framework/basic/greetercenter"
)

type PlayService struct {}

func (playService *PlayService)  BreakfastPlay(ctx context.Context, requestmessage *greetercenter.BreakfastPlayRequestMessage) (*greetercenter.BreakfastPlayResponseMessage, error) {
	responseMessage := greetercenter.BreakfastPlayResponseMessage{Msg: "breakfast success", Code: 200, Data: "hello " + requestmessage.Name}
	return &responseMessage,nil
}


func (playService *PlayService)  LunchPlay(ctx context.Context, requestmessage *greetercenter.LunchPlayRequestMessage) (*greetercenter.LunchPlayResponseMessage, error) {
	responseMessage := greetercenter.LunchPlayResponseMessage{Msg: "breakfast success", Code: 200, Data: "hello " + requestmessage.Name}
	return &responseMessage,nil
}

func (playService *PlayService)  DinnerPlay(ctx context.Context, requestmessage *greetercenter.DinnerPlayRequestMessage) (*greetercenter.DinnerPlayResponseMessage, error) {
	responseMessage := greetercenter.DinnerPlayResponseMessage{Msg: "breakfast success", Code: 200, Data: "hello " + requestmessage.Name}
	return &responseMessage,nil
}
