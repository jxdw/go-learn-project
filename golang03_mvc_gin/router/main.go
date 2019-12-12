package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "go-learn-code/golang03_mvc_gin/router/docs"
	"net/http"
)

func main() {
	engine:=gin.Default();
	engine.GET("/someGet",getting)
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.Group("/api/v1")
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

// @Summary 获取查询的信息
// @Produce  json
// @Param one path string true "查询的信息"
// @Success 200 {string} string "{"status":"OK","message":"ok"}"
// @Router /someGet?one={one} [get]
func getting(context *gin.Context) {
	msg:=context.Query("one");
	context.JSON(http.StatusOK,gin.H{"status":"OK","message":msg})
}
