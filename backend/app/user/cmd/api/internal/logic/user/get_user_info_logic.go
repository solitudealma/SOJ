package user

import (
	"context"

	"github.com/solitudealma/SOJ/backend/app/user/cmd/api/internal/svc"
	"github.com/solitudealma/SOJ/backend/app/user/cmd/api/internal/types"
	"github.com/solitudealma/SOJ/backend/app/user/model"
	"github.com/solitudealma/SOJ/backend/common/tool"
	"github.com/solitudealma/SOJ/backend/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoReq) (resp *types.GetUserInfoResp, err error) {
	userId := tool.GetUserId(l.svcCtx.Config.Auth.AccessSecret, l.ctx)

	user, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetUserInfo find user db err , id:%d , err:%v", user.Id, err)
	}

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("该用户信息不存在"), "can't find the userInfo")
	}

	return &types.GetUserInfoResp{
		UserInfo: types.UserInfo{
			Id:       user.Id,
			Username: user.Username,
			Avatar:   user.Avatar,
		},
	}, nil
}
