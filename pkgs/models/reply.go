package models

import (
	"gorm.io/gorm"
	"time"
)

type Reply struct {
	ID        int64          `json:"ID" gorm:"type:bigint not null auto_increment comment 'ID';primaryKey"`
	Cid       int64          `json:"cid" gorm:"type:bigint not null; index:idx_cid_root"`
	Root      int64          `json:"root" gorm:"type:bigint not null default '-1'; index:idx_cid_root"`
	Parent    int64          `json:"parent" gorm:"type:bigint"`
	Uid       int64          `json:"uid" gorm:"type:bigint not null"`
	Real      int            `json:"real" gorm:"type:tinyint not null comment '是否实名'"`
	Aid       int64          `json:"a_uid" gorm:"type:bigint"`
	Nick      string         `json:"nick" gorm:"type:varchar(64) not null"`
	Avatar    string         `json:"avatar" gorm:"type:varchar(128) not null"`
	AtName    int            `json:"atName" gorm:"type:tinyint not null"`
	Message   string         `json:"message" gorm:"type:varchar(256) not null"`
	CreatedAt time.Time      `json:"createdAt" gorm:"type:datetime not null default current_timestamp() comment '创建时间'"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"datetime not null default current_timestamp() on update current_timestamp() comment '跟新时间'"`
	DeleteUid int64          `json:"deleteUid" gorm:"bigint"`
	DeleteAt  gorm.DeletedAt `json:"deleteAt" gorm:"index"`
}

type AtName struct {
	ID       int64          `json:"id" gorm:"type:bigint not null auto_increment comment 'ID';primaryKey"`
	ReplyID  int64          `json:"replyID" gorm:"type:bigint not null; index"`
	Uid      int64          `json:"uid" gorm:"type:bigint not null"`
	Text     string         `json:"text" gorm:"type:varchar(64) not null"`
	DeleteAt gorm.DeletedAt `json:"deleteAt" gorm:"index"`
}
