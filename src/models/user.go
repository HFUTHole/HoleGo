package models

import (
	"hole/src/models/role"
	"time"
)

type User struct {
	ID        int64         `json:"id" gorm:"type:bigint not null auto_increment comment 'ID';primaryKey"`
	StudentId int64         `json:"studentId" gorm:"type:bigint not null comment '学号'; index" `
	Username  string        `json:"username" gorm:"type:varchar(64) not null comment '昵称'"`
	Password  string        `json:"password" gorm:"type:varchar(128) not null comment '密码'"`
	Role      role.UserRole `json:"role" gorm:"type:int not null default '0' comment '角色'"`
	AuthTime  time.Time     `json:"authTime" gorm:"type:datetime comment '认证时间'"`
	Sex       int           `json:"sex" gorm:"type:tinyint not null default '-1' comment '性别, -1: 未知, 0: 男, 1: 女'"`
	Avatar    string        `json:"avatar" gorm:"type:varchar(256) not null default '' comment '头像'"`
	Email     string        `json:"email" gorm:"type:varchar(128) not null default '' comment '邮箱'"`
	About     string        `json:"about" gorm:"type:varchar(256) not null default '' comment '个人简介'"`
	CreatedAt time.Time     `json:"createdAt" gorm:"type:datetime not null default current_timestamp() comment '创建时间'"`
	UpdatedAt time.Time     `json:"updatedAt" gorm:"datetime not null default current_timestamp() on update current_timestamp() comment '跟新时间'"`
}

func (u *User) ClearPassword() {
	u.Password = ""
}

func (u *User) SexToString() string {
	if u.Sex == 0 {
		return "男"
	}
	if u.Sex == 1 {
		return "女"
	}
	return "未知"
}

func (u *User) IsSelf(id int64) bool {
	return u.ID == id
}
