package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hole/pkgs/common/utils"
	"net/http"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth, err := utils.TokenValid(ctx)
		if err != nil {
			authFailure(ctx)
			ctx.Abort()
			return
		}

		ctx.Set("auth", auth)
		ctx.Next()
	}
}

func authFailure(ctx *gin.Context) {
	zap.L().Info("jwt authorized failure")
	ctx.JSON(200, gin.H{"code": http.StatusOK, "msg": "jwt authorized failure"})
}
