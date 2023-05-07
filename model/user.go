package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Dob      string `json:"dob" form:"dob"`
	Gender   string `json:"gender" form:"gender"`
	Position string `json:"position" form:"position"`
	Token    string `gorm:"-"`
	Sales    []Sale
}
