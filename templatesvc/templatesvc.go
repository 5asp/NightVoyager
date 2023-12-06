package main

import (
	"flag"
	"fmt"

	"github.com/aheadIV/NightVoyager/templatesvc/internal/config"
	"github.com/aheadIV/NightVoyager/templatesvc/internal/server"
	"github.com/aheadIV/NightVoyager/templatesvc/internal/svc"
	"github.com/aheadIV/NightVoyager/templatesvc/types/templatesvc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/templatesvc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		templatesvc.RegisterTemplatesvcServer(grpcServer, server.NewTemplatesvcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
