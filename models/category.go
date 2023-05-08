package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryName string `validate:"required" json:"categoryName"`
	Products     []Product
	ProductLogs  []ProductLog
}
