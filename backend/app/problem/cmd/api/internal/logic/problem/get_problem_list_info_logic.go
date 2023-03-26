package problem

import (
	"context"
	"math/rand"

	"github.com/solitudealma/SOJ/backend/app/problem/cmd/api/internal/svc"
	"github.com/solitudealma/SOJ/backend/app/problem/cmd/api/internal/types"
	"github.com/solitudealma/SOJ/backend/app/problem/model"
	"github.com/solitudealma/SOJ/backend/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetProblemListInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProblemListInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProblemListInfoLogic {
	return &GetProblemListInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProblemListInfoLogic) GetProblemListInfo(req *types.GetProblemListInfoReq) (resp *types.GetProblemListInfoResp, err error) {

	list, err := l.svcCtx.ProblemModel.FindPageListByPage(l.ctx, req.CurrentPage, 50, "")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetProblemListInfo failed, page: %d db err: %v", req.CurrentPage, err)
	}

	total, err := l.svcCtx.ProblemModel.FindCount(l.ctx)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "FindCount failed, page: %d db err: %v", req.CurrentPage, err)
	}

	problems := make([]types.Problem, 0)
	if len(list) > 0 {
		for i := range list {
			var flag bool
			if i%2 == 0 {
				flag = true
			} else {
				flag = false
			}

			problem := types.Problem{
				ProblemId:   list[i].ProblemId,
				Title:       list[i].Title,
				PassingRate: rand.Float64()*100 + 1,
				Difficulty:  list[i].Difficulty,
				IsAccepted:  flag,
			}

			problems = append(problems, problem)
		}
	}

	return &types.GetProblemListInfoResp{
		Problems: problems,
		Current:  req.CurrentPage,
		Total:    total,
	}, nil
}
