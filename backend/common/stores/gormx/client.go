package gormx

import (
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	DSN      string
	LogLevel logger.LogLevel
}

func BuildDBManager(conf Config) (*gorm.DB, error) {
	mysqlConfig := mysql.Config{
		DSN:                       conf.DSN,
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // 慢 SQL 阈值, 200ms
			LogLevel:                  logger.Info,            // 日志级别
			IgnoreRecordNotFoundError: true,                   // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,                   // 彩色打印
		},
	)
	client, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,      // 禁用自动创建外键约束
		Logger:                                   newLogger, // 使用自定义 Logger
	})

	if err != nil {
		return nil, err
	}

	// TODO
	// db.SetMaxOpenConns(conf.MaxOpenConns) // 打开数据库连接的最大数量
	// db.SetMaxIdleConns(conf.MaxIdleConns) // 空闲连接池中连接的最大数量
	// db.SetConnMaxLifetime(conf.ConnMaxLifeTime)

	return client, nil
}
