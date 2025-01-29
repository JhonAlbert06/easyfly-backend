package responses

import (
	"github.com/JhonAlbert06/easyfly-backend/initializers"
	"github.com/JhonAlbert06/easyfly-backend/models"
	"github.com/google/uuid"
)

type ResRole struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

func GetResRole(roleId uuid.UUID) ResRole {

	var DB = initializers.DB

	var role models.Role
	DB.First(&role, "id = ?", roleId)

	return ResRole{
		Name: role.Name,
		Description: role.Description,
	}
}