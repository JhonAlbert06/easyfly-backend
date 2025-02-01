package responses

import (
	"github.com/JhonAlbert06/easyfly-backend/models"
	"github.com/google/uuid"
)

type ResAirport struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Code    string    `json:"code"`
	City    string    `json:"city"`
	Country string    `json:"country"`
}

func GetResAirport(airport models.Airport) ResAirport {
	return ResAirport{
		ID:      airport.ID,
		Name:    airport.Name,
		Code:    airport.Code,
		City:    airport.City,
		Country: airport.Country,
	}
}
