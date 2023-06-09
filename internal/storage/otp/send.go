package otp

import (
	"cl/internal/structs"
	"cl/pkg/utils"
)

func (p *provider) SendOTP(application structs.Application) (err error) {

	application.Phone, err = utils.CheckPhone(application.Phone)

	if err != nil {
		return
	}
	// Validate fileds
	if err = application.Validate(); err != nil {
		return
	}

	// Formation of the text of the SMS message
	// otpCode := utils.GenOtp()
	otpCode := "1234"
	// _messageText := fmt.Sprintf(p.config.SMS.RegistrationText, otpCode)

	// Try send sms
	// err = p.gSms.Send(phone, _messageText)
	// if err != nil {
	// 	return
	// }

	tempRedisPayload, _ := p.redis.Get(application.Phone)
	tempRedisPayload, err = utils.ParseOTP(tempRedisPayload)

	if err != nil {
		return
	}

	redisPayload := structs.OtpRegistration{
		OtpSession: structs.OtpSession{
			CreatedAt: tempRedisPayload["created_at"].(int64),
			Attempts:  tempRedisPayload["attempts"].(int) + 1,
			// Attempts: tempRedisPayload["attempts"].(int),
			Code:  otpCode,
			Phone: application.Phone,
		},
		FullName:        application.FullName,
		City:            application.City,
		ApplicationType: application.ApplicationType,
	}

	if err = p.redis.Save(application.Phone, redisPayload); err != nil {
		return
	}

	return
}
