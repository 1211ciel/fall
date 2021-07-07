package handler

import (
	"net/http"

	"github.com/1211ciel/fall/test/go-zero/77/gateway/internal/logic"
	"github.com/1211ciel/fall/test/go-zero/77/gateway/internal/svc"
	"github.com/1211ciel/fall/test/go-zero/77/gateway/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func RegisterHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), ctx)
		resp, err := l.Register(req)
		if err != nil {
			httpx.OkJson(w, err.Error())
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
