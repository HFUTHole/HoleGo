package models

import (
	"gorm.io/gorm"
	"time"
)

type Reply struct {
	ID        int64          `json:"ID" gorm:"type:bigint comment 'ID'"`
	Root      int64          `json:"root" gorm:"type:bigint"`
	Parent    int64          `json:"parent" gorm:"type:bigint"`
	Cid       int64          `json:"cid" gorm:"type:bigint not null; index"`
	Uid       int            `json:"uid" gorm:"type:bigint not null"`
	Real      int            `json:"real" gorm:"type:tinyint not null comment '是否实名'"`
	Nick      string         `json:"nick"`
	Avatar    string         `json:"avatar"`
	Message   string         `json:"message" gorm:"type:varchar(256) not null"`
	CreatedAt time.Time      `json:"createdAt" gorm:"type:datetime not null default current_timestamp() comment '创建时间'"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"datetime not null default current_timestamp() on update current_timestamp() comment '跟新时间'"`
	DeleteAt  gorm.DeletedAt `json:"deleteAt" gorm:"index"`
}

type AtName struct {
	ID       int64          `json:"id" gorm:"type:bigint not null"`
	ReplyID  int64          `json:"replyID" gorm:"type:bigint not null"`
	Text     string         `json:"text" gorm:"type:varchar(64) not null"`
	Uid      int64          `json:"uid" gorm:"type:bigint not null"`
	DeleteAt gorm.DeletedAt `json:"deleteAt" gorm:"index"`
}
