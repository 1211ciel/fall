package svc

import (
	model "github.com/1211ciel/fall/model/user"
	"github.com/1211ciel/fall/test/go-zero/77/service/user/internal/config"
	"github.com/1211ciel/word-of-wind/calendar/dao"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config    config.Config
	UserDB    *gorm.DB
	UserRedis *redis.Pool
	Cache     dao.Cache
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	pool := dao.NewDefaultRedis(c.MyRedis.Dns)
	userDb := dao.NewDefaultDB(c.Mysql.Dns, c.Mysql.Debug)
	return &ServiceContext{
		Config:    c,
		UserDB:    userDb,
		UserRedis: pool,
		Cache:     dao.NewCache(pool),
		UserModel: model.NewUserModel(userDb, pool),
	}
}
