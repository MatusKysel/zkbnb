package account

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/bnb-chain/zkbas/service/apiserver/internal/logic/account"
	"github.com/bnb-chain/zkbas/service/apiserver/internal/svc"
	"github.com/bnb-chain/zkbas/service/apiserver/internal/types"
)

func GetAccountHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqGetAccount
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := account.NewGetAccountLogic(r.Context(), svcCtx)
		resp, err := l.GetAccount(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
