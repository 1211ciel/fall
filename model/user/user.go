package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Pid   uint64 `gorm:"column:pid;type:bigint;default:;comment:上级id;index"`
	Uname string `gorm:"column:uname;type:varchar(64) not null unique;default:;comment:用户名;index"`
	Icon  string `gorm:"column:icon;type:varchar(64) not null;default:;comment:头像;"`
	Pwd   string `gorm:"column:pwd;type:varchar(64) not null;default:;comment:密码;"`
	Phone string `gorm:"column:phone;type:varchar(64) not null;default:;comment:手机号;index"`
}

func (*User) TableName() string {
	return "us_user"
}
