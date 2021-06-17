package logic

import (
	"context"

	"github.com/1211ciel/fall/im/logic/internal/svc"
	"github.com/1211ciel/fall/im/logic/logicpb"

	"github.com/tal-tech/go-zero/core/logx"
)

type ConnectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConnectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConnectLogic {
	return &ConnectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  提供给comet 用户连接时进行鉴权
func (l *ConnectLogic) Connect(in *logicpb.ConnectReq) (*logicpb.ConnectReply, error) {
	// todo: add your logic here and delete this line

	return &logicpb.ConnectReply{}, nil
}
