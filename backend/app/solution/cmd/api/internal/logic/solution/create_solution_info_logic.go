package solution

import (
	"context"

	"github.com/pkg/errors"
	"github.com/solitudealma/SOJ/backend/app/solution/cmd/api/internal/svc"
	"github.com/solitudealma/SOJ/backend/app/solution/cmd/api/internal/types"
	"github.com/solitudealma/SOJ/backend/app/solution/model"
	"github.com/solitudealma/SOJ/backend/common/xerr"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSolutionInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSolutionInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSolutionInfoLogic {
	return &CreateSolutionInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSolutionInfoLogic) CreateSolutionInfo(req *types.CreateSolutionInfoReq) (resp *types.CreateSolutionInfoResp, err error) {
	s, err := l.svcCtx.SolutionModel.FindOneByAuthorId(l.ctx, req.AuthorId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "authorId:%d, err:%+v", req.AuthorId, err)
	}

	if err == nil {
		s.ProblemId = req.ProblemId
		s.ProblemSource = req.ProblemSource
		s.ProblemDifficulty = req.ProblemDifficulty
		s.Title = req.Title
		s.Content = req.Content
		s.ProblemLink = req.ProblemLink

		if req.Type == 0 {
			s.Type = 0
		} else {
			s.Type = 1
		}

		err = l.svcCtx.SolutionModel.Transaction(l.ctx, func(db *gorm.DB) error {
			err := l.svcCtx.SolutionModel.Update(l.ctx, db, s)
			if err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrMsg("保存失败"), "save solution failed, because of the mysql database, err:%+v", err)
		}

		return &types.CreateSolutionInfoResp{
			SolutionId: s.Id,
		}, nil
	}

	solution := &model.Solution{
		ProblemId:         req.ProblemId,
		ProblemSource:     req.ProblemSource,
		ProblemDifficulty: req.ProblemDifficulty,
		Title:             req.Title,
		AuthorId:          req.AuthorId,
		AuthorName:        req.AuthorName,
		AuthorAvatar:      req.AuthorAvatar,
		Content:           req.Content,
		ProblemLink:       req.ProblemLink,
	}

	err = l.svcCtx.SolutionModel.Transaction(l.ctx, func(db *gorm.DB) error {
		err := l.svcCtx.SolutionModel.Insert(l.ctx, db, solution)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("提交失败"), "sumit solution failed, because of the mysql database, error: %v", err)
	}

	return &types.CreateSolutionInfoResp{}, nil
}
