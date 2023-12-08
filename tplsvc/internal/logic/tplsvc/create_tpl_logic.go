package tplsvclogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/tplsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/tplsvc/types/tplsvc"

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

func (l *CreateTplLogic) CreateTpl(in *tplsvc.CreateTplReq) (*tplsvc.CreateTplResp, error) {
	// todo: add your logic here and delete this line

	return &tplsvc.CreateTplResp{}, nil
}
