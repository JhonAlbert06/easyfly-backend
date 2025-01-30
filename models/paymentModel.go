package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Pagos

type Payment struct {
	ID           	uuid.UUID `gorm:"type:char(36);primary_key;"`
	BookingID 		uuid.UUID `gorm:"type:char(36)"`
	UserID 		  	uuid.UUID `gorm:"type:char(36)"`
	Amount      	float64
	PaymentMethod 	string // Visa, MasterCard, PayPal
	Status 			string // Pending, Confirmed, Cancelled, Refunded, Failed
	gorm.Model
}