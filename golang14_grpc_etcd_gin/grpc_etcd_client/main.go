package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	protocol "go-learn-project/golang14_grpc_etcd_gin/proto"
	"go-learn-project/golang14_grpc_etcd_gin/registercenter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"net/http"

	"log"
	"time"
)

type RequestData struct {
	MethodName string `form:"methodName" json:"methodName" xml:"methodName"`
	Data string `form:"data" json:"data" xml:"data"`
}
var r	resolver.Builder

func init() {

}
func main() {
	router := gin.Default()
	//这里的://authority可以随意写但是需要://开头
	//conn, err := grpc.Dial(r.Scheme()+"/GreeterService", grpc.WithBalancerName("round_robin"),grpc.WithInsecure())
	router.POST("/c", func(c *gin.Context) {
        var requestData RequestData
		var err error
		contentType := c.Request.Header.Get("Content-Type")
		switch contentType {
		case "application/json":
			err = c.BindJSON(&requestData)
		case "application/x-www-form-urlencoded":
			err = c.BindWith(&requestData, binding.Form)
		}
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			  "returnData": newGreeterClient(requestData),
		})
	})
	router.Run(":8000")
}
func newGreeterClient(data RequestData)  string {
	EtcdAddr:="192.168.172.5:2379;192.168.172.6:2379;192.168.172.7:2379"
	//EtcdAddr:="127.0.0.1:2379"
	r = registercenter.NewResolver(EtcdAddr)
	resolver.Register(r)
	conn, err := grpc.Dial(r.Scheme()+"://abc/"+data.MethodName, grpc.WithBalancerName("round_robin"),grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := protocol.NewGreeterClient(conn)

	ctx,cancel:=context.WithTimeout(context.Background(),time.Second*30)
	defer cancel()

	response,err:=c.Helloservice(ctx,&protocol.RequestMessage{Name:data.Data})
	if err!=nil {
		log.Fatalln("could not  greet :%v",err)
	}
	fmt.Printf("greeting: %s",response.Msg)
	return response.Msg
}
