package logic

import (
	"context"
	"flag"
	"github.com/1211ciel/fall/test/go-zero/78/service/user/internal/config"
	"github.com/1211ciel/fall/test/go-zero/78/service/user/internal/svc"
	"github.com/1211ciel/fall/test/go-zero/78/service/user/user"
	"github.com/tal-tech/go-zero/core/conf"
	"testing"
)

func TestNewPingLogic(t *testing.T) {
	_, err := NewRegisterLogic(getConf()).Register(&user.RegisterReq{
		Uname: "ciel", Pwd: "123",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}
func getConf() (context.Context, *svc.ServiceContext) {
	path := "../../etc/user.yaml"
	var configFile = flag.String("f", path, "the config file")
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	return context.Background(), ctx
}
