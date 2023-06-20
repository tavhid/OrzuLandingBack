package merchant

import (
	"cl/internal/structs"
	"cl/pkg/bootstrap/http/misc/response"
	"fmt"
	"strings"
)

func (p *provider) GetList(search, city, category string, page, pageLimit uint) (merchant []map[string]interface{}, maxPage int64, err error) {
	search = strings.ToLower(search)
	search = "%" + search + "%"
	var where string

	if search != "" {
		where += fmt.Sprintf("LOWER(name) LIKE '%s'", search)
	}
	if category != "" {
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
		Select("id", "name", "logo").
		Where(where).
		Count(&maxPage).
		Limit(int(pageLimit)).
		Offset((int(page) - 1) * int(pageLimit)).
		Scan(&merchant).
		Error

	if err != nil {
		p.logger.Error(err)
		return
	}

	if merchant == nil {
		err = response.ErrDataNotFound
	}

	return
}
