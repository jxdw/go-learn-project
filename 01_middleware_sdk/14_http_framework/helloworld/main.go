package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	engine:=gin.Default()
	engine.Any("/",webroot)
	engine.Run(":2008")
}

func webroot(context *gin.Context) {
	context.String(http.StatusOK,"hello world")
}

