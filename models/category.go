package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryName string       `validate:"required" json:"categoryName"`
	Products     []Product    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	ProductLogs  []ProductLog `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
