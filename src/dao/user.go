package dao

// CheckUserExist 检查用户名是否存在
//func CheckUserExist(username string) (error error) {
//	sqlstr := `select count(user_id) from user where username = ?`
//	var count int
//	if err := db.Get(&count,sqlstr,username);err != nil {
//		return err
//	}
//	if count > 0 {
//		return errors.New("用户已存在")
//	}
//	return
//}

// SignUp 创建用户
//func SignUp(user models.User) (error error) {
//	db.Create(user)
//	return
//}
