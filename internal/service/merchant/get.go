package merchant

import (
	"cl/internal/structs"
	"cl/pkg/bootstrap/http/misc/response"
)

func (p *provider) Get(id uint) (merchant structs.Merchant, err error) {

	err = p.postgres.
		Model(&structs.Merchant{}).
		Select("*").
		Where("id = ?", id).
		Scan(&merchant).
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
