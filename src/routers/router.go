package routers

import (
	"hole/src/middlewares"
	"hole/src/pkg/utils"
	"hole/src/settings"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 设置发布或者生产模式
	}
	r := gin.New()
	r.Use(settings.GinLogger(), settings.GinRecovery(true))

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, settings.Conf.Version)
	})

	v1 := r.Group("/v1")
	v1.Use(middlewares.JwtAuthMiddleware())

	v1.POST("/signup",
		middlewares.JwtAuthMiddleware(),
		func(ctx *gin.Context) {
			value, _ := ctx.Get("auth")
			authorized := value.(*utils.Authorized)
			ctx.JSON(200, gin.H{"message": authorized})
		})

	return r
}
