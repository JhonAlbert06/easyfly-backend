package models

import "github.com/google/uuid"

type Ticket struct {
	ID           	uuid.UUID `gorm:"type:char(36);primary_key;"`
	BookingID 		uuid.UUID `gorm:"type:char(36)"`
	QRCode      	string
	Passenger   	string
	FlightNumber 	string
	SeatNumber 		string
	DepartureTime 	string
	ArrivalTime 	string
	Status 			string // Issued, Used, Cancelled
}