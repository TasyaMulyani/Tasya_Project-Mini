package model

import "gorm.io/gorm"

type drug struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	Name   string `json:"name" form:"name"`
	price  string `json:"price" form:"price"`
	stock  string `json:"stock" form:"stock"`
}
