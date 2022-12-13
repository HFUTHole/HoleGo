package models

type User struct {
	Id        int    `json:"id" gorm:"primarykey"`
	StudentId int    `json:"studentId"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}
