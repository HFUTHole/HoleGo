package controllers

import (
	"github.com/gin-gonic/gin"
	"hole/pkgs/common/response"
	"hole/pkgs/common/utils"
	"hole/pkgs/service"
)

type SignupReq struct {
	Username  string `json:"username" validate:"required,min=4,max=63"`
	StudentId int64  `json:"studentId" validate:"required"`
	Password  string `json:"password" validate:"required,min=8,max=63"`
}

func (s *SignupReq) check() error {
	err := utils.CheckUsername(s.Username)
	if err != nil {
		return err
	}

	err = utils.CheckPassword(s.Password)
	if err != nil {
		return err
	}

	return nil
}

func Signup() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var signup SignupReq
		err := ctx.BindJSON(&signup)
		if err != nil {
			response.Error403(ctx, "参数解析错误")
			return
		}

		err = signup.check()
		if err != nil {
			response.Error403(ctx, err.Error())
			return
		}

		userVO, err := service.CreateUser(signup.Username, signup.StudentId, signup.Password)
		if err != nil {
			response.HandleBusinessException(ctx, err)
			return
		}

		response.Success(ctx, userVO)
	}
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var login LoginReq
		err := ctx.BindJSON(&login)
		if err != nil {
			response.Error403(ctx, "参数解析错误")
			return
		}

		loginVO, err := service.Login(login.Username, login.Password, "")

		if err != nil {
			response.HandleBusinessException(ctx, err)
			return
		}

		response.Success(ctx, loginVO)
	}
}
