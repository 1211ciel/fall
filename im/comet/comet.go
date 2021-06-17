package main

import (
	"flag"
	"fmt"
	"github.com/1211ciel/fall/im/comet/cometpb"
	"github.com/1211ciel/fall/im/comet/internal"
	"github.com/1211ciel/fall/im/comet/internal/config"
	"github.com/1211ciel/fall/im/comet/internal/handler"
	"github.com/1211ciel/fall/im/comet/internal/server"
	"github.com/1211ciel/fall/im/comet/internal/svc"
	"github.com/tal-tech/go-zero/rest"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/comet.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewCometCliServer(ctx)

	step := 0 // 启动ws服务
	hub := internal.NewHub(ctx)
	go hub.Run()

	step = 1 // 启动http服务
	go func() {
		httpServer := rest.MustNewServer(c.RestConf)
		handler.RegisterHandlers(httpServer, ctx, hub)
		fmt.Printf("Starting http  httpServer at %s:%d...\n", c.Host, c.Port)
		httpServer.Start()
		defer httpServer.Stop()
	}()

	step = 2 // 启动rpc服务
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		cometpb.RegisterCometCliServer(grpcServer, srv)
	})
	defer s.Stop()
	fmt.Printf("Starting rpc httpServer at %s...\n", c.ListenOn)
	s.Start()
	step++
}
