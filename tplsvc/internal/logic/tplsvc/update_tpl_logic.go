package tplsvclogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/tplsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/tplsvc/types/tplsvc"

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

func (l *UpdateTplLogic) UpdateTpl(in *tplsvc.UpdateTplReq) (*tplsvc.UpdateTplResp, error) {
	// todo: add your logic here and delete this line

	return &tplsvc.UpdateTplResp{}, nil
}
