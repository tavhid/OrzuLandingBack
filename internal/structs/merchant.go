package structs

import (
	"time"

	"gorm.io/gorm"
)

type Merchant struct {
	Model
	Name     string `json:"name"`
	Category uint   `json:"category"`
	Logo     string `json:"logo"`
	Address  string `json:"address"`
	Contact  string `json:"contact"`
	City     string `json:"city"`
}

type Model struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
