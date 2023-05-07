package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	dob      string `json:"dob" form:"dob"`
	gender   string `json:"gender" form:"gender"`
	position string `json:"position" form:"position"`
	Token    string `gorm:"-"`
	Sales    []Sale
}
