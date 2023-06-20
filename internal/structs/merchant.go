package structs

import (
	"cl/pkg/utils"

	"github.com/go-playground/validator/v10"
)

type Merchant struct {
	Model
	Name     string `json:"name"`
	Category string `json:"category,omitempty"`
	Logo     string `json:"logo"`
	City     string `json:"city,omitempty"`
}

type MerchantApplication struct {
	Model
	FullName    string `json:"full_name" validate:"required"`
	CompanyName string `json:"company_name" validate:"required"`
	Address     string `json:"address" validate:"required"`
	Phone       string `json:"phone" validate:"required,len=9,number"`
	Industry    string `json:"industry" validate:"required"`
}

type MonthCommissionOfMerchant struct {
	Model
	Date       float32 `json:"date"`
	Commission int     `json:"commission"`
	MerchantID uint    `json:"merchant_id,omitempty"`
}

type AffiliateOfMerchant struct {
	Model
	Address    string `json:"address"`
	Contact    string `json:"contact"`
	MerchantID uint   `json:"merchant_id,omitempty"`
}

// Validate — validates the data specified in the field tags
func (a *MerchantApplication) Validate() (err error) {
	validate = validator.New()

	errs := validate.Struct(a)
	if errs != nil {
		return utils.PlaygroundValidator(errs)
	}

	return
}
