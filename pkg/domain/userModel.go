package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" validate:"required,min=8,max=24"`
	Password string `json:"password" validate:"required,min=8,max=16"`
	Email    string `json:"email" validate:"email,required"`
}

type Notes struct {
	Sid  string `json:"sid"`
	Note string `json:"note"`
}
