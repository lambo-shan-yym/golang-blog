package model

import (
	"time"
)

type Blog struct {
	Id              int64     `xorm:"pk autoincr BIGINT(20)" json:"id"`
	Title           string    `xorm:"not null default '' VARCHAR(64)" json:"title"`
	Content         string    `xorm:"not null default '' LONGTEXT" json:"content"`
	Flag            string    `xorm:"not null default '' VARCHAR(16)" json:"flag"`
	FirstPictureUrl string    `xorm:"not null default '' VARCHAR(128)" json:"first_picture_url"`
	Description     string    `xorm:"not null default '' VARCHAR(256)" json:"description"`
	ViewCount       int32     `xorm:"not null default 0 BIGINT(20)" json:"view_count"`
	LikeCount       int32     `xorm:"not null default 0 BIGINT(20)" json:"like_count"`
	PublishType     int32     `xorm:"not null default 1 TINYINT(1)" json:"publish_type"`
	CommentAble     bool      `xorm:"not null default 1 TINYINT(1)" json:"comment_able"`
	UserId          int64     `xorm:"not null default 0 BIGINT(20)" json:"user_id"`
	TypeId          int64     `xorm:"not null default 0 index(idx_type_id) BIGINT(20)" json:"type_id"`
	CreateTime      time.Time `xorm:"not null default '1970-01-01 00:00:00' DATETIME" json:"create_time"`
	UpdateTime      time.Time `xorm:"not null default '1970-01-01 00:00:00' DATETIME" json:"update_time"`
}
