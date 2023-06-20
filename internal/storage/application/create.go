package application

import (
	"cl/internal/structs"
)

func (p *provider) Create(otp_data structs.OtpInput) (err error) {

	err = p.otp.CheckOTP(otp_data)

	if err != nil {
		return
	}

	payload, err := p.redis.Get(otp_data.PhoneNumber)

	data := payload["data"].(map[string]interface{})

	if data["type_of_client"] == "merchant" {
		err = p.application.CreateMerchantApplication(data)
	} else {
		err = p.application.CreateUserApplication(data)
	}

	p.redis.Delete(otp_data.PhoneNumber)

	return

}
