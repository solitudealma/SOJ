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

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var (
	ErrGetTokenError       = xerr.NewErrMsg("获取token失败")
	ErrSetLoginStatusError = xerr.NewErrMsg("设置登录状态失败")
	ErrGenerateTokenError  = xerr.NewErrMsg("生成token失败")
	ErrUsernamePwdError    = xerr.NewErrMsg("账号或密码不正确")
)

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "根据用户名查询信息失败，err: %v, username:%s", err, req.Username)
	}

	if user == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_NOT_FOUND), "登录用户不存在或已注销")
	}

	if match, _ := tool.ComparePasswordAndHash(req.Password, user.Password); !match {
		return nil, errors.Wrapf(ErrUsernamePwdError, "密码错误")
	}

	token, err := l.TokenNext(user)
	if err != nil || token == "" {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.TOKEN_GENERATE_ERROR), "登录失败，请稍后重试")
	}

	return &types.LoginResp{
		Token: token,
		UserInfo: types.UserInfo{
			Id:       user.Id,
			Username: user.Username,
			Avatar:   user.Avatar,
		},
	}, nil
}

// TokenNext 登录以后签发jwt
func (l *LoginLogic) TokenNext(user *model.User) (token string, err error) {

	// 唯一签名
	j := &tool.JWT{SigningKey: []byte(l.svcCtx.Config.Auth.AccessSecret)}
	tokenExpireTime := l.svcCtx.Config.Auth.AccessExpire
	claims := j.CreateClaims(
		l.svcCtx.Config.Auth.BufferTime,
		l.svcCtx.Config.Auth.AccessExpire,
		tool.BaseClaims{
			UserId:   user.Id,
			Username: user.Username,
		},
	)

	token, err = j.CreateToken(claims)
	if err != nil {
		l.Logger.Error("获取token失败! err: %s", err.Error())

		return "", errors.Wrapf(ErrGetTokenError, "Get token failed!")
	}

	cacheKey := globalkey.CacheUserTokenKey + user.Username
	if _, err := l.svcCtx.RedisClient.GetCtx(l.ctx, cacheKey); err == nil {
		if err := l.svcCtx.RedisClient.SetexCtx(l.ctx, cacheKey, token, int(tokenExpireTime)); err != nil {
			l.Logger.Error("设置登录状态失败! err: %s", err.Error())

			return "", errors.Wrapf(ErrSetLoginStatusError, "设置登录状态失败!")
		}

		return token, nil
	} else if err != nil {
		l.Logger.Error("设置登录状态失败! err: %s", err.Error())

		return "", errors.Wrapf(ErrSetLoginStatusError, "the login status fail to set")
	} else {
		if _, err := l.svcCtx.RedisClient.DelCtx(l.ctx, cacheKey); err != nil {
			l.Logger.Errorf("jwt作废失败， err: %s", err.Error())

			return "", errors.Wrapf(xerr.NewErrMsg("jwt作废失败"), "jwt pull back failed")
		}

		if err = l.svcCtx.RedisClient.SetexCtx(l.ctx, cacheKey, token, int(tokenExpireTime)); err != nil {
			l.Logger.Errorf("设置登录状态失败, err: %s", err.Error())
			return "", errors.Wrapf(ErrSetLoginStatusError, "the login status fail to set")
		}

		return token, nil
	}
}
