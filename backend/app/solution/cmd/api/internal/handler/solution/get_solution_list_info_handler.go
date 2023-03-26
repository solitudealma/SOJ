package solution

import (
	"net/http"

	"github.com/solitudealma/SOJ/backend/app/solution/cmd/api/internal/logic/solution"
	"github.com/solitudealma/SOJ/backend/app/solution/cmd/api/internal/svc"
	"github.com/solitudealma/SOJ/backend/app/solution/cmd/api/internal/types"
	"github.com/solitudealma/SOJ/backend/common/result"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetSolutionListInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetSolutionListInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		if err := svcCtx.Validator.Struct(&req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := solution.NewGetSolutionListInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetSolutionListInfo(&req)
		result.HttpResult(r, w, resp, err)
	}
}
