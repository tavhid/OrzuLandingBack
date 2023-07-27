package merchant

func (p *provider) GetList(search, category, city string, page, pageLimit uint) (merchantList []map[string]interface{}, maxPage int64, cities []string, categories, err error) {

	filteredMerchantList, maxPage, err := p.merchant.GetList(search, category, city, page, pageLimit)
	for _, merchant := range filteredMerchantList{
		merchantList = append(merchantList, map[string]interface{}{
			"id": merchant.MerchantID,
			"logo": merchant.Logo,
			"name": merchant.Name,
		})
	}
	return
}
