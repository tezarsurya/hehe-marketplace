package models

type Category struct {
	Model
	CategoryName string       `gorm:"type:varchar(255)" validate:"required" json:"category_name"`
	Products     []Product    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products,omitempty"`
	ProductLogs  []ProductLog `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"product_logs,omitempty"`
}
