package svc

import (
	"github.com/1211ciel/fall/model/user"
	"github.com/1211ciel/fall/toy/fall/config"
	"github.com/1211ciel/fall/utils/dbutil"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	R         *redis.Redis
	UserModel model.UserModel
)

func InitSvc() {
	DB, R = getDBR()
}
func getDBR() (*gorm.DB, *redis.Redis) {
	r := redis.NewRedis(config.C.CacheRedis.Host, redis.NodeType, "")
	d := dbutil.GetDefaultMysql(config.C.Mysql.DataSource, false)
	return d, r
}
