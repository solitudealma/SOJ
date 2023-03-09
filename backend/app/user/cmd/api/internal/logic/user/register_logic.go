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
	"gorm.io/gorm"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var (
	ErrPasswordError = xerr.NewErrMsg("两次输入的密码不一致")
)

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	u, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "username:%s, err:%+v", req.Username, err)
	}

	if u != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_HAS_REGISTERED), "the username has been registered")
	}

	user := &model.User{
		Username: req.Username,
	}

	password, err := tool.GenerateFromPassword(req.Password)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("注册失败"), "密码摘要失败")
	}

	user.Password = password
	err = l.svcCtx.UserModel.Transaction(l.ctx, func(db *gorm.DB) error {
		err := l.svcCtx.UserModel.Insert(l.ctx, db, user)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("注册失败"), "register failed, because of the mysql database")
	}

	return &types.RegisterResp{}, nil
}
