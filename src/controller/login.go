package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const userName = "user"
const password = "password"

type LoginRequest struct {
	UserName string `json:"user"`
	Password string `json:"pass"`
}

func Login(c *gin.Context) {
	var loginRequest LoginRequest
	if c.ShouldBind(&loginRequest) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "request bind error.",
		})
	}

	if loginRequest.UserName != userName || loginRequest.Password != password {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "login error.",
		})
	} else {
		session := sessions.Default(c)
		session.Set("userName", loginRequest.UserName)
		session.Save()

		c.JSON(http.StatusOK, gin.H{
			"message": "login success.",
		})
	}
}
