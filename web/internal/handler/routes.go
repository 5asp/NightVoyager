// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/aheadIV/NightVoyager/web/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/app/create",
				Handler: CreateAppHandler(serverCtx),
			},
		},
		rest.WithMaxBytes(1048576),
	)
}
