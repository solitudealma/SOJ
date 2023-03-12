package problem

import (
	"context"

	"github.com/solitudealma/SOJ/backend/app/problem/cmd/api/internal/svc"
	"github.com/solitudealma/SOJ/backend/app/problem/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProblemInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProblemInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProblemInfoLogic {
	return &GetProblemInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProblemInfoLogic) GetProblemInfo(req *types.GetProblemInfoReq) (resp *types.GetProblemInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
