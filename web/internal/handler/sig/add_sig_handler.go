package sig

import (
	"net/http"

	"github.com/aheadIV/NightVoyager/web/internal/logic/sig"
	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddSigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SigAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sig.NewAddSigLogic(r.Context(), svcCtx)
		resp, err := l.AddSig(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
