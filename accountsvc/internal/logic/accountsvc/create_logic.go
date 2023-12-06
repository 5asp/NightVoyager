package accountsvclogic

import (
	"context"
	"time"

	"github.com/aheadIV/NightVoyager/accountsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/accountsvc/model"
	"github.com/aheadIV/NightVoyager/accountsvc/types/accountsvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *accountsvc.CreateReq) (*accountsvc.CreateResp, error) {
	if in != nil {
		data := &model.Account{
			Account:   in.Data.Account,
			Password:  in.Data.Password,
			Status:    1,
			CreatedAt: time.Now(),
		}
		err := l.svcCtx.DB.Insert(l.ctx, data)
		if err != nil {
			return nil, err
		}
		return &accountsvc.CreateResp{Id: int64(data.ID)}, nil
	}
	return &accountsvc.CreateResp{}, nil
}
