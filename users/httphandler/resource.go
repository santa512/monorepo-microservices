package httphandler

import "github.com/santa512/monorepo-microservices/users/model"

type (
	// For Get - /users
	UsersResource struct {
		Data []model.User `json:"data"`
	}
	// For Post/Put - /users
	UserResource struct {
		Data model.User `json:"data"`
	}
)
