package merchant

import (
	"cl/internal/structs"
	"cl/pkg/bootstrap/http/misc/response"
)

func (p *provider) Get(id uint) (merchant structs.Merchant, commissions []structs.MonthCommissionOfMerchant, affiliates []structs.AffiliateOfMerchant, err error) {

	err = p.postgres.
		Model(&structs.Merchant{}).
		Select("id", "name", "logo").
		Where("id = ?", id).
		Scan(&merchant).
		Error

	if err != nil {
		p.logger.Error(err)
		return
	}

	err = p.postgres.
		Model(&structs.MonthCommissionOfMerchant{}).
		Select("date", "commission").
		Where("merchant_id = ?", id).
		Scan(&commissions).
		Error

	if err != nil {
		p.logger.Error(err)
		return
	}

	err = p.postgres.
		Model(&structs.AffiliateOfMerchant{}).
		Select("address", "contact").
		Where("merchant_id = ?", id).
		Scan(&affiliates).
		Error

	if err != nil {
		p.logger.Error(err)
		return
	}

	if merchant.ID == 0 {
		err = response.ErrDataNotFound
	}

	return
}
