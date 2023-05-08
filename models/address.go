package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID         uint
	AddressTitle   string        `validate:"required" json:"addressTitle"`
	ReceiverName   string        `validate:"required" json:"receiverName"`
	Phone          string        `validate:"required" json:"phone"`
	AddressDetails string        `json:"addressDetails"`
	Transactions   []Transaction `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
