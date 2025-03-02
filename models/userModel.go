package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Usuarios

type User struct {
	ID           	uuid.UUID `gorm:"type:char(36);primary_key;"`
	RoleId 		  	uuid.UUID `gorm:"type:char(36)"`
	FullName      	string
	Email           string    `gorm:"unique"`
	Password        string
	PasswordVersion uuid.UUID `gorm:"type:char(36)"`
	IsEnabled       bool
	gorm.Model
}