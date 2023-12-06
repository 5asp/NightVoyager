package sig

import (
	"net/http"

	"github.com/aheadIV/NightVoyager/web/internal/logic/sig"
	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SigLsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SigLsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sig.NewSigLsLogic(r.Context(), svcCtx)
		resp, err := l.SigLs(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
