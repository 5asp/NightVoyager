package logic

import (
	"context"

	"github.com/aheadIV/NightVoyager/accountsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/accountsvc/types/accountsvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByIDLogic {
	return &GetByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetByIDLogic) GetByID(in *accountsvc.Request) (*accountsvc.Response, error) {
	// todo: add your logic here and delete this line

	return &accountsvc.Response{}, nil
}
