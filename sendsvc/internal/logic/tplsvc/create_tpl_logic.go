package tplsvclogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/sendsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/sendsvc/types/sendsvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTplLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTplLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTplLogic {
	return &CreateTplLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateTplLogic) CreateTpl(in *sendsvc.CreateTplReq) (*sendsvc.CreateTplResp, error) {
	// todo: add your logic here and delete this line

	return &sendsvc.CreateTplResp{}, nil
}
