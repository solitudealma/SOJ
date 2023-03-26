package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ CommentModel = (*customCommentModel)(nil)

type (
	// CommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentModel.
	CommentModel interface {
		commentModel
		customCommentLogicModel
	}

	customCommentModel struct {
		*defaultCommentModel
	}

	customCommentLogicModel interface {
		FindCountBySid(ctx context.Context) (int64, error)
		FindListBySId(ctx context.Context, solutionId int64) ([]*Comment, error)
	}
)

// NewCommentModel returns a model for the database table.
func NewCommentModel(conn *gorm.DB, c cache.CacheConf) CommentModel {
	return &customCommentModel{
		defaultCommentModel: newCommentModel(conn, c),
	}
}

func (m *defaultCommentModel) customCacheKeys(data *Comment) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

func (m *customCommentModel) FindCountBySid(ctx context.Context) (int64, error) {

	var resp int64
	err := m.QueryNoCacheCtx(ctx, &resp, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Comment{}).Count(&resp).Error
	})
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *customCommentModel) FindListBySId(ctx context.Context, solutionId int64) ([]*Comment, error) {

	var resp []*Comment
	err := m.QueryNoCacheCtx(ctx, &resp, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Comment{}).Where("sid = ?", solutionId).Find(&resp).Error
	})
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
