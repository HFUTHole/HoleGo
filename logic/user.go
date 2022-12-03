package logic

import (
	"hole_go/dao/mysql"
	"hole_go/models"
)

func SignUp(u *models.User) (err error) {
	// 1.判断用户是否存在
	//err := mysql.checkUserExist(u.Username)
	//if err != nil {
	//	// 数据库查询出错
	//	return err
	//}

	// 构造一个User实例保存进数据库
	user := models.User{
		Id:            u.Id,
		StudentNumber: u.StudentNumber,
		Username:      u.Username,
		Password:      u.Username,
	}
	return mysql.SignUp(user)
}
