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
var redisClient *redis.Client
func init()  {
	var err error
	redisClient=redis.NewClient(&redis.Options{
		Addr:"192.168.172.14:6379",
		Password:"",
		DB:0,
		PoolSize: 100,
	})
	pong,err:=redisClient.Ping().Result()
	if err!=nil {
		panic(err)
	}
	fmt.Println("connect result:",pong,err)
}

func main(){
	value := &numberStruct{Number: 42}
	marshalValue, _ :=json.Marshal(value)
	_ = redisClient.Set("key123456", marshalValue, time.Second*3600).Err()
	val, _ := redisClient.Get("key123456").Result()
	fmt.Println("key123456", val)
	_ = redisClient.Set("key123457", marshalValue, time.Second*3600).Err()
	val, _ = redisClient.Get("key123457").Result()
	fmt.Println("key123457", val)
	redisClient.RPush("book","python")
	redisClient.RPush("book","java")

	fmt.Println(redisClient.LRange("book",0,-1).Val())

	//fmt.Println(redisClient.LPop("book"))
	redisClient.SAdd("blacklist","Obama")
	redisClient.SAdd("blacklist","Hillary")
	//fmt.Println(redisClient.Info())

	//模糊删除
	delKey,_:=redisClient.Keys("key12345*").Result();
	//delKey:=[]int{1,2,3,4}

	for i,v:= range delKey{
		fmt.Println(i,"---",v)
	}
	redisClient.Del(delKey...)
	delKey,_=redisClient.Keys("b*").Result();
	//delKey:=[]int{1,2,3,4}

	for i,v:= range delKey{
		fmt.Println(i,"---",v)
	}
	redisClient.Del(delKey...)
}
