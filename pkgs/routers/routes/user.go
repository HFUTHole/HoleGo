package routes

import (
	"github.com/gin-gonic/gin"
	"hole/pkgs/service/user"
)

func SetupUserRoutes(e *gin.Engine) {
	r := e.Group("/user")

	r.GET("/profile", user.GetProfile)
}
