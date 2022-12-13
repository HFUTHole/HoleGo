package auth

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"hole/pkgs/common/app"
	"hole/pkgs/models"
	"net/http"
	"time"
)

type LoginForm struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func LoginAuthenticatorHandlers(c *gin.Context) (interface{}, error) {
	var form LoginForm

	if err := c.ShouldBind(&form); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	return &models.User{
		Username:  form.Username,
		StudentId: 2021217986,
	}, nil
}

func login() {}

func LoginResponseHandler(mw *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return func(c *gin.Context) {
		mw.LoginResponse = func(c *gin.Context, code int, token string, expireTime time.Time) {
			// redis 缓存token
			fmt.Println("当前的token是", token)

			appG := app.Gin{c}

			appG.Response(http.StatusOK, 200, "登录成功", token)
		}
		c.Next()
	}
}
