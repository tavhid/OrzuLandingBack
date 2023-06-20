package merchant

func (p *provider) GetList(search, category, city string, page, pageLimit uint) (merchant []map[string]interface{}, maxPage int64, err error) {

	return p.merchant.GetList(search, city, category, page, pageLimit)

}
