package model

import (
	"fmt"
	"github.com/1211ciel/fall/common/base"
	"github.com/1211ciel/fall/common/consts"
	"github.com/1211ciel/fall/core/stores/myshadow"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type LoginLog struct {
	base.Model
	Uid   uint64 `gorm:"column:uid;type:bigint unsigned not null;;comment:用户ID;index"`
	Uname string `gorm:"column:uname;type:varchar(255) not null;default:;comment:用户名;index"`
	Ip    string `gorm:"column:ip;type:varchar(255) null;default:;comment:IP;index"`
}

func (LoginLog) TableName() string {
	return "us_login_log"
}

var (
	cacheLoginLogId = "cache#LoginLog#id#%v"
)

type (
	LoginLogModel interface {
		FindLoginLogById(id uint64) (*LoginLog, error)
		CreateLoginLog(data *LoginLog) error
		UpdateLoginLog(data *LoginLog) error
		DelLoginLogById(id uint64) error
	}
	defaultLoginLogModel struct {
		myshadow.Shadow
	}
)

func (d defaultLoginLogModel) DelLoginLogById(id uint64) error {
	if err := d.delCache(id); err != nil {
		return err
	}
	if err := d.DB.Delete(&LoginLog{}, "id = ?", id).Error; err != nil {
		logx.Error(err.Error())
		return err
	}
	return nil
}

func (d defaultLoginLogModel) UpdateLoginLog(data *LoginLog) error {
	if err := d.delCache(data.ID); err != nil {
		return err
	}
	if err := d.DB.Model(data).Where("id = ?", data.ID).Updates(data).Error; err != nil {
		logx.Error(err.Error())
		return consts.ErrService
	}
	return nil
}

func (d defaultLoginLogModel) delCache(id uint64) error {
	key := fmt.Sprintf(cacheLoginLogId, id)
	if _, err := d.R.Del(key); err != nil {
		logx.Error(err.Error())
		return consts.ErrService
	}
	return nil
}

func (d defaultLoginLogModel) CreateLoginLog(data *LoginLog) error {
	if err := d.DB.Create(&data).Error; err != nil {
		return consts.ErrService
	}
	return nil
}

func (d defaultLoginLogModel) FindLoginLogById(id uint64) (*LoginLog, error) {
	var resp LoginLog
	key := fmt.Sprintf(cacheLoginLogId, id)
	err := d.Shadow.QueryRow(&resp, key, func(db *gorm.DB, v interface{}) error {
		var tempData LoginLog
		if err := d.DB.Model(&tempData).
			Where("id = ?", id).
			First(&tempData).Error; err != nil {
			logx.Error(err.Error())
			if err == gorm.ErrRecordNotFound {
				return consts.ErrDataNotFound
			}
			return consts.ErrService
		}
		*v.(*LoginLog) = tempData
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func NewLoginLogModel(db *gorm.DB, r *redis.Redis) LoginLogModel {
	return defaultLoginLogModel{myshadow.NewNodeConn(db, r)}
}
