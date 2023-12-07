package accountlogsvclogic

import (
	"context"
	"time"

	"github.com/aheadIV/NightVoyager/accountsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/accountsvc/model"
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
	if in != nil {
		data := in.GetData()
		err := l.svcCtx.DB.Insert(l.ctx, &model.AccountLog{
			AccountID: int(data.AccountId),
			IpAddr:    data.IpAddr,
			Device:    data.Device,
			CreatedAt: time.Now(),
		})
		if err != nil {
			return nil, err
		}
	}
	return &accountsvc.InsertLogResp{}, nil
}
