package merchant

import (
	"cl/internal/structs"
	"cl/pkg/bootstrap/http/misc/response"
	"fmt"
)

func (p *provider) GetList(search, city string, category uint) (merchant []structs.Merchant, err error) {
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
		Scan(&merchant).
		Error

	if err != nil {
		p.logger.Error(err)
		return
	}
	if len(merchant) == 0 {
		err = response.ErrDataNotFound
	}

	return
}
