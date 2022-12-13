package service

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"hole/src/config/logger"
	"hole/src/config/mysql"
	"hole/src/dao"
	"hole/src/exception"
	"hole/src/models"
	"hole/src/models/role"
	"hole/src/models/vo"
	"hole/src/pkg/utils"
	"strconv"
	"time"
)

func CreateUser(username string, studentId int64, password string) (*vo.UserVO, error) {
	db := mysql.GetDB()

	generatePassword, err := utils.GeneratePassword(password)
	if err != nil {
		return &vo.UserVO{}, &exception.BusinessException{Msg: "密码不合法"}
	}

	user := &models.User{
		StudentId: studentId,
		Username:  username,
		Password:  generatePassword,
		Role:      role.CreateRole(role.NormalUserRole, role.UnauthorizedRole),
		Avatar:    "default_avatar.png",
		Email:     "",
		About:     "",
	}

	err = AuthorizedUser(user)
	if err != nil {
		return nil, &exception.BusinessException{Msg: "信息校验失败"}
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		err = dao.CreateUser(tx, user)
		if err != nil {
			logger.GetLogger().Error("用户名已存在", zap.Error(err))
			return &exception.ClientException{Msg: "用户名已存在"}
		}
		return nil
	})

	user.ClearPassword()
	if err != nil {
		return vo.ConvertUserVO(user), &exception.ClientException{
			Msg: err.Error(),
		}
	}
	return vo.ConvertUserVO(user), nil
}

func AuthorizedUser(user *models.User) error {
	user.Role = user.Role.Revoke(role.UnauthorizedRole)
	user.AuthTime = time.Now()
	return nil
}

func Login(uid string, password string, device string) (*vo.LoginVO, error) {
	var user *models.User
	db := mysql.GetDB()
	if len(uid) < 1 {
		return nil, &exception.ClientException{Msg: "用户不正确"}
	}

	if uid[0] == '#' {
		id, err := strconv.Atoi(uid[1:])
		if err == nil {
			user, err = dao.GetUserByID(db, int64(id))
			if err != nil {
				return nil, &exception.ClientException{Msg: "用户名不存在"}
			}
		} else {
			return nil, &exception.ClientException{Msg: "用户名不合法"}
		}
	} else {
		id, err := strconv.Atoi(uid)
		if err == nil {
			user, err = dao.GetUserByStudentID(db, int64(id))
			if err != nil {
				return nil, &exception.ClientException{Msg: "用户名不存在"}
			}
		} else {
			return nil, &exception.ClientException{Msg: "用户名不合法"}
		}
	}

	if err := utils.VerifyPassword(password, user.Password); err != nil {
		return nil, &exception.ClientException{Msg: "密码错误"}
	}

	return nil, nil
}
