package controllers

import (
	"net/http"

	"github.com/JhonAlbert06/easyfly-backend/initializers"
	"github.com/JhonAlbert06/easyfly-backend/models"
	"github.com/JhonAlbert06/easyfly-backend/responses"
	"github.com/gin-gonic/gin"
)

func GetFlights(c *gin.Context) {
	var flights []models.Flight
	err := initializers.DB.Find(&flights).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get flights",
		})

		return
	}

	//c.JSON(http.StatusOK, flights)

	var resFlight []responses.ResFlight
	for _, flight := range flights {
		resFlight = append(resFlight, responses.GetResFlight(flight))
	}

	if(len(resFlight) == 0){
		c.JSON(http.StatusOK, []responses.ResFlight{})
		return
	}

	c.JSON(http.StatusOK, resFlight)

}