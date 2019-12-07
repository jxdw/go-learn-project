package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine:=gin.Default();
	engine.GET("/someGet",getting)
	engine.POST("/postSome",posting)
	engine.Run(":8082")
}

func posting(context *gin.Context) {
	buf:=bytes.Buffer{}
	buf.ReadFrom(context.Request.Body)
	context.Request.Body.Close();
	s:=buf.String();
	context.String(http.StatusOK,"request is:"+s);
}

func getting(context *gin.Context) {
	msg:=context.Query("one");
	context.JSON(http.StatusOK,gin.H{"status":"OK","message":msg})
}
