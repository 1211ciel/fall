package logic

import (
	"context"

	"github.com/1211ciel/fall/im/comet/cometpb"
	"github.com/1211ciel/fall/im/comet/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type PushMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPushMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushMsgLogic {
	return &PushMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PushMsgLogic) PushMsg(in *cometpb.PushReq) (*cometpb.CometEmpty, error) {
	// todo: add your logic here and delete this line

	return &cometpb.CometEmpty{}, nil
}
