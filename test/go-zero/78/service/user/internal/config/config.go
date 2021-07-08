package config

import (
	"github.com/1211ciel/word-of-wind/calendar/conf"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	MyRedis conf.Redis
	MySql   conf.Db
}
