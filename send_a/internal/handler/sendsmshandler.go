package handler

import (
	"net/http"

	"github.com/aheadIV/NightVoyager/send/internal/logic"
	"github.com/aheadIV/NightVoyager/send/internal/svc"
	"github.com/aheadIV/NightVoyager/send/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SendSMSHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSendSMSLogic(r.Context(), svcCtx)
		resp, err := l.SendSMS(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
