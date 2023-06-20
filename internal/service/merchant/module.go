package merchant

import (
	"cl/internal/structs"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Module ...
var Module = fx.Provide(MerchantInfo)

// Merchant ...
type Merchant interface {
	GetList(search, category, city string, page, pageLimit uint) (merchant []map[string]interface{}, maxPage int64, err error)
	Get(id uint) (merchant structs.Merchant, commissions []structs.MonthCommissionOfMerchant, affiliates []structs.AffiliateOfMerchant, err error)
}

// Dependencies ...
type Dependencies struct {
	fx.In
	Postgres *gorm.DB
	Logger   *logrus.Logger
}

type provider struct {
	postgres *gorm.DB
	logger   *logrus.Logger
}

// MerchantInfo ...
func MerchantInfo(params Dependencies) Merchant {
	return &provider{
		postgres: params.Postgres,
		logger:   params.Logger,
	}
}
