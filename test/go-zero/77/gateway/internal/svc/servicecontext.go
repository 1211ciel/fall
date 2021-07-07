package svc

import (
	"github.com/1211ciel/fall/test/go-zero/77/gateway/internal/config"
	"github.com/1211ciel/fall/test/go-zero/77/service/user/userclient"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
