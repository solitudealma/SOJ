package gormx

import (
	"context"
	"errors"
	"time"

	"github.com/solitudealma/SOJ/backend/common/tool"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Logger struct {
	SlowThreshold time.Duration
}

var _ logger.Interface = (*Logger)(nil)

func NewGormLogger() *Logger {
	return &Logger{}
}

func (l *Logger) LogMode(lev logger.LogLevel) logger.Interface {
	return &Logger{
		SlowThreshold: l.SlowThreshold,
	}
}

func (l *Logger) Info(ctx context.Context, msg string, data ...interface{}) {
	logx.WithContext(ctx).Infof(msg, data)
}

func (l *Logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	logx.WithContext(ctx).Errorf(msg, data)
}

func (l *Logger) Error(ctx context.Context, msg string, data ...interface{}) {
	logx.WithContext(ctx).Errorf(msg, data)
}

func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	// 获取运行时间
	elapsed := time.Since(begin)
	// 获取 SQL 语句和返回条数
	sql, rows := fc()
	// 通用字段
	logFields := []logx.LogField{
		logx.Field("sql", sql),
		logx.Field("time", tool.MicrosecondsStr(elapsed)),
		logx.Field("rows", rows),
	}
	// Gorm 错误
	if err != nil {
		// 记录未找到的错误使用 warning 等级
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logx.WithContext(ctx).Infow("Database ErrRecordNotFound", logFields...)
		} else {
			// 其他错误使用 error 等级
			logFields = append(logFields, logx.Field("catch error", err))
			logx.WithContext(ctx).Errorw("Database Error", logFields...)
		}
	}
	// 慢查询日志
	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		logx.WithContext(ctx).Sloww("Database Slow Log", logFields...)
	}
	// 非生产模式下，记录所有 SQL 请求
	//if l.c.Mode != service.ProMode {
	logx.WithContext(ctx).Infow("Database Query", logFields...)
	//}
}
