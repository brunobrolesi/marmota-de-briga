package factory

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate
var once = &sync.Once{}

func ValidatorFactory() *validator.Validate {
	once.Do(func() {
		validate = validator.New(validator.WithRequiredStructEnabled())
	})

	return validate
}
