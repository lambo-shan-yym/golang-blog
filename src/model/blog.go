package model

import (
	"time"
)

type Blog struct {
	Id              int64     `xorm:"pk autoincr BIGINT(20)"`
	Title           string    `xorm:"not null default '' VARCHAR(64)"`
	Content         string    `xorm:"not null default '' LONGTEXT`
	Flag            string    `xorm:"not null default '' VARCHAR(16)"`
	FirstPictureUrl string    `xorm:"not null default '' VARCHAR(128)"`
	Description     string    `xorm:"not null default '' VARCHAR(256)"`
	ViewCount       int32     `xorm:"not null default 0 BIGINT(20)"`
	LikeCount       int32     `xorm:"not null default 0 BIGINT(20)"`
	PublishType     int32     `xorm:"not null default 1 TINYINT(1)"`
	CommentAble     bool      `xorm:"not null default 1 TINYINT(1)"`
	UserId          int64     `xorm:"not null default 0 BIGINT(20)"`
	TypeId          int64     `xorm:"not null default 0 index(idx_type_id) BIGINT(20)"`
	CreateTime      time.Time `xorm:"not null default '1970-01-01 00:00:00' DATETIME"`
	UpdateTime      time.Time `xorm:"not null default '1970-01-01 00:00:00' DATETIME"`
}
