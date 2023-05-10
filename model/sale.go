package model

import "gorm.io/gorm"

type Sale struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	Status string `json:"status" form:"status"`
	Qty    string `json:"total_qty" form:"total_qty"`
	price  string `json:"total_pay" form:"total_price"`
}
