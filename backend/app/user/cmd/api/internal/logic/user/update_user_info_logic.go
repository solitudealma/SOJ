package user

import (
	"context"

	"github.com/pkg/errors"
	"github.com/solitudealma/SOJ/backend/app/user/cmd/api/internal/svc"
	"github.com/solitudealma/SOJ/backend/app/user/cmd/api/internal/types"
	"github.com/solitudealma/SOJ/backend/app/user/model"
	"github.com/solitudealma/SOJ/backend/common/tool"
	"github.com/solitudealma/SOJ/backend/common/xerr"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(req *types.UpdateUserInfoReq) (resp *types.UpdateUserInfoResp, err error) {
	userId := tool.GetUserId(l.svcCtx.Config.Auth.AccessSecret, l.ctx)

	user, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetUserInfo find user db err , id:%d , err:%v", user.Id, err)
	}

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("该用户信息不存在"), "can't find the userInfo")
	}

	user.Username = req.Username
	// 需要同步其他数据库的修改
	err = l.svcCtx.UserModel.Transaction(l.ctx, func(db *gorm.DB) error {
		err := l.svcCtx.UserModel.Update(l.ctx, db, user)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("更新用户信息失败"), "update userInfo failed, because of the mysql database, err:%v", err)
	}

	return &types.UpdateUserInfoResp{}, nil
}
