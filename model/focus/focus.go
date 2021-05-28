package model

import (
	"fmt"
	"github.com/1211ciel/fall/common/base"
	"github.com/1211ciel/fall/common/consts"
	"github.com/1211ciel/fall/core/stores/myshadow"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"gorm.io/gorm"
	"time"
)

// Focus 专注
type Focus struct {
	base.Model
	Name             string    `gorm:"column:name;type:varchar(255) not null;comment:专注名称;index"`
	Uid              uint64    `gorm:"column:uid;type:bigint unsigned not null;comment:用户ID;index"`
	TypeCode         uint8     `gorm:"column:type_code;type:tinyint unsigned not null;comment:类型Code 0倒计时 1正计时;index"`
	TotalTime        uint32    `gorm:"column:total_time;type:bigint unsigned;default 0;comment:总完成时间 单位分;"`
	FinishNum        uint32    `gorm:"column:finish_num;type:int unsigned;default 0;comment:总完成次数;"`
	Logo             string    `gorm:"column:logo;type:varchar(255);comment:图标;"`
	Desc             string    `gorm:"column:desc;type:varchar(255);comment:描述;"`
	FlagTime         uint32    `gorm:"column:flag_time;type:int unsigned;default 0;comment:目标时间,单位分钟;"`
	FinishedPoints   float32   `gorm:"column:finished_points;type:decimal(10,2);comment:当天完成百分比;"`
	CurrentTime      uint64    `gorm:"column:current_time;type:bigint unsigned not null;default:;comment:当前完成时间;"`
	CurrentBeginTime time.Time `gorm:"column:current_begin_time;type:datetime;comment:当次开始时间;"`
	LastFinishedTime time.Time `gorm:"column:last_finished_time;type:datetime;comment:最后完成时间;"`
}

var (
	cacheFocusId = "cache#Focus#id#%v"
)

func (f *Focus) TypeCodeDesc() string {
	switch f.TypeCode {
	case 0:
		return "倒计时(0)"
	case 1:
		return "正计时(1)"
	default:
		return "ERROR"
	}
}

type (
	FocusModel interface {
		FindFocusById(id uint64) (*Focus, error)
		CreateFocus(data *Focus) error
		UpdateFocus(data *Focus) error
		DelFocusById(id uint64) error
	}
	defaultFocusModel struct {
		myshadow.Shadow
	}
)

func (s *defaultFocusModel) DelFocusById(id uint64) error {
	data, err := s.FindFocusById(id)
	if err != nil {
		return err
	}
	if err := s.DB.Delete(&Focus{}, "id = ?", id).Error; err != nil {
		logx.Error(err.Error())
		return err
	}
	s.Shadow.DelCache(fmt.Sprintf(cacheFocusId, data.ID))
	return nil
}

func (s *defaultFocusModel) UpdateFocus(data *Focus) error {
	if err := s.DB.Model(data).Where("id = ?", data.ID).Updates(data).Error; err != nil {
		logx.Error(err.Error())
		return consts.ErrService
	}
	s.Shadow.DelCache(fmt.Sprintf(cacheFocusId, data.ID))
	return nil
}

func (s *defaultFocusModel) CreateFocus(data *Focus) error {
	if err := s.DB.Create(&data).Error; err != nil {
		return consts.ErrService
	}
	return nil
}

func (s *defaultFocusModel) FindFocusById(id uint64) (*Focus, error) {
	var data Focus
	key := fmt.Sprintf(cacheFocusId, id)
	if err := s.Cache.Take(&data, key, func(v interface{}) error {
		var tempData Focus
		if err := s.DB.Model(&tempData).
			Where("id = ?", id).
			First(&tempData).Error; err != nil {
			logx.Error(err.Error())
			if err == gorm.ErrRecordNotFound {
				return consts.ErrDataNotFound
			}
			return consts.ErrService
		}
		*v.(*Focus) = tempData
		return nil
	}); err != nil {
		logx.Error(err.Error())
		return nil, err
	}
	return &data, nil
}
func (s *defaultFocusModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf(cacheFocusId, primary)
}
func (s *defaultFocusModel) primaryQuery(db *gorm.DB, v, primaryKey interface{}) error {
	var data Focus
	err := db.Model(&data).Where("id = ?", primaryKey).First(&data).Error
	if err != nil {
		return err
	}
	*v.(*Focus) = data
	return nil
}

func NewFocusModel(mb *gorm.DB, r *redis.Redis) FocusModel {
	return &defaultFocusModel{
		myshadow.NewNodeConn(mb, r),
	}
}
