package logic

import (
	"context"

	"github.com/1211ciel/fall/test/go-zero/78/gateway/internal/svc"
	"github.com/tal-tech/go-zero/core/logx"
)

type TestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) TestLogic {
	return TestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestLogic) Test() error {
	// todo: add your logic here and delete this line

	return nil
}
