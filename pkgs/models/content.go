package models

import (
	"gorm.io/gorm"
	"time"
)

type Content struct {
	ID        int64          `json:"id" gorm:"type:bigint ;primaryKey"`
	Uid       int64          `json:"uid" gorm:"type:bigint ;index"`
	Nick      string         `json:"nick" gorm:"varchar(64) not null"`
	Avatar    string         `json:"Avatar" gorm:"type:varchar(64) not null"`
	Like      int64          `json:"like" gorm:"type:int default '0'"`
	Real      int            `json:"real" gorm:"type:tinyint not null default '0' comment '是否实名'"`
	Aid       int64          `json:"aid" gorm:"type:bigint"`
	Voting    int            `json:"voting" gorm:"type:tinyint not null default '0'"` // 0 未开启 1 开启
	EndTime   time.Time      `json:"endTime" gorm:"type:datetime default null"`
	Title     string         `json:"title" gorm:"type:varchar(32) not null ;index"`
	Text      string         `json:"text" gorm:"type:varchar(2048) not null"`
	CreatedAt time.Time      `json:"createdAt" gorm:"type:datetime not null default current_timestamp() comment '创建时间'; index"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"datetime not null default current_timestamp() on update current_timestamp() comment '跟新时间'"`
	DeleteUid int64          `json:"deleteUid" gorm:"type:bigint"`
	DeleteAt  gorm.DeletedAt `json:"deleteAt" gorm:"index"`
}

type ContentImage struct {
	ID  int64  `json:"id" gorm:"type:bigint comment 'ID'; primaryKey"`
	Cid int64  `json:"cid" gorm:"type:bigint not null; index"`
	URL string `json:"url" gorm:"type:varchar(128) not null"`
}

type ContentJumpUrl struct {
	ID       int64          `json:"id" gorm:"type:bigint comment 'ID'; primaryKey"`
	Cid      int64          `json:"cid" gorm:"type:bigint not null; index"`
	Text     string         `json:"text" gorm:"type:varchar(64) not null comment '关键词' "`
	JumpCid  int64          `json:"jumpCid" gorm:"type:bigint not null comment '跳转 ID'"`
	DeleteAt gorm.DeletedAt `json:"deleteAt" gorm:"index"`
}

type Tag struct {
	ID   int64  `json:"id" gorm:"type:bigint; primaryKey"`
	Text string `json:"text" gorm:"type:varchar(32) not null; uniqueIndex"`
}

type ContentTag struct {
	ID  int64 `json:"id" gorm:"type:bigint; primaryKey"`
	Cid int64 `json:"cid" gorm:"type:bigint; index"`
	Tid int64 `json:"tid" gorm:"type:bigint"`
}

type ContentLiked struct {
	ID        int64     `json:"id" gorm:"type:bigint ;primaryKey"`
	Cid       int64     `json:"cid" gorm:"type:bigint not null; index:idx_cid_uid,unique"`
	Uid       int64     `json:"uid" gorm:"type:bigint not null; index:idx_cid_uid,unique"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:datetime not null default current_timestamp() comment '创建时间'"`
}
