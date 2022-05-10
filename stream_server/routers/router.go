package routers

import (
	"github.com/gin-gonic/gin"
	"video_server/stream_server/controllers"
)

func SetRouters() *gin.Engine {
	r := gin.Default()
	r.GET("/video/:vid", controllers.StreamHandler)
	r.POST("/video/:vid", controllers.UploadHandler)
	r.Use(Cors())
	return r
}

// Cors 开启跨域函数
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Origin")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
