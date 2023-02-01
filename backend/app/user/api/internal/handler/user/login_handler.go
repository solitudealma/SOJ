package user

import (
	"net/http"

	"github.com/solitudealma/SOJ/backend/app/user/api/internal/logic/user"
	"github.com/solitudealma/SOJ/backend/app/user/api/internal/svc"
	"github.com/solitudealma/SOJ/backend/app/user/api/internal/types"
	"github.com/solitudealma/SOJ/backend/common/result"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		if err := svcCtx.Validator.Struct(&req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		result.HttpResult(r, w, resp, err)
	}
}
