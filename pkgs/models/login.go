package models

import "time"

type TokenInfo struct {
	ID        int64     `gorm:"type:bigint not null auto_increment comment 'ID';primaryKey"`
	Audience  int64     `gorm:"type:bigint not null comment '用户 ID'; index" `
	ExpiresAt int64     `gorm:"type:bigint not null comment '过期时间'" `
	IssuedAt  int64     `gorm:"type:bigint not null comment '签发时间'" `
	Issuer    string    `gorm:"type:varchar(64) not null comment '签发者';"`
	NotBefore int64     `gorm:"type:bigint not null comment '可用时间'" `
	Subject   string    `gorm:"type:varchar(64) not null comment '设备';"`
	CreatedAt time.Time `gorm:"type:datetime not null default current_timestamp() comment '创建时间'"`
	UpdatedAt time.Time `gorm:"datetime not null default current_timestamp() on update current_timestamp() comment '跟新时间'"`
}
