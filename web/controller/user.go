package controller

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct{}
func (u *User) Login(c *gin.Context) {

	log.Print("Received Login API request")
	var req map[string]interface{}
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"result": false, "error": err.Error()})
		return
	}
	if req["username"] == nil {
		err := errors.New("Field username is rquired.")
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"result": false, "error": err.Error()})
		return
	}
	if req["password"] == nil {
		err := errors.New("Field password is rquired.")
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"result": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": true, "data": req})
}