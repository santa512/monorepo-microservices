package httphandler

import "github.com/santa512/monorepo-microservices/auth/model"

type (
	// For Post/Put - /users
	UserResource struct {
		Data model.User `json:"data"`
	}
)
