package svc

import (
	"github.com/go-playground/validator/v10"
	"github.com/solitudealma/SOJ/backend/app/user/api/internal/config"
	"github.com/solitudealma/SOJ/backend/app/user/api/internal/middleware"
	"github.com/solitudealma/SOJ/backend/common/validate"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config
	Auth   rest.Middleware

	Validator *validator.Validate
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Auth:   middleware.NewAuthMiddleware().Handle,

		Validator: validate.Validate(),
	}
}
