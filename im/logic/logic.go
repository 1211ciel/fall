package main

import (
	"flag"
	"fmt"

	"github.com/1211ciel/fall/im/logic/internal/config"
	"github.com/1211ciel/fall/im/logic/internal/server"
	"github.com/1211ciel/fall/im/logic/internal/svc"
	"github.com/1211ciel/fall/im/logic/logicpb"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/logic.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewLogicCliServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		logicpb.RegisterLogicCliServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
