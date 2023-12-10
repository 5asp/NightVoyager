package tplsvclogic

import (
	"context"
	"time"

	"github.com/aheadIV/NightVoyager/pkg/utils"
	"github.com/aheadIV/NightVoyager/sendsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/sendsvc/model"
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
	if in != nil {
		tpl := &model.Template{
			Status:     1,
			CreatedAt:  time.Now(),
			AccountID:  int(in.GetAccountId()),
			SignID:     int(in.GetSignId()),
			TplContent: in.GetContent(),
			QueueID:    utils.CreateID(),
		}
		err := l.svcCtx.DB.Insert(l.ctx, tpl)
		if err != nil {
			return nil, err
		}
		return &sendsvc.CreateTplResp{TplId: tpl.ID}, nil
	}

	return &sendsvc.CreateTplResp{}, nil
}
