package accountsvclogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/accountsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/accountsvc/types/accountsvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByIdLogic {
	return &GetByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetByIdLogic) GetById(in *accountsvc.GetByIdReq) (*accountsvc.GetByIdResp, error) {
	// todo: add your logic here and delete this line

	return &accountsvc.GetByIdResp{}, nil
}
