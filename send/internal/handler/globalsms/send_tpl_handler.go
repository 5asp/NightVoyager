package globalsms

import (
	"net/http"

	"github.com/aheadIV/NightVoyager/send/internal/logic/globalsms"
	"github.com/aheadIV/NightVoyager/send/internal/svc"
	"github.com/aheadIV/NightVoyager/send/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
)

func SendTplHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.XsendReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := globalsms.NewSendTplLogic(r.Context(), svcCtx)
		resp, err := l.SendTpl(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
