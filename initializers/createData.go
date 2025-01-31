package initializers

import (
	"errors"
	"log"

	"github.com/JhonAlbert06/easyfly-backend/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
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
		var existingRole models.Role
		result := DB.Where("name = ?", role.Name).First(&existingRole)
	
		if result.Error != nil { // Si no se encuentra, lo creamos
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				DB.Create(&role)
			} else {
				log.Println("Error checking role:", result.Error)
			}
		}
	}

	airports := []models.Airport{
		{ ID: uuid.New(), Name: "El Dorado International Airport", Code: "BOG", City: "Bogotá", Country: "Colombia" },
		{ ID: uuid.New(), Name: "Miami International Airport", Code: "MIA", City: "Miami", Country: "United States" },
		{ ID: uuid.New(), Name: "John F. Kennedy International Airport", Code: "JFK", City: "New York", Country: "United States" },
		{ ID: uuid.New(), Name: "Heathrow Airport", Code: "LHR", City: "London", Country: "United Kingdom" },
		{ ID: uuid.New(), Name: "Charles de Gaulle Airport", Code: "CDG", City: "Paris", Country: "France" },
		{ ID: uuid.New(), Name: "Los Angeles International Airport", Code: "LAX", City: "Los Angeles", Country: "United States" },
		{ ID: uuid.New(), Name: "Dubai International Airport", Code: "DXB", City: "Dubai", Country: "United Arab Emirates" },
		{ ID: uuid.New(), Name: "Madrid-Barajas Adolfo Suárez Airport", Code: "MAD", City: "Madrid", Country: "Spain" },
		{ ID: uuid.New(), Name: "São Paulo/Guarulhos International Airport", Code: "GRU", City: "São Paulo", Country: "Brazil" },
		{ ID: uuid.New(), Name: "Toronto Pearson International Airport", Code: "YYZ", City: "Toronto", Country: "Canada" },
		{ ID: uuid.New(), Name: "Tokyo Haneda Airport", Code: "HND", City: "Tokyo", Country: "Japan" },
		{ ID: uuid.New(), Name: "Sydney Kingsford Smith Airport", Code: "SYD", City: "Sydney", Country: "Australia" },
		{ ID: uuid.New(), Name: "Frankfurt Airport", Code: "FRA", City: "Frankfurt", Country: "Germany" },
		{ ID: uuid.New(), Name: "Hong Kong International Airport", Code: "HKG", City: "Hong Kong", Country: "China" },
		{ ID: uuid.New(), Name: "Changi Airport", Code: "SIN", City: "Singapore", Country: "Singapore" },
		{ ID: uuid.New(), Name: "Amsterdam Schiphol Airport", Code: "AMS", City: "Amsterdam", Country: "Netherlands" },
		{ ID: uuid.New(), Name: "Mexico City International Airport", Code: "MEX", City: "Mexico City", Country: "Mexico" },
		{ ID: uuid.New(), Name: "Beijing Capital International Airport", Code: "PEK", City: "Beijing", Country: "China" },
		{ ID: uuid.New(), Name: "O'Hare International Airport", Code: "ORD", City: "Chicago", Country: "United States" },
		{ ID: uuid.New(), Name: "San Francisco International Airport", Code: "SFO", City: "San Francisco", Country: "United States" },
		{ ID: uuid.New(), Name: "Seoul Incheon International Airport", Code: "ICN", City: "Seoul", Country: "South Korea" },
		{ ID: uuid.New(), Name: "Munich Airport", Code: "MUC", City: "Munich", Country: "Germany" },
		{ ID: uuid.New(), Name: "Zürich Airport", Code: "ZRH", City: "Zürich", Country: "Switzerland" },
		{ ID: uuid.New(), Name: "Vienna International Airport", Code: "VIE", City: "Vienna", Country: "Austria" },
		{ ID: uuid.New(), Name: "Narita International Airport", Code: "NRT", City: "Tokyo", Country: "Japan" },
		{ ID: uuid.New(), Name: "Doha Hamad International Airport", Code: "DOH", City: "Doha", Country: "Qatar" },
		{ ID: uuid.New(), Name: "Indira Gandhi International Airport", Code: "DEL", City: "New Delhi", Country: "India" },
		{ ID: uuid.New(), Name: "Moscow Sheremetyevo International Airport", Code: "SVO", City: "Moscow", Country: "Russia" },
		{ ID: uuid.New(), Name: "Istanbul Airport", Code: "IST", City: "Istanbul", Country: "Turkey" },
		{ ID: uuid.New(), Name: "Las Américas International Airport", Code: "SDQ", City: "Santo Domingo", Country: "Dominican Republic" },
		{ ID: uuid.New(), Name: "Punta Cana International Airport", Code: "PUJ", City: "Punta Cana", Country: "Dominican Republic" },
		{ ID: uuid.New(), Name: "Cibao International Airport", Code: "STI", City: "Santiago", Country: "Dominican Republic" },
		{ ID: uuid.New(), Name: "Gregorio Luperón International Airport", Code: "POP", City: "Puerto Plata", Country: "Dominican Republic" },
	}

	for _, airport := range airports {
		var existingAirport models.Airport
		result := DB.Where("code = ?", airport.Code).First(&existingAirport)
	
		if result.Error != nil { // Si no se encuentra, lo creamos
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				DB.Create(&airport)
			} else {
				log.Println("Error checking airport:", result.Error)
			}
		}
	}

}