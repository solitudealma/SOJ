package problem

import (
	"net/http"

	"github.com/solitudealma/SOJ/backend/app/problem/cmd/api/internal/logic/problem"
	"github.com/solitudealma/SOJ/backend/app/problem/cmd/api/internal/svc"
	"github.com/solitudealma/SOJ/backend/app/problem/cmd/api/internal/types"
	"github.com/solitudealma/SOJ/backend/common/result"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetProblemListInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetProblemListInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		if err := svcCtx.Validator.Struct(&req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := problem.NewGetProblemListInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetProblemListInfo(&req)
		result.HttpResult(r, w, resp, err)
	}
}
