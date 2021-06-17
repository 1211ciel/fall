package handler

import (
	"github.com/1211ciel/fall/im/comet/internal"
	"github.com/1211ciel/fall/im/comet/internal/logic"
	"github.com/1211ciel/fall/im/comet/internal/svc"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func wsHandler(ctx *svc.ServiceContext, hub *internal.Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewWsLogic(r.Context(), ctx)
		err := l.Ws(hub, w, r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}

