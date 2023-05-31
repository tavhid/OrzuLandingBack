package structs

import "gorm.io/gorm"

type MainData struct {
	gorm.Model
	Limit   int
}
