package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo":  "bar1",
			"temp": "1111",
		}
		c.JSONP(http.StatusOK, data)
	})
	return router
}
