package user

import (
	"context"

	"github.com/solitudealma/SOJ/backend/app/user/cmd/api/internal/svc"
	"github.com/solitudealma/SOJ/backend/app/user/cmd/api/internal/types"
	"github.com/solitudealma/SOJ/backend/app/user/model"
	"github.com/solitudealma/SOJ/backend/common/globalkey"
	"github.com/solitudealma/SOJ/backend/common/tool"
	"github.com/solitudealma/SOJ/backend/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(req *types.LogoutReq) (resp *types.LogoutResp, err error) {
	userId := tool.GetUserId(l.svcCtx.Config.Auth.AccessSecret, l.ctx)

	user, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetUserInfo find user db err , id:%d , err:%v", user.Id, err)
	}

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("该用户信息不存在"), "注销登录失败")
	}

	cacheKey := globalkey.CacheUserTokenKey + user.Username
	if _, err = l.svcCtx.RedisClient.DelCtx(l.ctx, cacheKey); err != nil {
		l.Logger.Errorf("用户注销登录失败, err: %s", err.Error())

		return
	}

	// 无论是否删除失败(指有无key)都直接当作注销登录成功，因为此时redis已经没有key
	return
}
