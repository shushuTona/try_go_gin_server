package main

import (
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"try_go_gin_server/controller"
)

func setupRouter() *gin.Engine {
	var router = gin.Default()

	// setting session
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("login-session", store))

	router.POST("/login", controller.Login)
	router.POST("/logout", controller.Logout)

	return router
}

func main() {
	// gin.SetMode(gin.ReleaseMode)

	var router = setupRouter()

	router.Run(":8080")
}
