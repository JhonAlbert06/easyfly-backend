package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Aeropuertos

type Airport struct {
	ID           	uuid.UUID `gorm:"type:char(36);primary_key;"`
	Name 		  	string `gorm:"unique"`
	Code 		  	string `gorm:"unique"`
	City 		  	string
	Country 	  	string
	gorm.Model
}