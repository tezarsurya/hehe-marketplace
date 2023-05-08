package models

import "gorm.io/gorm"

type ProductLog struct {
	gorm.Model
	ProductID           uint
	StoreID             uint
	CategoryID          uint
	TransactionDetailID uint
	ProductName         string
	Slug                string
	ResellerPrice       string
	CustomerPrice       string
	Description         string `gorm:"type:text"`
}
