package model

import "time"

type User struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	UserName   string    `xorm:"not null default '' unique(uni_username) VARCHAR(64)"`
	NickName   string    `xorm:"not null default '' VARCHAR(64)"`
	Password   string    `xorm:"not null default '' VARCHAR(64)"`
	Salt       string    `xorm:"not null default '' VARCHAR(64)"`
	Email      string    `xorm:"not null default '' unique(uni_email) VARCHAR(64)"`
	CreateTime time.Time `xorm:"not null default '1970-01-01 00:00:00' DATETIME"`
	UpdateTime time.Time `xorm:"not null default '1970-01-01 00:00:00' DATETIME"`
}
