package logic

import (
	"context"

	"github.com/aheadIV/NightVoyager/appsvc/types/appsvc"
	"github.com/aheadIV/NightVoyager/send/internal/svc"
	"github.com/aheadIV/NightVoyager/send/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSMSLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendSMSLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSMSLogic {
	return &SendSMSLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendSMSLogic) SendSMS(req *types.SendReq) (resp *types.SendResp, err error) {
	// todo: add your logic here and delete this line
	if req.AppID != 0 && req.Secret != "" {
		data, _ := l.svcCtx.AppRPC.GetApp(l.ctx, &appsvc.GetAppReq{AppId: int64(req.AppID)})
		if data.Secret != req.Secret {

		}
	}
	return
}
