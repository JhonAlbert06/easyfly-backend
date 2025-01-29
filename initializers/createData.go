package initializers

import (
	"github.com/JhonAlbert06/easyfly-backend/models"
	"github.com/google/uuid"
)

func CreateData() {

	roles := []models.Role{
		{ ID: uuid.New(), Name: "Admin", Description: "Superuser with full system access. Can manage flights, users, payments, and bookings." },
		{ ID: uuid.New(), Name: "Customer", Description: "User who purchases airline tickets, manages their bookings, and makes payments." },
		{ ID: uuid.New(), Name: "Airline Agent", Description: "Airline employee who can manage flights, seat availability, and verify passengers." },
		{ ID: uuid.New(), Name: "Support", Description: "Support staff who assist customers with booking and payment issues."},
		{ ID: uuid.New(), Name: "Guest", Description: "Unregistered user who can view flight schedules and prices." },
	}

	for _, role := range roles {
		DB.Create(&role)
	}

}