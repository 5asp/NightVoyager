package main

import (
	"flag"
	"fmt"

	"github.com/aheadIV/NightVoyager/addrsvc/addrsvc"
	"github.com/aheadIV/NightVoyager/addrsvc/internal/config"
	"github.com/aheadIV/NightVoyager/addrsvc/internal/server"
	"github.com/aheadIV/NightVoyager/addrsvc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/addrsvc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		addrsvc.RegisterAddrsvcServer(grpcServer, server.NewAddrsvcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}