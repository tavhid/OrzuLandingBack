package structs

import (
	"cl/pkg/utils"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

type Application struct {
	gorm.Model
	FullName          string `json:"full_name" validate:"required"`
	Phone             string `json:"phone" validate:"required"`
	RegionID          uint   `json:"region_id" validate:"required,min=1"`
	ApplicationTypeID uint   `json:"application_type_id" validate:"required,min=1"`
}

// Validate — validates the data specified in the field tags
func (a *Application) Validate() (err error) {
	validate = validator.New()

	errs := validate.Struct(a)
	if errs != nil {
		return utils.PlaygroundValidator(errs)
	}

	return
}