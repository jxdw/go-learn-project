package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type numberStruct struct {
	Number int
}

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
	value := &numberStruct{Number: 42}
	marshalValue,err:=json.Marshal(value)
	err = client.Set("key123456", marshalValue, time.Second*3600).Err()
	val, err := client.Get("key123456").Result()
	fmt.Println("key123456", val)

}
