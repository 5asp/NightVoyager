package app

import (
	"net/http"

	"github.com/aheadIV/NightVoyager/pkg/response"
	"github.com/aheadIV/NightVoyager/web/internal/logic/app"
	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UdAppHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AppUdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := app.NewUdAppLogic(r.Context(), svcCtx)
		resp, err := l.UdApp(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			response.Response(w, resp, err)
		}
	}
}
