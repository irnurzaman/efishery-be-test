package models

import (
	"github.com/golang-jwt/jwt"
)


type Configuration struct {
	APIHost    string `long:"host" description:"Service REST API host" default:"0.0.0.0"`
	APIPort    int    `long:"port" description:"Service REST API port" default:"5000"`
	DB         string `long:"db" description:"SQLite filename" default:"efishery.db"`
	Secret     string `long:"secret" description:"JWT secret key" required:"true"`
	Service    string `long:"service" descrition:"Service name" required:"true"`
}

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
