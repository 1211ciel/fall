package handler

import (
	"net/http"

	"github.com/1211ciel/fall/test/go-zero/77/gateway/internal/logic"
	"github.com/1211ciel/fall/test/go-zero/77/gateway/internal/svc"
	"github.com/1211ciel/fall/test/go-zero/77/gateway/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GatewayHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGatewayLogic(r.Context(), ctx)
		resp, err := l.Gateway(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
