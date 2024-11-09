package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Login    string `json:"login" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"` // Exclude from JSON responses
}
