package models

import "gorm.io/gorm"

type TransactionDetail struct {
	gorm.Model
	StoreID     uint
	Quantity    int
	TotalPrice  int
	ProductLogs []ProductLog
}
