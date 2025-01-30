package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Reservas

type Booking struct {
	ID           	uuid.UUID `gorm:"type:char(36);primary_key;"`
	UserID 		  	uuid.UUID `gorm:"type:char(36)"`
	FlightID      	uuid.UUID `gorm:"type:char(36)"`
	SeatID 		  	uuid.UUID `gorm:"type:char(36)"`
	Status 			string // Pending, Confirmed, Cancelled
	TotalPrice 		float64
	gorm.Model
}