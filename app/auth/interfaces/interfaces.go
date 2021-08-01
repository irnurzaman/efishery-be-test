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
	RegisterUser(user models.ReqRegisterUser) (pwd string, err error)
	VerifyUser(user models.ReqVerifyUser) (token string, err error)
	VerifyToken(token models.ReqVerifyToken) (claims models.Token, err error)
}