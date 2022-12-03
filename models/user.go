package models

type User struct {
	Id            int `json:"id" gorm:"primarykey"`
	StudentNumber int	`json:"studentNumber"`
	Username      string `json:"username"`
	Password string `json:"password"`
}

