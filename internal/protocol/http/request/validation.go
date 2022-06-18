package request

import (
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	Validator *validator.Validate
}

func (cv *Validator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
