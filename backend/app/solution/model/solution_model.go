package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ SolutionModel = (*customSolutionModel)(nil)

type (
	// SolutionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSolutionModel.
	SolutionModel interface {
		solutionModel
		customSolutionLogicModel
	}

	customSolutionModel struct {
		*defaultSolutionModel
	}

	customSolutionLogicModel interface {
		FindCount(ctx context.Context) (int64, error)
		FindPageListByPage(ctx context.Context, page, pageSize int, orderBy string) ([]*Solution, error)
	}
)

// NewSolutionModel returns a model for the database table.
func NewSolutionModel(conn *gorm.DB, c cache.CacheConf) SolutionModel {
	return &customSolutionModel{
		defaultSolutionModel: newSolutionModel(conn, c),
	}
}

func (m *defaultSolutionModel) customCacheKeys(data *Solution) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

func (m *customSolutionModel) FindCount(ctx context.Context) (int64, error) {

	var resp int64
	err := m.QueryNoCacheCtx(ctx, &resp, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Solution{}).Count(&resp).Error
	})
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *customSolutionModel) FindPageListByPage(ctx context.Context, page, pageSize int, orderBy string) ([]*Solution, error) {
	if orderBy == "" {
		orderBy = "id ASC"
	}

	if page < 1 {
		page = 1
	}

	offset := (page - 1) * pageSize

	var resp []*Solution
	err := m.QueryNoCacheCtx(ctx, &resp, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Solution{}).
			Select([]string{"solution_id", "problem_id", "title", "problem_source", "author_name", "author_avatar", "author_id", "create_time", "read", "problem_difficulty"}).
			Offset(offset).Limit(pageSize).
			Order(orderBy).
			Find(&resp).Error
	})
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
