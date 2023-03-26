package comment

import (
	"net/http"

	"github.com/solitudealma/SOJ/backend/app/comment/cmd/api/internal/logic/comment"
	"github.com/solitudealma/SOJ/backend/app/comment/cmd/api/internal/svc"
	"github.com/solitudealma/SOJ/backend/app/comment/cmd/api/internal/types"
	"github.com/solitudealma/SOJ/backend/common/result"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCommentListInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCommentListInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		if err := svcCtx.Validator.Struct(&req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := comment.NewGetCommentListInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetCommentListInfo(&req)
		result.HttpResult(r, w, resp, err)
	}
}
