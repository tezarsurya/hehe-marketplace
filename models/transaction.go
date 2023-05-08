package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID        uint
	AddressID     uint
	TotalPrice    int
	InvoiceCode   string
	PaymentMethod string
}
