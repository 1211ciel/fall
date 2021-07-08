package svc

import (
	"github.com/1211ciel/fall/test/go-zero/78/gateway/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
