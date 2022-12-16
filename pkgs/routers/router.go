package routers

import (
	"github.com/gin-gonic/gin"
	"hole/pkgs/config"
	"hole/pkgs/config/logger"
	"hole/pkgs/routers/controllers"
)

func Setup() *gin.Engine {
	if config.GetMode() == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 设置发布或者生产模式
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/version", controllers.Version())

	api := r.Group("/api")
	//api.Use(middlewares.JwtAuthMiddleware())

	api.POST("/signup",
		controllers.Signup(),
	)

	api.POST("/content/create",
		controllers.CreateContent(),
	)

	api.GET("/content/one",
		controllers.GetContent())

	api.GET("/content/page",
		controllers.GetContentPage(),
	)

	return r
}
