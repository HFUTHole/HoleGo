package vo

import (
	"time"
)

type ReplyVO struct {
	ID        int64               `json:"id"`
	Cid       int64               `json:"cid"`
	Uid       int64               `json:"uid"`
	Nick      string              `json:"nick"`
	Avatar    string              `json:"avatar"`
	Message   string              `json:"message"`
	CreatedAt time.Time           `json:"createdAt"`
	AtName    map[string]AtNameVO `json:"atName"`
	List      []ReplyChildVO      `json:"list"`
}

type ReplyChildVO struct {
	ID        int64               `json:"id"`
	Cid       int64               `json:"cid"`
	Root      int64               `json:"root"`
	Parent    int64               `json:"parent"`
	Uid       int64               `json:"uid"`
	Nick      string              `json:"nick"`
	Avatar    string              `json:"avatar"`
	Message   string              `json:"message"`
	AtName    map[string]AtNameVO `json:"atName"`
	CreatedAt time.Time           `json:"createdAt"`
}

type AtNameVO struct {
	ID      int64  `json:"id"`
	ReplyID int64  `json:"rid"`
	Text    string `json:"text"`
	Uid     int64  `json:"uid"`
}
