package dao

import (
	"github.com/1211ciel/fall/calendar/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func NewDB(c *conf.Db) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       c.Addr, // DSN data source name
		DefaultStringSize:         256,    // string 类型字段的默认长度
		DisableDatetimePrecision:  true,   // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,   // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,   // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,  // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDb.SetMaxIdleConns(c.MaxIdleConnes)
	sqlDb.SetMaxOpenConns(c.MaxIdleConnes)
	sqlDb.SetConnMaxLifetime(time.Hour * time.Duration(c.MaxLifTime))
	if c.Debug {
		return db.Debug() //开启debug模式
	} else {
		return db
	}
	return nil
}

func NewDefaultDB(addr string, debug bool) *gorm.DB {
	return NewDB(&conf.Db{
		Addr:          addr,
		MaxIdleConnes: 10,
		MaxOpenConns:  20,
		MaxLifTime:    1,
		Debug:         debug,
	})
}
