package tplsvclogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/sendsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/sendsvc/types/sendsvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type TplListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTplListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TplListLogic {
	return &TplListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TplListLogic) TplList(in *sendsvc.TplListReq) (*sendsvc.TplListResp, error) {
	// todo: add your logic here and delete this line

	return &sendsvc.TplListResp{}, nil
}
