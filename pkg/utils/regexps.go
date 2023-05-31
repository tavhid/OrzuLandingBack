package utils

import (
	"cl/pkg/bootstrap/http/misc/response"
	"regexp"
)

// CheckPhone ...
func CheckPhone(phone string) (result_phone string, err error) {
	// tjk phone checker
	reg, err := regexp.Compile(`^(992)?\d{9}$`)

	if !reg.MatchString(phone) {
		err = response.ErrInvalidPhoneFormat
		return "", err
	}

	return phone, nil
}
