package structs

import (
	"cl/pkg/utils"

	"github.com/go-playground/validator/v10"
)

// OtpSession ...
type OtpSession struct {
	CreatedAt int64  `json:"created_at"` // Время создания
	Attempts  int    `json:"attempts"`   // Количество попыток
	Code      string `json:"code"`       // Код
	Phone     string `json:"phone"`      // Телефон пользователя
}

// OtpRegistration ...
type OtpRegistration struct {
	OtpSession
	FullName        string `json:"full_name"`
	City            string `json:"city" validate:"required,min=1"`
	ApplicationType string `json:"application_type" validate:"required,min=1"`
}

// OtpInput ..
type OtpInput struct {
	PhoneNumber string `json:"phone" validate:"required"`
	OtpCode     string `json:"code"  validate:"required,len=4"`
}

// Validate — validates the data specified in the field tags
func (a *OtpInput) Validate() (err error) {
	validate = validator.New()

	errs := validate.Struct(a)
	if errs != nil {
		return utils.PlaygroundValidator(errs)
	}

	return
}
