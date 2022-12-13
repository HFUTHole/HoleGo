package user

import (
	"github.com/gin-gonic/gin"
	"hole/pkgs/common/app"
	"net/http"
)

func GetProfile(c *gin.Context) {
	appG := app.Gin{C: c}

	appG.Response(http.StatusNotFound, http.StatusNotFound, "用户不存在", nil)
}
