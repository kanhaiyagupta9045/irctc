package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Email    string `json:"email" validate:"required,email" gorm:"unique;not null"`
	Username string `json:"username" validate:"required" gorm:"not null"`
	Password string `json:"password" validate:"required,min=8" gorm:"not null"`
	UserType string `json:"usertype" validate:"required" gorm:"not null"`
}

type LoginData struct {
	Email    string `json:"email" validate:"required,email" gorm:"unique;not null"`
	Password string `json:"password" validate:"required,min=8" gorm:"not null"`
}
