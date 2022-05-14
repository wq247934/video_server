package routers

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"video_server/stream_server/controllers"
	"video_server/stream_server/middleware"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	// programatically set swagger info
	//docs.SwaggerInfo.Title = lib.GetStringConf("base.swagger.title")
	//docs.SwaggerInfo.Description = lib.GetStringConf("base.swagger.desc")
	//docs.SwaggerInfo.Version = "1.0"
	//docs.SwaggerInfo.Host = lib.GetStringConf("base.swagger.host")
	//docs.SwaggerInfo.BasePath = lib.GetStringConf("base.swagger.base_path")
	//docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	UserLoginRouter := router.Group("/user")
	store, err := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	if err != nil {
		log.Fatalf("sessions.NewRedisStore err: %+v", err)
	}
	UserLoginRouter.Use(
		sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.TranslationMiddleware())
	{
		controllers.UserLoginControllerRegister(UserLoginRouter)
	}

	return router
}
