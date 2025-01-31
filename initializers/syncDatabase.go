package initializers

import (
	"fmt"

	"github.com/JhonAlbert06/easyfly-backend/models"
)

func SyncDatabase() {
	err := DB.AutoMigrate(
		&models.Airport{},
		&models.Booking{},
		&models.Flight{},
		&models.Payment{},
		&models.Role{},
		&models.Seat{},
		&models.Ticket{},
		&models.User{},
	)
	if err != nil {
		fmt.Println(err)
		return
	}
}