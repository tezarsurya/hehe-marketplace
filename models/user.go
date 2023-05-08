package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string    `validate:"required" json:"name"`
	Password     string    `validate:"required,min=8" json:"password,omitempty"`
	Phone        string    `gorm:"unique" validate:"required" json:"phone"`
	Birthday     time.Time `gorm:"type:date" validate:"required" json:"birthday"`
	Gender       string    `validate:"required" json:"gender"`
	About        string    `gorm:"type:text" json:"about,omitempty"`
	Profession   string    `validate:"required" json:"profession"`
	Email        string    `gorm:"unique" validate:"required,email" json:"email"`
	ProvinceId   string    `validate:"required" json:"provinceId"`
	CityId       string    `validate:"required" json:"cityId"`
	IsAdmin      *bool     `validate:"required" json:"isAdmin"`
	Store        Store     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Addresses    []Address `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Transactions []Transaction
}
