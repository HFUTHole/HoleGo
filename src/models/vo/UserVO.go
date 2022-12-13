package vo

import (
	"hole/src/models"
	"time"
)

type UserVO struct {
	ID        int64     `json:"id"`
	StudentId int64     `json:"studentId"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	AuthTime  time.Time `json:"authTime"`
	Sex       string    `json:"sex"`
	Avatar    string    `json:"avatar"`
	Email     string    `json:"email"`
	About     string    `json:"about"`
	CreatedAt time.Time `json:"createdAt"`
}

func ConvertUserVO(u *models.User) *UserVO {
	return &UserVO{
		ID:        u.ID,
		StudentId: u.StudentId,
		Username:  u.Username,
		Role:      u.Role.String(),
		AuthTime:  u.AuthTime,
		Sex:       u.SexToString(),
		Avatar:    u.Avatar,
		Email:     u.Email,
		About:     u.About,
		CreatedAt: u.CreatedAt,
	}
}
