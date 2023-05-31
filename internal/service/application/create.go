package application

import (
	"cl/internal/structs"
)

func (p *provider) Create(fullName, phone string, regionID, applicationTypeID uint) (err error) {
	var applicationID int

	err = p.postgres.Model(structs.Application{}).
		Where("phone = ?", phone).
		Select("id").
		Scan(&applicationID).
		Error

	if applicationID != 0 || err != nil {
		return
	}
	result := p.postgres.
		Model(&structs.Application{}).
		Create(&structs.Application{
			FullName:          fullName,
			Phone:             phone,
			RegionID:          regionID,
			ApplicationTypeID: applicationTypeID,
		}).
		Error

	if result != nil {
		err = result
		return
	}

	return
}
