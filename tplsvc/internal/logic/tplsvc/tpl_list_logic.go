package tplsvclogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/tplsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/tplsvc/types/tplsvc"

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

func (l *TplListLogic) TplList(in *tplsvc.TplListReq) (*tplsvc.TplListResp, error) {
	// todo: add your logic here and delete this line

	return &tplsvc.TplListResp{}, nil
}
