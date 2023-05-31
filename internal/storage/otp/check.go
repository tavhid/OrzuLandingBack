package otp

import (
	"cl/internal/structs"
	"cl/pkg/bootstrap/http/misc/response"
	"cl/pkg/utils"
)

func (p *provider) CheckOTP(data structs.OtpInput) (err error) {

	phone, err := utils.CheckPhone(data.PhoneNumber)
	if err != nil {
		return
	}
	if err = data.Validate(); err != nil {
		return
	}

	payload, err := p.redis.Get(phone)
	if err != nil {
		return
	}

	if data.OtpCode != payload["code"] {
		return response.ErrIncorrectOtpCode
	}

	if (payload["attempts"]).(float64) > 3 {
		return response.ErrLimitExceeded
	}

	err = p.application.Create(
		payload["full_name"].(string),
		payload["phone"].(string),
		uint(payload["region_id"].(float64)),
		uint(payload["application_type_id"].(float64)),
	)

	p.redis.Delete(phone)

	return
}
