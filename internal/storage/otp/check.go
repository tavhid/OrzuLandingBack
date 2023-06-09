package otp

import (
	"cl/internal/structs"
	"cl/pkg/bootstrap/http/misc/response"
	"cl/pkg/utils"
	"log"
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

	log.Println(payload)
	err = p.application.Create(
		payload["full_name"].(string),
		payload["phone"].(string),
		payload["city"].(string),
		payload["application_type"].(string),
	)

	p.redis.Delete(phone)

	return
}
