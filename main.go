package main

import (
	"github.com/JhonAlbert06/easyfly-backend/controllers"
	"github.com/JhonAlbert06/easyfly-backend/initializers"
	"github.com/JhonAlbert06/easyfly-backend/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()  
	initializers.SyncDatabase()
	initializers.CreateData()
}

func main() {

	r := gin.Default()

	// Allowing CORS from any origin
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// User
	r.POST("/user/signup", controllers.SignUp)
	r.POST("/user/login", controllers.Login)

	r.GET("/user/loadUser", middleware.RequireAuth, controllers.LoadUser)
	r.POST("/user/changeUserPassword", middleware.RequireAuth, controllers.ChangePassword)

	// Airports
	r.GET("/airports", controllers.GetAirports)

	err := r.Run()
	if err != nil {
		return
	}
}
