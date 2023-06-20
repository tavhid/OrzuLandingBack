package application

import (
	"cl/internal/service/application"
	"cl/internal/storage/helper/otp"
	"cl/internal/structs"
	cotp "cl/pkg/cache/otp"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Provide(MerchantInfo)

// Application ...
type Application interface {
	Create(application structs.OtpInput) (err error)
}

// Dependencies ...
type Dependencies struct {
	fx.In
	Logger *logrus.Logger
	Otp otp.Otp
	Redis cotp.COtp
	Application application.Application
}

type provider struct {
	logger *logrus.Logger
	otp otp.Otp
	redis cotp.COtp
	application application.Application
}

// MerchantInfo ...
func MerchantInfo(params Dependencies) Application {
	return &provider{
		logger: params.Logger,
		otp: params.Otp,
		redis: params.Redis,
		application: params.Application,
	}
}
