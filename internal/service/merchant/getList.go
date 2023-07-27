package merchant

import (
	"cl/internal/structs"
	"log"
	"math"
	"strings"
)

func (p *provider) GetList(search, city, category string, page, pageLimit uint) (merchantList []structs.Merchant, maxPage int64, err error) {
	merchantList, err = p.cft.GetMerchantList()
	if err != nil {
		log.Println(err)
		return
	}

	//filtering CFT data
	search = strings.ToLower(search)
	for i := 0; i < len(merchantList); i++ {

		//if the list of merchant cities contains a city that the user needs, then we change haveCity
		haveCity := false
		for _, val := range merchantList[i].City {
			if val == city || city == "" {
				haveCity = true
				break
			}
		}
		//filtering by search, city and category
		if !strings.Contains(strings.ToLower(merchantList[i].Name), search) ||
			!haveCity ||
			(category != "" && merchantList[i].Category != category) {
			//removing item with index i from merchant list
			merchantList = append(merchantList[:i], merchantList[i+1:]...)
			i--
		}
	}

	//We count the number of possible pages in pagination depending on the selected limit (default 16 if not specified)
	maxPage = int64((len(merchantList) + int(pageLimit) - 1) / int(pageLimit))

	paginationStart := int(math.Max(0, float64((page-1)*pageLimit)))
	pagintaionEnd := int(math.Min(float64(len(merchantList)-1), float64(page*pageLimit)))

	if paginationStart >= len(merchantList) || pagintaionEnd < 0 {
		merchantList = []structs.Merchant{}
		return
	}
	merchantList = merchantList[paginationStart:pagintaionEnd]

	// log.Println(merchantList, maxPage)
	return
}
