package merchant

import (
	"cl/internal/service/merchant"
	"cl/internal/structs"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Provide(MerchantInfo)

// Merchant ...
type Merchant interface {
	GetList(search, city string, category uint) (merchant []structs.Merchant, err error)
	Get(id uint) (merchant structs.Merchant, err error)
}

// Dependencies ...
type Dependencies struct {
	fx.In
	Logger *logrus.Logger

	Merchant merchant.Merchant
}

type provider struct {
	logger   *logrus.Logger
	merchant merchant.Merchant
}

// MerchantInfo ...
func MerchantInfo(params Dependencies) Merchant {
	return &provider{
		logger:   params.Logger,
		merchant: params.Merchant,
	}
}
