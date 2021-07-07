package model

import (
	"errors"
	"fmt"
	"github.com/1211ciel/word-of-wind/calendar/consts"
	"github.com/1211ciel/word-of-wind/calendar/dao"
	"github.com/gomodule/redigo/redis"
	"github.com/tal-tech/go-zero/core/logx"
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

func NewUser(pid uint64, uname, icon, pwd, phone string) User {
	return User{Pid: pid, Uname: uname, Icon: icon, Pwd: pwd, Phone: phone}
}

var (
	cacheUserId = "cache#User#id#%v"
)

type (
	UserModel interface {
		FindUserById(id uint64) (*User, error)
		CreateUser(data *User) error
		UpdateUser(data *User) error
		DelUserById(id uint64) error
		FindByUsername(uname string) (*User, error)
		CheckUserNameExist(uname string) error
	}
	defaultUserModel struct {
		dao.Shadow
		DB *gorm.DB
	}
)

func (s *defaultUserModel) CheckUserNameExist(uname string) error {
	var count int64
	err := s.DB.Model(&User{}).Where("uname = ?", uname).Count(&count).Error
	if err != nil {
		return consts.ErrData
	}
	if count != 0 {
		return consts.ErrDataAlreadyExist
	}
	return nil
}

func (s *defaultUserModel) FindByUsername(uname string) (*User, error) {
	var u User
	err := s.DB.Model(&u).Where("uname = ?", uname).First(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("数据不存在")
		}
	}
	return &u, nil
}

func NewUserModel(db *gorm.DB, r *redis.Pool) UserModel {
	return &defaultUserModel{
		dao.NewCache(r),
		db,
	}
}

func (s *defaultUserModel) DelUserById(id uint64) error {
	data, err := s.FindUserById(id)
	if err != nil {
		return err
	}
	if err := s.DB.Delete(&User{}, "id = ?", id).Error; err != nil {
		logx.Error(err.Error())
		return err
	}
	return s.Shadow.Del(fmt.Sprintf(cacheUserId, data.ID))
}

func (s *defaultUserModel) UpdateUser(data *User) error {
	if err := s.DB.Model(data).Where("id = ?", data.ID).Updates(data).Error; err != nil {
		logx.Error(err.Error())
		return consts.ErrService
	}
	return s.Shadow.Del(fmt.Sprintf(cacheUserId, data.ID))
}

func (s *defaultUserModel) CreateUser(data *User) error {
	if err := s.DB.Create(&data).Error; err != nil {
		logx.Error(err.Error())
		return consts.ErrService
	}
	return nil
}

func (s *defaultUserModel) FindUserById(id uint64) (*User, error) {
	var data User
	key := fmt.Sprintf(cacheUserId, id)
	if err := s.Take(&data, key, func(v interface{}) error {
		var temp User
		if err := s.DB.Model(&temp).
			Where("id = ?", id).
			First(&temp).Error; err != nil {
			logx.Error(err.Error())
			if err == gorm.ErrRecordNotFound {
				return consts.ErrDataNotFound
			}
			return consts.ErrService
		}
		*v.(*User) = temp
		return nil
	}); err != nil {
		return nil, err
	}
	return &data, nil
}
func (s *defaultUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf(cacheUserId, primary)
}
func (s *defaultUserModel) primaryQuery(db *gorm.DB, v, primaryKey interface{}) error {
	var data User
	err := db.Model(&data).Where("id = ?", primaryKey).First(&data).Error
	if err != nil {
		return err
	}
	*v.(*User) = data
	return nil
}
