package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	// session.Delete("userName")
	session.Clear()
	session.Options(sessions.Options{MaxAge: -1})
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"message": "logout success.",
	})
}
