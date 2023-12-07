package sendlogsvclogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/appsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/appsvc/types/appsvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertLogLogic {
	return &InsertLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertLogLogic) InsertLog(in *appsvc.Req) (*appsvc.Resp, error) {
	// todo: add your logic here and delete this line

	return &appsvc.Resp{}, nil
}
