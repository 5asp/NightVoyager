package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/aheadIV/NightVoyager/accountsvc/types/accountsvc"
	"github.com/aheadIV/NightVoyager/pkg/utils"
	"github.com/aheadIV/NightVoyager/web/internal/svc"
	"github.com/aheadIV/NightVoyager/web/internal/types"
	"github.com/golang-jwt/jwt"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogLogic {
	return &LogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogLogic) Log(req *types.LogReq) (resp *types.LogResp, err error) {
	res, err := l.svcCtx.AccountRPC.GetByAccount(l.ctx, &accountsvc.GetByAccountReq{Account: req.Account})
	if err != nil {
		return nil, err
	}
	accountData := res.GetData()
	if accountData == nil {
		return &types.LogResp{Err: "帐号不存在"}, nil
	}
	newPass := req.Password + l.svcCtx.Config.Auth.AccessSecret
	if !utils.CheckPassword(res.Data.Password, newPass) {
		return &types.LogResp{Err: "非法帐号"}, nil
	}
	// ip := l.ctx.Value("X-Real-Ip")
	// fmt.Println(ip)
	_, err = l.svcCtx.AccountLogRPC.InsertLog(l.ctx, &accountsvc.InsertLogReq{Data: &accountsvc.AccountLog{
		AccountId: accountData.Id,
		IpAddr:    l.svcCtx.ClientIP,
		Device:    l.svcCtx.Device,
	}})
	if err != nil {
		return nil, err
	}
	var accessExpire = l.svcCtx.Config.Auth.AccessExpire
	jwtMap := make(map[string]interface{}, 0)
	jwtMap["account"] = accountData.Account
	jwtMap["account_id"] = accountData.Id
	fmt.Println(accountData.Id)
	now := time.Now().Unix()
	accessToken, err := l.GenToken(now, l.svcCtx.Config.Auth.AccessSecret, jwtMap, accessExpire)
	if err != nil {
		return nil, err
	}

	return &types.LogResp{Token: accessToken}, nil
}

func (l *LogLogic) GenToken(iat int64, secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	for k, v := range payloads {
		claims[k] = v
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
