package interactor

import (
	"github.com/santa512/monorepo-microservices/auth/model"
)

type AuthenticationInteractor interface {
	Login(user model.User) (model.Token, error)
	Authenticate(authToken, refreshToken, csrfToken string) (model.Token, error)
	RevokeRefreshToken(refreshToken string) error
}
