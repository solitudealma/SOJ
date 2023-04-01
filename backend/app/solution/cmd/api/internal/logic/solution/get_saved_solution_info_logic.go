package solution

import (
	"context"

	"github.com/pkg/errors"
	"github.com/solitudealma/SOJ/backend/app/solution/cmd/api/internal/svc"
	"github.com/solitudealma/SOJ/backend/app/solution/cmd/api/internal/types"
	"github.com/solitudealma/SOJ/backend/app/solution/model"
	"github.com/solitudealma/SOJ/backend/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSavedSolutionInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSavedSolutionInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSavedSolutionInfoLogic {
	return &GetSavedSolutionInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSavedSolutionInfoLogic) GetSavedSolutionInfo(req *types.GetSavedSolutionInfoReq) (resp *types.GetSavedSolutionInfoResp, err error) {
	s, err := l.svcCtx.SolutionModel.FindOneByAuthorId(l.ctx, req.AuthorId)

	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "authorId:%d, err:%+v", req.AuthorId, err)
	}

	if err != nil {
		return nil, nil
	}

	return &types.GetSavedSolutionInfoResp{
		ProblemId:         s.ProblemId,
		ProblemSource:     s.ProblemSource,
		ProblemDifficulty: s.ProblemDifficulty,
		Title:             s.Title,
		Content:           s.Content,
		ProblemLink:       s.ProblemLink,
	}, nil
}
