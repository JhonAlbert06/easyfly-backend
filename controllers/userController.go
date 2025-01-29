package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/JhonAlbert06/easyfly-backend/initializers"
	"github.com/JhonAlbert06/easyfly-backend/models"
	"github.com/JhonAlbert06/easyfly-backend/responses"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var body struct {
		Email string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
		FullName string `json:"fullName" form:"fullName"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	if body.Email == "" || body.Password == "" || body.FullName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please fill all fields",
		})

		return
	}

	//Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	// get the Guest role by name
	var role models.Role
	err = initializers.DB.First(&role, "name = ?", "Guest").Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get role",
		})
		return
	}

	user := models.User{
		ID:              uuid.New(),
		Email:        body.Email,
		Password:        string(hash),
		FullName:        body.FullName,
		PasswordVersion: uuid.New(),
		RoleId: role.ID,
		IsEnabled: true,
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

func Login(c *gin.Context) {

	var body struct {
		Email string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid User Name",
		})

		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid phone number or password",
		})

		return
	}

	//Generate jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":         user.ID,
		"exp":         time.Now().Add(time.Hour * 24 * 30).Unix(),
		"pwd_version": user.PasswordVersion,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})

		return
	}

	//Sent it back
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func ChangePassword(c *gin.Context) {
	var body struct {
		CurrentPassword string `json:"currentPassword" form:"currentPassword"`
		NewPassword     string `json:"newPassword" form:"newPassword"`
	}

	if c.Bind(&body) != nil {
		fmt.Println(body)
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body",
		})
		return
	}

	user, _ := c.Get("user")
	if user.(models.User).ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user",
		})
		return
	}

	// Hash the new password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.NewPassword), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// Compare the current password with the stored hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.(models.User).Password), []byte(body.CurrentPassword)); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Incorrect current password",
		})
		return
	}

	// Change the password and save the user in the database
	if u, ok := user.(models.User); ok {
		u.Password = string(hash)
		u.PasswordVersion = uuid.New()
		result := initializers.DB.Save(&u)

		if result.Error != nil {
			fmt.Println(result.Error)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to change the password",
			})
			return
		}

		// Respond with a success message
		c.JSON(http.StatusOK, gin.H{})
	} else {
		fmt.Println(ok)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to load user",
		})
	}
}

func LoadUser(c *gin.Context) {
	user, _ := c.Get("user")

	if u, ok := user.(models.User); ok {
		resUser := responses.GetResUser(u)
		c.JSON(http.StatusOK, resUser)
	} else {
		fmt.Println(ok)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to load user",
		})
	}
}