package responses

import (
	"fmt"
	"time"

	"github.com/JhonAlbert06/easyfly-backend/initializers"
	"github.com/JhonAlbert06/easyfly-backend/models"
	"github.com/google/uuid"
)

type ResFlight struct {
	ID      uuid.UUID `json:"id"`
	AirLine string    `json:"airline"`
	FlightNumber string `json:"flight_number"`
	AirportOrigin ResAirport `json:"airport_origin"`
	AirportDestination ResAirport `json:"airport_destination"`
	DepartureTime time.Time `json:"departure_time"`
	ArrivalTime time.Time `json:"arrival_time"`
	Price float64 `json:"price"`
	SeatsTotal int `json:"seats_total"`
	SeatsBooked int `json:"seats_booked"`
	SeatsAvailable int `json:"seats_available"`
	Status string `json:"status"`
}

func GetResFlight(flight models.Flight) ResFlight {

	oAirport := models.Airport{}
	dAirport := models.Airport{}

	err := initializers.DB.First(&oAirport, "id = ?", flight.AirportOriginID).Error
	if err != nil {
		fmt.Println("Error getting origin airport:", err)
	}

	err = initializers.DB.First(&dAirport, "id = ?", flight.AirportDestinationID).Error
	if err != nil {
		fmt.Println("Error getting destination airport:", err)
	}
	
	return ResFlight{
		ID: flight.ID,
		AirLine: flight.AirLine,
		FlightNumber: flight.FlightNumber,
		AirportOrigin: GetResAirport(oAirport),
		AirportDestination: GetResAirport(dAirport),
		DepartureTime: flight.DepartureTime,
		ArrivalTime: flight.ArrivalTime,
		Price: flight.Price,
		SeatsTotal: flight.SeatsTotal,
		SeatsBooked: flight.SeatsBooked,
		SeatsAvailable: flight.SeatsAvailable,
		Status: flight.Status,
	}

}