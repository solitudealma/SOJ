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

type GetSolutionListInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSolutionListInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSolutionListInfoLogic {
	return &GetSolutionListInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSolutionListInfoLogic) GetSolutionListInfo(req *types.GetSolutionListInfoReq) (resp *types.GetSolutionListInfoResp, err error) {
	list, err := l.svcCtx.SolutionModel.FindPageListByPage(l.ctx, req.CurrentPage, 50, "")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetSolutionListInfo failed, page: %d db err: %v", req.CurrentPage, err)
	}

	total, err := l.svcCtx.SolutionModel.FindCount(l.ctx)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "FindCount failed, page: %d db err: %v", req.CurrentPage, err)
	}

	solutions := make([]types.Solution, 0)
	if len(list) > 0 {
		for i := range list {
			solution := types.Solution{
				SolutionId:        list[i].Id,
				ProblemId:         list[i].ProblemId,
				Title:             list[i].Title,
				ProblemSource:     list[i].ProblemSource,
				CreateTime:        list[i].CreateTime,
				UpdateTime:        list[i].UpdateTime,
				ProblemDifficulty: list[i].ProblemDifficulty,
				AuthorId:          list[i].AuthorId,
				AuthorName:        list[i].AuthorName,
				AuthorAvatar:      list[i].AuthorAvatar,
				Read:              list[i].Read,
			}

			solutions = append(solutions, solution)
		}
	}

	return &types.GetSolutionListInfoResp{
		Solutions: solutions,
		Current:   req.CurrentPage,
		Total:     total,
	}, nil
}
