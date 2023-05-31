package otp

import (
	"cl/internal/service/application"
	"cl/internal/structs"
	"cl/pkg/cache/otp"
	"cl/pkg/config"
	"cl/pkg/gateways/sms"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

// TODO: REVISION OTP MODULES

// Module ...
var Module = fx.Provide(NewOtpModule)

// IRegistration ...
type Otp interface {
	SendOTP(application structs.Application) (err error)
	CheckOTP(data structs.OtpInput) (err error)
}

// Dependencies ...
type Dependencies struct {
	fx.In
	Logger      *logrus.Logger
	Config      *config.Config
	Redis       otp.COtp
	GSms        sms.GSms
	Application application.Application
}

type provider struct {
	logger      *logrus.Logger
	config      *config.Config
	redis       otp.COtp
	gSms        sms.GSms
	application application.Application
}

// NewIRegistrationModule ...
func NewOtpModule(params Dependencies) Otp {
	return &provider{
		logger:      params.Logger,
		config:      params.Config,
		redis:       params.Redis,
		gSms:        params.GSms,
		application: params.Application,
	}
}
