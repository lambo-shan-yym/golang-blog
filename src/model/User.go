package model

import "time"

type User struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)" json:"id"`
	UserName   string    `xorm:"not null default '' unique(uni_username) VARCHAR(64)" json:"user_name"`
	NickName   string    `xorm:"not null default '' VARCHAR(64)" json:"nick_name"`
	Password   string    `xorm:"not null default '' VARCHAR(64)" json:"password"`
	Salt       string    `xorm:"not null default '' VARCHAR(64)" json:"salt"`
	Email      string    `xorm:"not null default '' unique(uni_email) VARCHAR(64)" json:"email"`
	CreateTime time.Time `xorm:"not null default '1970-01-01 00:00:00' DATETIME" json:"create_time"`
	UpdateTime time.Time `xorm:"not null default '1970-01-01 00:00:00' DATETIME" json:"update_time"`
}
