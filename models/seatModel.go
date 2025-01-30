package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Asientos

type Seat struct {
	ID           	uuid.UUID `gorm:"type:char(36);primary_key;"`
	FlightID      	uuid.UUID `gorm:"type:char(36)"`
	SeatNumber 	  	string
	IsBooked 	  	bool
	gorm.Model
}