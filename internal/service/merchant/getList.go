package merchant

import (
	"cl/internal/structs"
	"fmt"
)

func (p *provider) GetList(search, city string, category, page, pageLimit uint) (merchant []structs.Merchant, err error) {
	search = "%" + search + "%"
	var where string

	if search != "" {
		where += fmt.Sprintf("name LIKE '%s'", search)
	}
	if category != 0 {
		if where != "" {
			where += " AND "
		}
		where += fmt.Sprintf("category= %d", category)
	}
	if city != "" {
		if where != "" {
			where += " AND "
		}
		where += fmt.Sprintf("city = '%s'", city)
	}

	err = p.postgres.
		Model(&structs.Merchant{}).
		Select("*").
		Where(where).
		Limit(int(pageLimit)).
		Offset((int(page) - 1) * int(pageLimit)).
		Scan(&merchant).
		Error

	if err != nil {
		p.logger.Error(err)
		return
	}

	return
}
