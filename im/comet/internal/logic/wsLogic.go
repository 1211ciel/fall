package logic

import (
	"context"
	"github.com/1211ciel/fall/im/comet/internal"
	"github.com/1211ciel/fall/im/comet/internal/svc"
	"net/http"

	"github.com/tal-tech/go-zero/core/logx"
)

type WsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWsLogic(ctx context.Context, svcCtx *svc.ServiceContext) WsLogic {
	return WsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WsLogic) Ws(hub *internal.Hub, w http.ResponseWriter, r *http.Request) error {
	internal.ServeWs(hub, w, r)
	return nil
}
