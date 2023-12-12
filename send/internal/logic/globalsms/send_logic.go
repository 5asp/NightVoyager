package globalsms

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/aheadIV/NightVoyager/send/internal/svc"
	"github.com/aheadIV/NightVoyager/send/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendLogic {
	return &SendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendLogic) Send(req *types.SendReq) (resp *types.SendResp, err error) {
	//1.调用appsvc获取到accountID
	//2.调用tplsvc插入到sign
	//3.调用tplsvc插入到tpl
	//4.调用tplsvc插入到tpl_send,返回send_id 以及处理账户额度
	//5.调用appsvc发送短信

	openBracket := "【"
	closeBracket := "】"

	startIndex := strings.Index(req.Content, openBracket)
	endIndex := strings.Index(req.Content, closeBracket)

	if startIndex == -1 || endIndex == -1 {
		// 未找到匹配的字符串
		fmt.Println("未找到匹配的字符串")
		notMatch := errors.New("非法签名")
		return nil, notMatch
	}

	sig := req.Content[startIndex+len(openBracket) : endIndex]
	tpl := strings.TrimSpace(req.Content[endIndex+len(closeBracket):])

	return
}
