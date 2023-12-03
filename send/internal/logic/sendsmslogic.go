package logic

import (
	"context"

	"github.com/aheadIV/NightVoyager/appsvc/types/appsvc"
	"github.com/aheadIV/NightVoyager/send/internal/svc"
	"github.com/aheadIV/NightVoyager/send/internal/types"
	"github.com/nats-io/nats.go/jetstream"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSMSLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendSMSLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSMSLogic {
	return &SendSMSLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendSMSLogic) SendSMS(req *types.SendReq) (resp *types.SendResp, err error) {
	// todo: add your logic here and delete this line
	if req.AppID != 0 && req.Secret != "" {
		data, _ := l.svcCtx.AppRPC.GetApp(l.ctx, &appsvc.GetAppReq{AppId: int64(req.AppID)})
		if data.Secret != req.Secret {
			return nil, err
		}

	}
	CreateJetStream(l.ctx, l.svcCtx.JsStream, "ok", "sms.*")
	return
}

// golang nats jetstream produce
// https://medium.com/vlmedia-tech/distributed-message-streaming-in-golang-using-nats-jetstream-29f28be66dc6
func CreateJetStream(ctx context.Context, js jetstream.JetStream, streamName, subjectName string) {
	stream, _ := js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     streamName,
		Subjects: []string{subjectName},
	})

}
