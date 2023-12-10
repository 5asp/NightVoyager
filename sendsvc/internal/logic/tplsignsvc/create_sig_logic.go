package tplsignsvclogic

import (
	"context"
	"time"

	"github.com/aheadIV/NightVoyager/sendsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/sendsvc/model"
	"github.com/aheadIV/NightVoyager/sendsvc/types/sendsvc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSigLogic {
	return &CreateSigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateSigLogic) CreateSig(in *sendsvc.CreateSigReq) (*sendsvc.CreateSigResp, error) {
	if in != nil {
		sig := &model.TemplateSign{
			Status:    1,
			CreatedAt: time.Now(),
			Signature: in.GetSignature(),
			AccountID: int(in.GetAccountId()),
		}
		err := l.svcCtx.DB.Insert(l.ctx, sig)
		if err != nil {
			return nil, err
		}
		return &sendsvc.CreateSigResp{Id: int64(sig.ID)}, nil
	}
	return &sendsvc.CreateSigResp{}, nil
}
