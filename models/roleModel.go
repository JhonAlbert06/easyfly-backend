package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Roles

type Role struct {
	ID           	uuid.UUID `gorm:"type:char(36);primary_key;"`
	Name 		  	string
	Description     string
	gorm.Model
}