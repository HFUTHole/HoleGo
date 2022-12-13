package routers

import (
	"github.com/gin-gonic/gin"
	"hole/src/config"
	"hole/src/config/logger"
	"hole/src/routers/controllers"
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

	return r
}
