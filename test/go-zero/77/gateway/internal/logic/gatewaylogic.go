package logic

import (
	"context"
	"github.com/1211ciel/fall/test/go-zero/77/gateway/internal/svc"
	"github.com/1211ciel/fall/test/go-zero/77/gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GatewayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGatewayLogic(ctx context.Context, svcCtx *svc.ServiceContext) GatewayLogic {
	return GatewayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GatewayLogic) Gateway(req types.Request) (*types.Response, error) {
	return &types.Response{}, nil
}
