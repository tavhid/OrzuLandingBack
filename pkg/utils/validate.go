package utils

import (
	"cl/pkg/bootstrap/http/misc/response"

	"github.com/go-playground/validator/v10"
)

// PlaygroundValidator ...
func PlaygroundValidator(errs error) (err error) {
	for _, errs := range errs.(validator.ValidationErrors) {
		// fmt.Println("Field: ", errs.Field())
		// fmt.Println("Tag: ", errs.Tag())
		// fmt.Println("ActualTag: ", errs.ActualTag())
		// fmt.Println("Kind: ", errs.Kind())
		// fmt.Println("Type: ", errs.Type())
		// fmt.Println("Value: ", errs.Value())
		// fmt.Println("Param: ", errs.Param())
		// fmt.Println()

		if errs.Value() == "" {
			return response.ErrEmptyFileds
		}
		if errs.ActualTag() == "required" {
			return response.ErrEmptyFileds
		}
		if errs.ActualTag() == "max" || errs.ActualTag() == "min" || errs.ActualTag() == "len" {
			return response.ErrIncorrectFillFields
		}
	}

	return
}
