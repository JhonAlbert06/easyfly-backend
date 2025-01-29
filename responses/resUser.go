package responses

import "github.com/JhonAlbert06/easyfly-backend/models"

type ResUser struct {
	ID       string `json:"id"`
	Email string `json:"email"`
	FullName string `json:"fullName"`
	Role ResRole `json:"role"`
	IsEnabled bool `json:"isEnabled"`
}

func GetResUser(user models.User) ResUser {

	role := GetResRole(user.RoleId)

	return ResUser{
		ID:       user.ID.String(),
		Email: user.Email,
		FullName: user.FullName,
		Role: role,
		IsEnabled: user.IsEnabled,
	}
}