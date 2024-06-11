package router

import (
	"demo08/router/user"
	"demo08/server"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api")
	{
		v1.POST("/user/login", Wrapper(server.Login))
		user.UserApi(v1)
	}
	return router
}
