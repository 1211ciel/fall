package model

import (
	"fmt"
	"github.com/1211ciel/fall/common/base"
	"github.com/1211ciel/fall/common/consts"
	"github.com/1211ciel/fall/core/stores/myshadow"
	"github.com/1211ciel/fall/model"
	"github.com/1211ciel/fall/utils/pwdutil"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"gorm.io/gorm"
	"time"
)

type User struct {
	base.Model
	Uname         string    `gorm:"column:uname;type:varchar(255) not null unique;default:'';comment:用户名;index"`
	SysCode       string    `gorm:"column:sys_code;type:varchar(255) not null;default:'';comment:系统账号;index"`
	Pid           int       `gorm:"column:pid;type:bigint unsigned not null;default:0;comment:上级ID;"`
	Pwd           string    `gorm:"column:pwd;type:varchar(255) not null;default:'';comment:密码;"`
	Nickname      string    `gorm:"column:nickname;type:varchar(255) not null;default:'';comment:昵称;"`
	Icon          string    `gorm:"column:icon;type:varchar(255) not null;default:'';comment:头像;"`
	QrCode        string    `gorm:"column:qr_code;type:varchar(255) not null;default:'';comment:二维码;"`
	Phone         string    `gorm:"column:phone;type:varchar(255) not null;default:'';comment:电话;"`
	JoinIp        string    `gorm:"column:join_ip;type:varchar(255) not null;default:'';comment:注册IP;index"`
	LastLoginIp   string    `gorm:"column:last_login_ip;type:varchar(255) not null;default:'';comment:上次注册IP;index"`
	LastLoginTime time.Time `gorm:"column:last_login_time;type:datetime not null;default:current_timestamp;comment:上次登录时间"`
	LoginStatus   uint8     `gorm:"column:login_status;type:tinyint unsigned not null;default:0;comment:登录状态;"`
	Status        uint8     `gorm:"column:status;type:tinyint unsigned not null;default:1;comment:状态;index"`
}

func (User) TableName() string {
	return "us_user"
}

var (
	cacheUserId    = "cache#User#id#%v"
	cacheUserUname = "cache#User#Uname#%v"
)

type (
	UserModel interface {
		FindUserById(id uint64) (*User, error)
		CreateUser(data *User) error
		UpdateUser(data *User) error
		DelUserById(id uint64) error
		Register(uname string, pwd string) error
		FindUserByUname(uname string) (*User, error)
	}
	defaultUserModel struct {
		myshadow.Shadow
	}
)

// FindUserByUname  缓存值为用户ID,拿到用户ID再拿用户信息
func (s *defaultUserModel) FindUserByUname(uname string) (*User, error) {
	var data User
	key := fmt.Sprintf(cacheUserUname, uname)
	err := s.Shadow.QueryRowIndex(&data, key, s.formatPrimary, func(db *gorm.DB, v interface{}) (interface{}, error) {
		var tempData User
		err := db.Model(&data).Where("uname = ?", uname).First(&tempData).Error
		if err != nil {
			return nil, err
		}
		*v.(*User) = tempData
		return data.ID, nil
	}, s.primaryQuery)
	if err != nil {
		logx.Error(err.Error())
		return nil, err
	}
	return &data, nil
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
	s.Shadow.DelCache(fmt.Sprintf(cacheUserId, data.ID), fmt.Sprintf(cacheUserUname, data.Uname))
	return nil
}

func (s *defaultUserModel) UpdateUser(data *User) error {
	if err := s.DB.Model(data).Where("id = ?", data.ID).Updates(data).Error; err != nil {
		logx.Error(err.Error())
		return consts.ErrService
	}
	s.Shadow.DelCache(fmt.Sprintf(cacheUserId, data.ID), fmt.Sprintf(cacheUserUname, data.Uname))
	return nil
}

func (s *defaultUserModel) CreateUser(data *User) error {
	if err := s.DB.Create(&data).Error; err != nil {
		return consts.ErrService
	}
	return nil
}

func (s *defaultUserModel) FindUserById(id uint64) (*User, error) {
	var data User
	key := fmt.Sprintf(cacheUserId, id)
	if err := s.Cache.Take(&data, key, func(v interface{}) error {
		var tempData User
		if err := s.DB.Model(&tempData).
			Where("id = ?", id).
			First(&tempData).Error; err != nil {
			logx.Error(err.Error())
			if err == gorm.ErrRecordNotFound {
				return consts.ErrDataNotFound
			}
			return consts.ErrService
		}
		*v.(*User) = tempData
		return nil
	}); err != nil {
		logx.Error(err.Error())
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

func NewUserModel(mb *gorm.DB, r *redis.Redis) UserModel {
	return &defaultUserModel{
		myshadow.NewNodeConn(mb, r),
	}
}

func (s *defaultUserModel) Register(uname string, pwd string) error {
	var count int64
	if err := s.DB.Model(&User{}).Unscoped().Where("uname = ?", uname).Count(&count).Error; err != nil {
		return consts.ErrService
	}
	if count != 0 {
		return model.ErrUnameExisted
	}
	return s.CreateUser(&User{
		Uname:    uname,
		Nickname: uname,
		Pwd:      pwdutil.GenPwd(pwd),
	})
}
