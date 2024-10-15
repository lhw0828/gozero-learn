package account

import (
	"context"
	"time"
	"user-api/internal/biz"
	"user-api/internal/model"

	"user-api/internal/svc"
	"user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	userModel := model.NewUserModel(l.svcCtx.Conn)

	u, err := userModel.FindByUsernameAndPwd(l.ctx, req.Username, req.Password)

	if err != nil {
		l.Logger.Error("登录失败，", err)
		return nil, biz.DBError
	}

	if u == nil {
		return nil, biz.NameOrPwdError
	}

	// 登录成功 生成token
	secret := l.svcCtx.Config.Auth.Secret
	expire := l.svcCtx.Config.Auth.Expire

	token, err := biz.GetJwtToken(secret, time.Now().Unix(), expire, u.Id)

	if err != nil {
		l.Logger.Error("生成token失败", err)
		return nil, biz.TokenError
	}

	resp = &types.LoginResp{
		Token: token,
	}

	return
}
