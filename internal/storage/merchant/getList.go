package merchant

import "cl/internal/structs"

func (p *provider) GetList(search, city string, category, page, pageLimit uint) (merchant []structs.Merchant, err error) {

	return p.merchant.GetList(search, city, category, page, pageLimit)

}
