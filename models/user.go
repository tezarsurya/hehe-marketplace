package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string        `gorm:"type:varchar(255)" validate:"required" json:"name"`
	Password     string        `gorm:"type:varchar(255)" validate:"required,min=8" json:"password,omitempty"`
	Phone        string        `gorm:"unique;type:varchar(20)" validate:"required" json:"phone"`
	Birthday     time.Time     `gorm:"type:date" validate:"required" json:"birthday"`
	Gender       string        `gorm:"type:enum('male','female')" validate:"required" json:"gender"`
	About        string        `gorm:"type:text" json:"about,omitempty"`
	Profession   string        `gorm:"type:varchar(255)" validate:"required" json:"profession"`
	Email        string        `gorm:"unique;type:varchar(255)" validate:"required,email" json:"email"`
	ProvinceId   string        `gorm:"type:varchar(255)" validate:"required" json:"provinceId"`
	CityId       string        `gorm:"type:varchar(255)" validate:"required" json:"cityId"`
	IsAdmin      *bool         `validate:"required" json:"isAdmin"`
	Store        Store         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Addresses    []Address     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Transactions []Transaction `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
