package application

import (
	"cl/internal/structs"
	"cl/pkg/bootstrap/http/misc/response"
)

func (p *provider) CreateMerchantApplication(application map[string]interface{}) (err error) {
	var applicationID int

	err = p.postgres.Model(structs.MerchantApplication{}).
		Where("phone = ?", application["phone"]).
		Select("id").
		Scan(&applicationID).
		Error
	if applicationID != 0 {
		err = response.ErrAlreadyRegistered
		return
	} else if err != nil {
		return
	}
	delete(application, "type_of_client")

	result := p.postgres.
		Model(&structs.MerchantApplication{}).
		Create(&application).
		Error

	if result != nil {
		err = result
		return
	}

	return
}
