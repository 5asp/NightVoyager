package accountlogsvclogic

import (
	"context"

	"github.com/aheadIV/NightVoyager/accountsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/accountsvc/types/accountsvc"

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

func (l *InsertLogLogic) InsertLog(in *accountsvc.InsertLogReq) (*accountsvc.InsertLogResp, error) {
	// if in != nil {
	// 	data := &model.AccountLog{
	// 		AccountID: ,
	// 	}
	// 	err := l.svcCtx.DB.Insert(l.ctx, data)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return &accountsvc.CreateResp{Id: int64(data.ID)}, nil
	// }
	return &accountsvc.InsertLogResp{}, nil
}
