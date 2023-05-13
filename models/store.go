package models

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	UserID      uint
	StoreName   string       `json:"store_name,omitempty"`
	StoreURL    string       `json:"store_url,omitempty"`
	Products    []Product    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProductLogs []ProductLog `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
