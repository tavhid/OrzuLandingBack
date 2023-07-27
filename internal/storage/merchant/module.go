package merchant

import (
	"cl/internal/service/cft"
	"cl/internal/service/merchant"
	"cl/internal/structs"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Provide(MerchantInfo)

// Merchant ...
type Merchant interface {
	GetList(search, category, city string, page, pageLimit uint) (merchantList []map[string]interface{}, maxPage int64, cities []string, categories, err error)
	Get(id uint) (merchant structs.Merchant, commissions []structs.MonthCommissionOfMerchant, affiliates []structs.AffiliateOfMerchant, err error)
}

// Dependencies ...
type Dependencies struct {
	fx.In
	Logger *logrus.Logger

	Merchant merchant.Merchant
	Cft      cft.CFT
}

type provider struct {
	logger   *logrus.Logger
	merchant merchant.Merchant
	cft      cft.CFT
}

// MerchantInfo ...
func MerchantInfo(params Dependencies) Merchant {
	return &provider{
		logger:   params.Logger,
		merchant: params.Merchant,
		cft:      params.Cft,
	}
}
