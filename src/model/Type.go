package model

import "time"

type Type struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)" json:"id"`
	Name       string    `xorm:"not null default '' unique(uni_name) VARCHAR(64)" json:"name"`
	CreateTime time.Time `xorm:"not null default '1970-01-01 00:00:00' DATETIME" json:"create_time"`
	UpdateTime time.Time `xorm:"not null default '1970-01-01 00:00:00' DATETIME" json:"update_time"`
}
