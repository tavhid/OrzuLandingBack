package structs

import (
	"cl/pkg/utils"

	"github.com/go-playground/validator/v10"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

type UserApplication struct {
	Model
	FullName        string `json:"full_name" validate:"required"`
	Phone           string `json:"phone" validate:"required,len=9,number"`
	City            string `json:"city" validate:"required,min=1"`
	ApplicationType string `json:"application_type" validate:"required,min=1"`
}

// Validate — validates the data specified in the field tags
func (a *UserApplication) Validate() (err error) {
	validate = validator.New()

	errs := validate.Struct(a)
	if errs != nil {
		return utils.PlaygroundValidator(errs)
	}

	return
}
