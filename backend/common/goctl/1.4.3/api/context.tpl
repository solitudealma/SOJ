package svc

import (
	"github.com/solitudealma/SOJ/backend/common/validate"
	"github.com/go-playground/validator/v10"
    {{.configImport}}
)

type ServiceContext struct {
	Config {{.config}}
	{{.middleware}}
    Validator   *validator.Validate
}

func NewServiceContext(c {{.config}}) *ServiceContext {
	return &ServiceContext{
		Config: c, 
		{{.middlewareAssignment}}
        Validator:   validate.Validate(),
	}
}
