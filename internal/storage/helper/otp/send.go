package otp

import (
	"cl/internal/structs"
	"cl/pkg/utils"
)

func (p *provider) SendOTP(phone, message string, data interface{}) (err error) {
	// Formation of the text of the SMS message
	// otpCode := utils.GenOtp()
	otpCode := "1234"
	// _messageText := fmt.Sprintf(p.config.SMS.RegistrationText, otpCode)

	// Try send sms
	// err = p.gSms.Send(phone, _messageText)
	// if err != nil {
	// 	return
	// }

	tempRedisPayload, _ := p.redis.Get(phone)
	tempRedisPayload, err = utils.ParseOTP(tempRedisPayload)

	if err != nil {
		return
	}

	redisPayload := structs.OtpSession{
		CreatedAt: tempRedisPayload["created_at"].(int64),
		Attempts:  tempRedisPayload["attempts"].(int) + 1,
		// Attempts: tempRedisPayload["attempts"].(in~t),
		Code:    otpCode,
		Phone:   phone,
		Data:    data,
		Message: message,
	}

	if err = p.redis.Save(phone, redisPayload); err != nil {
		return
	}

	return
}
