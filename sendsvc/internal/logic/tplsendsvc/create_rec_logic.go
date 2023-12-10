package tplsendsvclogic

import (
	"context"
	"time"

	"github.com/aheadIV/NightVoyager/sendsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/sendsvc/model"
	"github.com/aheadIV/NightVoyager/sendsvc/types/sendsvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRecLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRecLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRecLogic {
	return &CreateRecLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateRecLogic) CreateRec(in *sendsvc.CreateRecReq) (*sendsvc.CreateRecResp, error) {
	if in != nil {
		tpl := &model.TemplateSend{
			Mobile:        in.GetMobile(),
			NationCode:    in.GetNationcode(),
			AppID:         int(in.GetAppId()),
			TemplateID:    in.GetTplId(),
			TemplateParam: in.GetTplParam(),
			SenderIP:      in.GetSenderIp(),
			CreatedAt:     time.Now(),
		}
		err := l.svcCtx.DB.Insert(l.ctx, tpl)
		if err != nil {
			return nil, err
		}
		return &sendsvc.CreateRecResp{}, nil
	}

	return &sendsvc.CreateRecResp{}, nil
}
