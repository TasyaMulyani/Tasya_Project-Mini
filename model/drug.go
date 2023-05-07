package model

import "gorm.io/gorm"

type Drug struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	Name   string `json:"name" form:"name"`
	Price  string `json:"price" form:"price"`
	Stock  string `json:"stock" form:"stock"`
}
