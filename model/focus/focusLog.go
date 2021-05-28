package model

import (
	"github.com/1211ciel/fall/common/base"
	"time"
)

type FocusTimeLog struct {
	base.Model
	Uid        uint64    `gorm:"column:uid;type:bigint unsigned;comment:用户ID;index"`
	Fid        uint64    `gorm:"column:fid;type:bigint unsigned;comment:专注ID;index"`
	BeginTime  time.Time `gorm:"column:begin_time;type:datetime;comment:开始时间;"`
	EndTime    time.Time `gorm:"column:end_time;type:datetime;comment:结束时间;"`
	FinishTime uint      `gorm:"column:finish_time;type:int unsigned;comment:完成时间点,值为整点数;"`
	TotalTime  uint32    `gorm:"column:total_time;type:int unsigned;comment:完成分钟数,单位分;"`
}

func (*FocusTimeLog) TableName() string {
	return "focus_time_log"
}
