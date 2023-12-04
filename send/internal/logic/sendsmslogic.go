package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aheadIV/NightVoyager/appsvc/types/appsvc"
	"github.com/aheadIV/NightVoyager/send/internal/svc"
	"github.com/aheadIV/NightVoyager/send/internal/types"
	"github.com/aheadIV/NightVoyager/send/model"
	"github.com/nats-io/nats.go/jetstream"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	StreamName     = "REVIEWS"
	StreamSubjects = "REVIEWS.*"

	SubjectNameReviewCreated  = "REVIEWS.rateGiven"
	SubjectNameReviewAnswered = "REVIEWS.rateAnswer"
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
	fmt.Println(111122222)
	_, err = l.svcCtx.JsStream.CreateStream(l.ctx, jetstream.StreamConfig{
		Name:     StreamName,
		Subjects: []string{StreamSubjects},
	})

	if err != nil {
		return nil, err
	}

	data := &model.Review{
		To:      req.To,
		Content: req.Content,
		SendAt:  time.Now(),
	}
	fmt.Println("xxxxxxxxxxxx")
	reviewString, _ := json.Marshal(data)
	_, err = l.svcCtx.JsStream.Publish(l.ctx, SubjectNameReviewCreated, reviewString)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Publisher  =>  Message:%s\n", data.Content)
	}

	return
}

// golang nats jetstream produce
// https://medium.com/vlmedia-tech/distributed-message-streaming-in-golang-using-nats-jetstream-29f28be66dc6
//https://natsdashboard.com/jetstream?url=http%3A%2F%2Flocalhost%3A8222
