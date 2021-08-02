package interfaces

import (
	"efishery-be-test/app/auth/entities"
	"efishery-be-test/app/auth/models"
)

type Repository interface {
	CreateUser(user entities.User) (err error)
	GetUser(phone string) (user entities.User, err error)
}

type Service interface {
	RegisterUser(req models.ReqRegisterUser) (pwd string, err error)
	VerifyUser(req models.ReqVerifyUser) (token string, err error)
	VerifyToken(req models.ReqVerifyToken) (claims models.Token, err error)
}