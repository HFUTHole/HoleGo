package utils

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

var TokenHourLifespan int64 = 60 * 60 * 48 // 48 hour
var Secret = "MDE0ODAwLCJpYXQiOjE2NzA4"

func InitToken(secret string, tokenHourLifespan int64) {
	Secret = secret
	TokenHourLifespan = tokenHourLifespan
}

type Authorized struct {
	Audience  int64  `json:"aud,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	Id        string `json:"jti,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   string `json:"sub,omitempty"`
}

func GenerateToken(id int64, sub string) (string, error) {
	claims := jwt.MapClaims{}
	claims["aud"] = id
	claims["sub"] = sub
	now := time.Now()
	claims["iat"] = now.Unix()
	claims["exp"] = now.Add(time.Duration(TokenHourLifespan) * time.Second).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(Secret))
}

func TokenValid(ctx *gin.Context) (*Authorized, error) {
	token := ExtractToken(ctx)

	parse, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(Secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := parse.Claims.(jwt.MapClaims)
	valid := claims.Valid()
	if ok && valid != nil {
		return nil, valid
	}

	return ParseClaims(claims), nil
}

func ExtractToken(ctx *gin.Context) string {
	token := ctx.Query("token")
	if token != "" {
		return token
	}

	bearerToken := ctx.GetHeader("Authorization")
	split := strings.Split(bearerToken, " ")

	if len(split) == 2 && split[0] == "Bearer" {
		return split[1]
	}

	return ""
}

func ParseClaims(claims jwt.MapClaims) *Authorized {
	auth := Authorized{}
	// 用户
	auth.Audience = toInt64(claims["aud"])
	// 到期时间
	auth.ExpiresAt = toInt64(claims["exp"])
	auth.IssuedAt = toInt64(claims["iat"])
	// JWT ID用于标识该JWT
	jti := claims["jti"]
	if jti != nil {
		auth.Id = (claims["jti"]).(string)
	}
	// 发行人
	iss := claims["iss"]
	if iss != nil {
		auth.Issuer = (claims["iss"]).(string)
	}
	// 在此之前不可用
	nbf := claims["nbf"]
	if nbf != nil {
		auth.NotBefore = toInt64(nbf)
	}
	// 主题
	sub := claims["sub"]
	if sub != nil {
		auth.Subject = (claims["sub"]).(string)
	}
	return &auth
}

func toInt64(value interface{}) int64 {
	switch nbf := value.(type) {
	case float64:
		return int64(value.(float64))
	case json.Number:
		v, _ := nbf.Int64()
		return v
	}

	return 0
}
