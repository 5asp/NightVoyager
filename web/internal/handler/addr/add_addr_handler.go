package addr

import (
	"net/http"

	"github.com/aheadIV/NightVoyager/web/internal/logic/addr"
	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddAddrHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddrAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := addr.NewAddAddrLogic(r.Context(), svcCtx)
		resp, err := l.AddAddr(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
