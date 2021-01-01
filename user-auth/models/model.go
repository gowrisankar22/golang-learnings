package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `json:"email" gorm:"not null;unique"`
	Password     string `json:"password" gorm:"-"`
	PasswordHash string `json:"password_hash"`
	Remember     string `json:"remember" gorm:"-"`
	RememberHash string	`json:"remember_hash"`
}
