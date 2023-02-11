package main

import (
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memcached"
	// "github.com/gin-contrib/sessions/cookie"
	"github.com/bradfitz/gomemcache/memcache"

	// "try_go_gin_server/controller"
)

func setupRouter() *gin.Engine {
	var router = gin.Default()

	store := memcached.NewStore(memcache.New("memcached:11211"), "", []byte("secret"))
	router.Use(sessions.Sessions("login-session", store))

	// setting session
	// store := cookie.NewStore([]byte("secret"))
	// router.Use(sessions.Sessions("login-session", store))

	// router.POST("/login", controller.Login)
	// router.POST("/logout", controller.Logout)

	router.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}

		session.Set("count", count)
		session.Save()

		c.JSON(200, gin.H{"count": count})
	})

	return router
}

func main() {
	// gin.SetMode(gin.ReleaseMode)

	var router = setupRouter()

	router.Run(":8080")
}
