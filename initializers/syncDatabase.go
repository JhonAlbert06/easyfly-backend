package initializers

import (
	"fmt"

	"github.com/JhonAlbert06/easyfly-backend/models"
)

func SyncDatabase() {
	err := DB.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		fmt.Println(err)
		return
	}
}