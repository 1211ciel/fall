package logic

import (
	"context"
	"github.com/1211ciel/fall/test/go-zero/77/service/user/user"

	"github.com/1211ciel/fall/test/go-zero/77/gateway/internal/svc"
	"github.com/1211ciel/fall/test/go-zero/77/gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.RegisterReq) (*types.RegisterRes, error) {
	register, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterReq{Uname: req.Name, Pwd: req.Pwd})
	if err != nil {
		return nil, err
	}
	return &types.RegisterRes{Ok: register.Ok}, nil
}
