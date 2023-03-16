package problem

import (
	"context"

	"github.com/pkg/errors"
	"github.com/solitudealma/SOJ/backend/app/problem/cmd/api/internal/svc"
	"github.com/solitudealma/SOJ/backend/app/problem/cmd/api/internal/types"
	"github.com/solitudealma/SOJ/backend/app/user/model"
	"github.com/solitudealma/SOJ/backend/common/xerr"

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
	problem, err := l.svcCtx.ProblemModel.FindOneByProblemId(l.ctx, req.ProblemId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetProblemInfo failed, db err: %v", err)
	}

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("该题目信息不存在"), "can't find the userInfo, err: %v", err)
	}

	problemInfo := types.ProblemInfo{
		ProblemId:           problem.ProblemId,
		Title:               problem.Title,
		Description:         problem.Description,
		Input:               problem.Input,
		Output:              problem.Output,
		Examples:            problem.Examples,
		Difficulty:          problem.Difficulty,
		TimeLimit:           problem.TimeLimit,
		MemoryLimit:         problem.MemoryLimit,
		TotalNumberOfPasses: 1,
		TotalAttempts:       1,
		Source:              problem.Source,
		Tags:                "dp,贪心",
	}

	return &types.GetProblemInfoResp{
		ProblemInfo: problemInfo,
	}, nil
}
