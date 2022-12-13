package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"hole/pkgs/routers/routes"
	"hole/pkgs/settings"
	"net/http"
	"time"
)

func setupGlobalMiddlewares(r *gin.Engine) {
	corsMiddleware := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})

	r.Use(corsMiddleware)

	// r.Use(settings.GinLogger(), settings.GinRecovery(true))
}

func Setup(r *gin.Engine, mode string) *gin.Engine {

	setupGlobalMiddlewares(r)

	mw := routes.SetupAuthRoute(r)

	r.Use(mw.MiddlewareFunc())

	routes.SetupUserRoutes(r)

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, settings.Conf.Version)
	})

	r.Run("127.0.0.1:8888")

	//v1 := r.Group("/v1")
	//v1.Use(middlewares.JwtAuthMiddleware())
	//
	//v1.POST("/signup",
	//	middlewares.JwtAuthMiddleware(),
	//	func(ctx *gin.Context) {
	//		value, _ := ctx.Get("auth")
	//		authorized := value.(*utils.Authorized)
	//		ctx.JSON(200, gin.H{"message": authorized})
	//	})

	return r
}
