package merchant

import (
	"cl/internal/structs"
)

func (p *provider) Get(id uint) (merchant structs.Merchant, commissions []structs.MonthCommissionOfMerchant, affiliates []structs.AffiliateOfMerchant, err error) {

	return p.merchant.Get(id)

}
