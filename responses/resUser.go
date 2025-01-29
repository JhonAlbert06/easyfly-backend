package responses

import "github.com/JhonAlbert06/easyfly-backend/models"



type ResUser struct {
	ID       string `json:"id"`
	Email string `json:"email"`
	FullName string `json:"fullName"`
}

func GetResUser(user models.User) ResUser {
	return ResUser{
		ID:       user.ID.String(),
		Email: user.Email,
		FullName: user.FullName,
	}
}