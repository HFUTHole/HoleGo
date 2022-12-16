package models

import (
	"gorm.io/gorm"
	"time"
)

type Vote struct {
	ID        int64          `json:"id" gorm:"type:bigint not null auto_increment comment 'ID';primaryKey"`
	CID       int64          `json:"cid" gorm:"type:bigint not null; index" `
	Text      string         `json:"password" gorm:"type:varchar(128) not null"`
	Total     int64          `json:"total" gorm:"type:bigint not null" `
	DeletedAt gorm.DeletedAt `json:"deleteAt" gorm:"index"`
}

type VoteList struct {
	ID       int64     `json:"id" gorm:"type:bigint not null auto_increment comment 'ID';primaryKey"`
	VID      int64     `json:"vid" gorm:"type:bigint; index"`
	UID      int64     `json:"uid" gorm:"type:bigint; index"`
	CreateAt time.Time `json:"deleteAt" gorm:"index"`
}
