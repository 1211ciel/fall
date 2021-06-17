package handler

import (
	"github.com/1211ciel/fall/im/comet/internal"
	"github.com/1211ciel/fall/im/comet/internal/svc"
	"github.com/tal-tech/go-zero/rest"
	"net/http"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext, hub *internal.Hub) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/ws",
				Handler: wsHandler(serverCtx, hub),
			},
		},
	)
}
