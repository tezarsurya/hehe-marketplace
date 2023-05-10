package models

import "gorm.io/gorm"

type TransactionDetail struct {
	gorm.Model
	StoreID     uint
	Quantity    int
	TotalPrice  int
	ProductLogs []ProductLog `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
