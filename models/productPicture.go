package models

import "gorm.io/gorm"

type ProductPicture struct {
	gorm.Model
	ProductID uint
	URL       string
}
