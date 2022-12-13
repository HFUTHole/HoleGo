package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"hole/pkgs/common/utils"
	"testing"
)

func TestGenerateToken(t *testing.T) {

	token, err := utils.GenerateToken(123456, "sub")
	if err != nil {
		t.Error("generate token error", err)
	}

	t.Log("token: ", token)

}

var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOjEyMzQ1NiwiZXhwIjoxNjcwNjU1NTg3LCJpYXQiOjE2NzA2NTU1MjcsInN1YiI6IjEyMzQ1Njc4OSJ9.qLG1u6Eg9vyCl1ey1yBDKWX1lWEzJ19ttJ1W9WOzuhg"

func TestTokenValid(t *testing.T) {
	//token, _ := GenerateToken(123456)
	parse, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(utils.Secret), nil
	})

	t.Logf("err: %v\n", err)
	claims := parse.Claims.(jwt.MapClaims)

	authorized := utils.ParseClaims(claims)
	t.Logf("aud: %v", authorized)
	t.Logf("token %v\n", parse)
}

func TestTokenValid1(t *testing.T) {
	token, _ = utils.GenerateToken(123456, "sub")
	err := TokenValidT(token)
	if err != nil {
		t.Error(err)
	}
}

func TokenValidT(token string) error {
	parse, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(utils.Secret), nil
	})

	if err != nil {
		return err
	}

	claims, ok := parse.Claims.(jwt.MapClaims)
	valid := claims.Valid()
	if ok && valid != nil {
		return valid
	}

	fmt.Println(utils.ParseClaims(claims))
	return nil
}
