package dao

import (
	"hole/pkgs/dao"
	"hole/pkgs/models"
	"testing"
)

func TestCreateLogin(t *testing.T) {
	db := GetDB()

	login := &models.TokenInfo{
		ID:        1,
		Audience:  1,
		ExpiresAt: 0,
		IssuedAt:  0,
		Issuer:    "zou yu",
		NotBefore: 0,
		Subject:   "Window 10",
	}

	err := dao.CreateToken(db, login)
	if err != nil {
		t.Error(err)
		return
	}
}