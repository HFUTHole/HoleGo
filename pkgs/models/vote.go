package models

import (
	"gorm.io/gorm"
	"time"
)

type VotingOption struct {
	ID        int64          `json:"id" gorm:"type:bigint not null auto_increment comment 'ID';primaryKey"`
	Cid       int64          `json:"cid" gorm:"type:bigint not null; index" `
	Text      string         `json:"text" gorm:"type:varchar(128) not null"`
	Total     int64          `json:"total" gorm:"type:bigint not null"`
	DeletedAt gorm.DeletedAt `json:"deleteAt" gorm:"index"`
}

type VotingInfo struct {
	ID        int64     `json:"id" gorm:"type:bigint not null auto_increment comment 'ID';primaryKey"`
	Cid       int64     `json:"cid" gorm:"type:bigint not null; index:idx_cid_uid,unique"`
	Uid       int64     `json:"uid" gorm:"type:bigint not null; index:idx_cid_uid,unique"`
	Vid       int64     `json:"vid" gorm:"type:bigint not null; index"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:datetime not null default current_timestamp() comment '创建时间'"`
}
