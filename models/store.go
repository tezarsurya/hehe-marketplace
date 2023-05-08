package models

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	UserID      uint
	StoreName   string `validate:"required" json:"storeName"`
	StoreURL    string `validate:"required" json:"storeUrl"`
	Products    []Product
	ProductLogs []ProductLog
}
