package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/etcd"
	"go-framework-01/web/controller"
	"log"
	"net/http"
)

func main() {
	registry:=etcd.NewRegistry(func(option *registry.Options) {
		option.Addrs=[]string{"127.0.0.1:2379"}
	})
	service := web.NewService(
		web.Name("go.micro.api.login"),
		web.Registry(registry),
	)
	service.Init(func(o *web.Options){
		o.Address="127.0.0.1:8081"
	},
	func(o *web.Options){
		log.Print("second init")
	})
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
	if err := service.Run(); err != nil {log.Fatal(err)}
}
