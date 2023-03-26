package comment

import (
	"context"

	"github.com/pkg/errors"
	"github.com/solitudealma/SOJ/backend/app/comment/cmd/api/internal/svc"
	"github.com/solitudealma/SOJ/backend/app/comment/cmd/api/internal/types"
	"github.com/solitudealma/SOJ/backend/app/comment/model"
	"github.com/solitudealma/SOJ/backend/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentListInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListInfoLogic {
	return &GetCommentListInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentListInfoLogic) GetCommentListInfo(req *types.GetCommentListInfoReq) (resp *types.GetCommentListInfoResp, err error) {
	list, err := l.svcCtx.CommentModel.FindListBySId(l.ctx, req.SolutionId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetCommentListInfo failed, SolutionId: %d, db err: %v", req.SolutionId, err)
	}

	total, err := l.svcCtx.CommentModel.FindCountBySid(l.ctx)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "FindCountBySid failed, SolutionId: %d, db err: %v", req.SolutionId, err)
	}

	comments := make([]types.Comment, 0)
	if len(list) > 0 {
		for i := range list {

			comment := types.Comment{
				CreateTime:    list[i].CreateTime,
				Layer:         list[i].Layer,
				FromUserId:    list[i].FromUid,
				FromName:      list[i].FromName,
				FromAvatar:    list[i].FromAvatar,
				Content:       list[i].Content,
				CommentId:     list[i].Id,
				RootCommentId: list[i].RootCid,
				ToCommentId:   list[i].ToCid,
				ToUserId:      list[i].ToUid,
				ToName:        list[i].ToName,
			}
			comments = append(comments, comment)
		}
	}

	return &types.GetCommentListInfoResp{
		Comments: comments,
		Total:    total,
	}, nil
}
