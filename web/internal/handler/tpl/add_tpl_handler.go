package tpl

import (
	"net/http"

	"github.com/aheadIV/NightVoyager/pkg/response"
	"github.com/aheadIV/NightVoyager/web/internal/logic/tpl"
	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddTplHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TplAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tpl.NewAddTplLogic(r.Context(), svcCtx)
		resp, err := l.AddTpl(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			response.Response(w, resp, err)
		}
	}
}
