package otp

import (
	"cl/pkg/gateways/sms"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"

	"go.uber.org/fx"
)

// Module ...
var Module = fx.Provide(NewCOtpModule)

// COtp ...
type COtp interface {
	Get(key string) (expandTo map[string]interface{}, err error)
	Save(key string, value interface{}) (err error)
	Update(key string, value interface{}) (err error)
	Delete(key string)
}

// Dependencies ...
type Dependencies struct {
	fx.In
	Logger *logrus.Logger
	RDB    *redis.Client
	GSms   sms.GSms
}

type provider struct {
	logger *logrus.Logger
	rdb    *redis.Client
	gSms   sms.GSms
}

// NewCOtpModule ...
func NewCOtpModule(params Dependencies) COtp {
	return &provider{
		logger: params.Logger,
		rdb:    params.RDB,
		gSms:   params.GSms,
	}
}
