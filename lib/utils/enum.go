package utils

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Enum interface {
	IsValid() bool
}

func ValidateEnum(fl validator.FieldLevel) bool {
	value := fl.Field().Interface().(Enum)
	return value.IsValid()
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("enum", ValidateEnum)
	}
}
