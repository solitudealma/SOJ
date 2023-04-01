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

type GetSolutionInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSolutionInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSolutionInfoLogic {
	return &GetSolutionInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSolutionInfoLogic) GetSolutionInfo(req *types.GetSolutionInfoReq) (resp *types.GetSolutionInfoResp, err error) {
	solution, err := l.svcCtx.SolutionModel.FindOne(l.ctx, req.SolutionId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetSolutionInfo failed, db err: %v", err)
	}

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("该题解信息不存在"), "can't find the solutionInfo, err: %v", err)
	}

	solutionInfo := types.SolutionInfo{
		ProblemId:         solution.ProblemId,
		Title:             solution.Title,
		ProblemSource:     solution.ProblemSource,
		Content:           solution.Content,
		ProblemLink:       solution.ProblemLink,
		UpdateTime:        solution.UpdateTime,
		ProblemDifficulty: solution.ProblemDifficulty,
		AuthorId:          solution.AuthorId,
		AuthorName:        solution.AuthorName,
		AuthorAvatar:      solution.AuthorAvatar,
		Read:              solution.Read,
	}

	return &types.GetSolutionInfoResp{
		SolutionInfo: solutionInfo,
	}, nil
}
