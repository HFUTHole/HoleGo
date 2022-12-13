package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"hole/pkgs/middlewares"
	"hole/pkgs/service/auth"
)

func SetupAuthRoute(e *gin.Engine) *jwt.GinJWTMiddleware {
	jwtMiddleware := middlewares.JwtMiddleWare()

	r := e.Group("/auth")

	r.POST("/login",
		auth.LoginResponseHandler(jwtMiddleware),
		jwtMiddleware.LoginHandler,
	)

	return jwtMiddleware
}
