package model

import (
	"github.com/1211ciel/fall/utils/dbutil"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"gorm.io/gorm"
	"testing"
)

func TestInitData(t *testing.T) {
	db, _ := getDBR()
	db.AutoMigrate(Focus{}, FocusTimeLog{})
}
func getDBR() (*gorm.DB, *redis.Redis) {
	r := redis.NewRedis("localhost:6379", redis.NodeType, "")
	d := dbutil.GetDefaultMysql("root:123456@tcp(localhost:3306)/ciel?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai", true)
	return d, r
}
