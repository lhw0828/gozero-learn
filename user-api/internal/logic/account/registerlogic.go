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

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	userModel := model.NewUserModel(l.svcCtx.Mysql)

	// 检查用户名是否存在
	u, err := userModel.FindByUsername(l.ctx, req.Username)
	if err != nil {
		l.Logger.Error("注册失败，", err)
		return nil, biz.DBError
	}
	if u != nil {
		return nil, biz.UserIsRegistered
	}

	// 注册
	_, err = userModel.Insert(l.ctx, &model.User{
		Username:      req.Username,
		Password:      req.Password,
		RegisterTime:  time.Now(),
		LastLoginTime: time.Now(),
	})
	if err != nil {
		l.Logger.Error("注册失败，", err)
		return nil, err
	}

	return
}
