package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main(){
	client:=redis.NewClient(&redis.Options{
		Addr:"192.168.172.4:6380",
		Password:"",
		DB:0,
	})
	pong,err:=client.Ping().Result()
	if err!=nil {
		panic(err)
	}
	fmt.Println("connect result:",pong,err)

}
