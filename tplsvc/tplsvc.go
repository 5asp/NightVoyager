package main

import (
	"flag"
	"fmt"

	"github.com/aheadIV/NightVoyager/tplsvc/internal/config"
	tplsendsvcServer "github.com/aheadIV/NightVoyager/tplsvc/internal/server/tplsendsvc"
	tplsignsvcServer "github.com/aheadIV/NightVoyager/tplsvc/internal/server/tplsignsvc"
	tplsvcServer "github.com/aheadIV/NightVoyager/tplsvc/internal/server/tplsvc"
	"github.com/aheadIV/NightVoyager/tplsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/tplsvc/types/tplsvc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/tplsvc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		tplsvc.RegisterTplSvcServer(grpcServer, tplsvcServer.NewTplSvcServer(ctx))
		tplsvc.RegisterTplSignSvcServer(grpcServer, tplsignsvcServer.NewTplSignSvcServer(ctx))
		tplsvc.RegisterTplSendSvcServer(grpcServer, tplsendsvcServer.NewTplSendSvcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
