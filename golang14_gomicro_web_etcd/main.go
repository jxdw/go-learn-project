package main

import (

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/etcdv3"
	"go-framework-01/golang14_gomicro_web_etcd/controller"
	"log"
	"net/http"
)


func main() {
	registry:=etcdv3.NewRegistry(func(options *registry.Options) {
		options.Addrs=[]string{"192.168.172.5:2379","192.168.172.6:2379","192.168.172.7:2379"}
	})
	service := web.NewService(
		web.Name("go.micro.api.login"),
		web.Registry(registry),
	)

	service.Init()
	user := new(controller.User)
	router := gin.Default()
	router.HandleMethodNotAllowed = true
	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"result": false, "error": "Method Not Allowed"})
		return
	})
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"result": false, "error": "Endpoint Not Found"})
		return
	})
	router.POST("/login", user.Login)
	service.Handle("/", router)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
