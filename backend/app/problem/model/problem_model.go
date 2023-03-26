package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ ProblemModel = (*customProblemModel)(nil)

type (
	// ProblemModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProblemModel.
	ProblemModel interface {
		problemModel
		customProblemLogicModel
	}

	customProblemModel struct {
		*defaultProblemModel
	}

	customProblemLogicModel interface {
		FindCount(ctx context.Context) (int64, error)
		FindPageListByPage(ctx context.Context, page, pageSize int, orderBy string) ([]*Problem, error)
	}
)

// NewProblemModel returns a model for the database table.
func NewProblemModel(conn *gorm.DB, c cache.CacheConf) ProblemModel {
	return &customProblemModel{
		defaultProblemModel: newProblemModel(conn, c),
	}
}

func (m *defaultProblemModel) customCacheKeys(data *Problem) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

func (m *customProblemModel) FindCount(ctx context.Context) (int64, error) {

	var resp int64
	err := m.QueryNoCacheCtx(ctx, &resp, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Problem{}).Count(&resp).Error
	})
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *customProblemModel) FindPageListByPage(ctx context.Context, page, pageSize int, orderBy string) ([]*Problem, error) {
	if orderBy == "" {
		orderBy = "id ASC"
	}

	if page < 1 {
		page = 1
	}

	offset := (page - 1) * pageSize

	var resp []*Problem
	err := m.QueryNoCacheCtx(ctx, &resp, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Problem{}).Select("problem_id", "title", "difficulty").Offset(offset).Limit(pageSize).Order(orderBy).Find(&resp).Error
	})
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
