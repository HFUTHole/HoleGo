package utils

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Validate(v interface{}) error {
	return validate.Struct(v)
}

func CheckUsername(username string) error {
	l := len(username)
	if l < 1 {
		return errors.New("用户名不能少于1位")
	}
	if l >= 32 {
		return errors.New("用户名不能多于32位")
	}

	return nil
}

func CheckPassword(password string) error {
	l := len(password)
	if l < 8 {
		return errors.New("密码不能少于8位")
	}
	if l >= 32 {
		return errors.New("密码不能多于32位")
	}

	var lower = true
	var upper = true
	var number = true
	var illegal = false

	// 保证每次校验的时间相同
	for _, ch := range password {
		if ch >= 'a' && ch <= 'z' {
			lower = false
			continue
		}
		if ch >= 'A' && ch <= 'Z' {
			upper = false
			continue
		}
		if ch >= '0' && ch <= '9' {
			number = false
			continue
		}

		if ch == '.' || ch == '_' || ch == '@' {
			continue
		}
		illegal = true
	}

	if illegal {
		return errors.New("密码只能包含[a-zA-Z._@]")
	}
	if lower {
		return errors.New("密码必须包含小写字母")
	}
	if upper {
		return errors.New("密码必须包含大写字母")
	}
	if number {
		return errors.New("密码必须包含数字")
	}
	return nil
}
