package tplsvclogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/sendsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/sendsvc/types/sendsvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTplLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTplLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTplLogic {
	return &UpdateTplLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTplLogic) UpdateTpl(in *sendsvc.UpdateTplReq) (*sendsvc.UpdateTplResp, error) {
	// todo: add your logic here and delete this line

	return &sendsvc.UpdateTplResp{}, nil
}
