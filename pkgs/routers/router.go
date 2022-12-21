package routers

import (
	"github.com/gin-gonic/gin"
	"hole/pkgs/config/base"
	"hole/pkgs/config/logger"
	"hole/pkgs/routers/controllers"
)

func Setup() *gin.Engine {
	if base.GetMode() == gin.ReleaseMode {
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

	api.POST("/image",
		controllers.UpdateImage(),
	)

	r.GET(
		"/image/:bucket/:id",
		controllers.DownloadImage(),
	)

	api.POST("/content",
		controllers.CreateContent(),
	)

	api.DELETE("/content/:cid",
		controllers.DeleteContent(),
	)

	api.GET("/content/:cid",
		controllers.GetContent())

	api.GET("/content/page",
		controllers.GetContentPage(),
	)
	api.POST("/content/voting",
		controllers.CreateContentVote(),
	)

	api.POST("/content/voting/:cid",
		controllers.Vote(),
	)

	api.DELETE("/content/voting/:cid",
		controllers.DeleteContentVote(),
	)
	return r
}
