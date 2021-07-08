package handler

import (
	"net/http"

	"github.com/1211ciel/fall/test/go-zero/78/gateway/internal/logic/user"
	"github.com/1211ciel/fall/test/go-zero/78/gateway/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func TestHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewTestLogic(r.Context(), ctx)
		err := l.Test()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
