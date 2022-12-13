package dao

import (
	"hole/src/models"
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

	err := CreateToken(db, login)
	if err != nil {
		t.Error(err)
		return
	}

}
