package svc

import (
	"github.com/go-playground/validator/v10"
	"github.com/solitudealma/SOJ/backend/app/user/cmd/api/internal/config"
	"github.com/solitudealma/SOJ/backend/app/user/cmd/api/internal/middleware"
	"github.com/solitudealma/SOJ/backend/app/user/model"
	"github.com/solitudealma/SOJ/backend/common/stores/gormx"
	"github.com/solitudealma/SOJ/backend/common/validate"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"

	"golang.org/x/sync/singleflight"
)

type ServiceContext struct {
	Config config.Config
	Auth   rest.Middleware

	Validator   *validator.Validate
	RedisClient *redis.Redis

	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	db, err := gormx.BuildDBManager(c.MySQL)
	if err != nil {
		logx.Errorv("connect to db failed")
		return &ServiceContext{}
	}

	concurrencyControl := &singleflight.Group{}
	redisClient := redis.New(c.Redis.Host, func(r *redis.Redis) {
		r.Type = c.Redis.Type
		r.Pass = c.Redis.Pass
	})

	return &ServiceContext{
		Config: c,
		Auth:   middleware.NewAuthMiddleware(c, concurrencyControl, redisClient).Handle,

		Validator:   validate.Validate(),
		RedisClient: redisClient,

		UserModel: model.NewUserModel(db, c.Cache),
	}
}
