package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	StoreID         uint
	CategoryID      uint
	ProductName     string
	Slug            string
	ResellerPrice   string
	CustomerPrice   string
	Available       int
	Description     string           `gorm:"type:text"`
	ProductPictures []ProductPicture `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	ProductLogs     []ProductLog     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
