package tplsvclogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/tplsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/tplsvc/types/tplsvc"

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

func (l *DeleteTplLogic) DeleteTpl(in *tplsvc.DeleteTplReq) (*tplsvc.DeletTplResp, error) {
	// todo: add your logic here and delete this line

	return &tplsvc.DeletTplResp{}, nil
}
