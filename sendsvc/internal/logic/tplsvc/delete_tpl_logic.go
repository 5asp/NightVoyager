package tplsvclogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/sendsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/sendsvc/types/sendsvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTplLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTplLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTplLogic {
	return &DeleteTplLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteTplLogic) DeleteTpl(in *sendsvc.DeleteTplReq) (*sendsvc.DeletTplResp, error) {
	// todo: add your logic here and delete this line

	return &sendsvc.DeletTplResp{}, nil
}
