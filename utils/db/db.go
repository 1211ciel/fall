package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
// DBC 不同库的连接池 扩展时考虑
//DBC = make(map[string]*gorm.DB)
)

func GetMysql(addr string, maxIdleConnes, maxOpenConns, maxLiftTime int, debug bool) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       addr,  // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDb.SetMaxIdleConns(maxIdleConnes)
	sqlDb.SetMaxOpenConns(maxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Hour * time.Duration(maxLiftTime))
	if debug {
		return db.Debug() //开启debug模式
	} else {
		return db
	}
}
func GetDefaultMysql(addr string, debug bool) *gorm.DB {
	return GetMysql(addr, 10, 20, 1, debug)
}
