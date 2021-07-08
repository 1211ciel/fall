package svc

import (
	"github.com/1211ciel/fall/model/user"
	"github.com/1211ciel/fall/test/go-zero/78/service/user/internal/config"
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
	userDB := dao.NewDefaultDB(c.MySql.Addr, c.MySql.Debug)
	userRedis := dao.NewDefaultRedis(c.MyRedis.Addr)
	return &ServiceContext{
		Config:    c,
		UserDB:    userDB,
		UserRedis: userRedis,
		Cache:     dao.NewCache(userRedis),
		UserModel: model.NewUserModel(userDB, userRedis),
	}
}
