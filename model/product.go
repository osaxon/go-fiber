package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title 	 string `gorm:"not null" json:"title"`
	Slug 	 string `gorm:"not null" json:"slug"`
	Description  string `gorm:"not null" json:"description"`
	Price 	 int `gorm:"not null" json:"price"`
}