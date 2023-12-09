package tplsendsvclogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/sendsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/sendsvc/types/sendsvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRecByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRecByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecByIdLogic {
	return &GetRecByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRecByIdLogic) GetRecById(in *sendsvc.GetRecByIdReq) (*sendsvc.GetRecByIdResp, error) {
	// todo: add your logic here and delete this line

	return &sendsvc.GetRecByIdResp{}, nil
}
