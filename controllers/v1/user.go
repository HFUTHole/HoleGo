package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hole_go/dao/mysql"
	e "hole_go/errorResponse"
	"hole_go/logic"
	"hole_go/models"
)

func SignUp(c *gin.Context) {
	var user *models.User
	// 参数错误
	if err := c.ShouldBind(&user); err != nil {
		e.ResponseErrorWithMsg(c, e.CodeInvalidParams)
		return
	}
	if err := logic.SignUp(user); err != nil {
		zap.L().Error("logic.signup failed", zap.Error(err))
		// 用户已存在
		if errors.Is(err, mysql.ErrorUserExit) {
			e.ResponseError(c, e.CodeUserExist)
			return
		}
		e.ResponseError(c, e.CodeServerBusy)
		return
	}
	e.ResponseSuccess(c, nil)
}






















