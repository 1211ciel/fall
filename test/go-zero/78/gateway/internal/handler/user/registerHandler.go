package handler

import (
	"net/http"

	"github.com/1211ciel/fall/test/go-zero/78/gateway/internal/logic/user"
	"github.com/1211ciel/fall/test/go-zero/78/gateway/internal/svc"
	"github.com/1211ciel/fall/test/go-zero/78/gateway/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func RegisterHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegistereReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), ctx)
		resp, err := l.Register(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
