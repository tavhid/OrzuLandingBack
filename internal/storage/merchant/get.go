package merchant

import (
	"cl/internal/structs"
)

func (p *provider) Get(id uint) (merchant structs.Merchant, err error) {

	return p.merchant.Get(id)

}
