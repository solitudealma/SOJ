package comment

import (
	"context"

	"github.com/pkg/errors"
	"github.com/solitudealma/SOJ/backend/app/comment/cmd/api/internal/svc"
	"github.com/solitudealma/SOJ/backend/app/comment/cmd/api/internal/types"
	"github.com/solitudealma/SOJ/backend/app/comment/model"
	"github.com/solitudealma/SOJ/backend/common/xerr"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCommentLogic) CreateComment(req *types.CreateCommentReq) (resp *types.CreateCommentResp, err error) {

	comment := &model.Comment{
		Sid:        req.Comment.SolutionId,
		CreateTime: req.Comment.CreateTime,
		Layer:      req.Comment.Layer,
		FromUid:    req.Comment.FromUserId,
		FromName:   req.Comment.FromName,
		FromAvatar: req.Comment.FromAvatar,
		Content:    req.Comment.Content,
		RootCid:    req.Comment.RootCommentId,
		ToCid:      req.Comment.ToCommentId,
		ToUid:      req.Comment.ToUserId,
		ToName:     req.Comment.ToName,
	}

	l.svcCtx.CommentModel.Transaction(l.ctx, func(db *gorm.DB) error {
		err := l.svcCtx.CommentModel.Insert(l.ctx, db, comment)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("创建评论失败"), "create comment failed, solutionId: %d, err: %v", req.Comment.SolutionId, err)
	}

	return &types.CreateCommentResp{}, nil
}
