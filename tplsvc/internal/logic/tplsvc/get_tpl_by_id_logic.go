package tplsvclogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/tplsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/tplsvc/types/tplsvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTplByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTplByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTplByIdLogic {
	return &GetTplByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTplByIdLogic) GetTplById(in *tplsvc.GetTplByIdReq) (*tplsvc.GetTplByIdResp, error) {
	// todo: add your logic here and delete this line

	return &tplsvc.GetTplByIdResp{}, nil
}
