package models

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	UserID      uint
	StoreName   string       `validate:"required" json:"storeName"`
	StoreURL    string       `validate:"required" json:"storeUrl"`
	Products    []Product    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	ProductLogs []ProductLog `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
