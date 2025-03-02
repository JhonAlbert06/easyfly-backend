package controllers

import (
	"net/http"

	"github.com/JhonAlbert06/easyfly-backend/initializers"
	"github.com/JhonAlbert06/easyfly-backend/models"
	"github.com/JhonAlbert06/easyfly-backend/responses"
	"github.com/gin-gonic/gin"
)

func GetAirports(c *gin.Context) {
	var airports []models.Airport
	err := initializers.DB.Find(&airports).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get airports",
		})

		return
	}

	var resAirports []responses.ResAirport
	for _, airport := range airports {
		resAirports = append(resAirports, responses.GetResAirport(airport))
	}

	if(len(resAirports) == 0){
		c.JSON(http.StatusOK, []responses.ResAirport{})
		return
	}

	c.JSON(http.StatusOK, resAirports)
}