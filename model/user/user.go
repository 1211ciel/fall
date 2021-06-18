package model

import "gorm.io/gorm"

type UsUser struct {
	gorm.Model
	Icon  string `gorm:"column:icon;type:varchar(64) not null;default:;comment:头像;"`
	Uname string `gorm:"column:uname;type:varchar(64) not null unique;default:;comment:用户名;index"`
	Pwd   string `gorm:"column:pwd;type:varchar(64) not null;default:;comment:密码;"`
	Phone string `gorm:"column:phone;type:varchar(64) not null;default:;comment:手机号;index"`
}

