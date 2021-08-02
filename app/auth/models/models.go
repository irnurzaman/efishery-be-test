package models

import (
	"github.com/golang-jwt/jwt"
)

type Token struct {
	Phone     string `json:"phone"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Timestamp string `json:"string"`
	jwt.StandardClaims
}

type ReqRegisterUser struct {
	Phone string `json:"phone"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}

type RespRegisterUser struct {
	Password string `json:"password"`
}

type ReqLoginUser struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type RespLoginUser struct {
	Token string `json:"token"`
}

type ReqVerifyToken struct {
	Token string `json:"token"`
}

type RespVerifyToken struct {
	Phone     string `json:"phone"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Timestamp string `json:"timestamp"`
}
