package model

import "gorm.io/gorm"

type Sale struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	status string `json:"status" form:"status"`
	qty    string `json:"total_qty" form:"total_qty"`
	pay    string `json:"total_pay" form:"total_pay"`
}
