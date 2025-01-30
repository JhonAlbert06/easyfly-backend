package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Vuelos

type Flight struct {
	ID				uuid.UUID `gorm:"type:char(36);primary_key;"`
	AirLine 		string
	FlightNumber 	string
	Origin 			string
	Destination 	string
	DepartureTime 	time.Time
	ArrivalTime 	time.Time
	Price 			float64
	SeatsTotal 		int
	SeatsBooked 	int
	SeatsAvailable 	int
	Status 			string // Scheduled, Delayed, Cancelled, Completed
	gorm.Model
}