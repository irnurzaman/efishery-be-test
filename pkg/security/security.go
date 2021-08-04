package security

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

type Authenticator struct {
	secret string
	Method jwt.SigningMethod
}

func (a Authenticator) GenerateToken(claims jwt.Claims) (stoken string, err error) {
	token := jwt.NewWithClaims(a.Method, claims)
	stoken, err = token.SignedString([]byte(a.secret))
	return
}

func (a Authenticator) ParseToken(stoken string, claims jwt.Claims) (err error) {
	token, err := jwt.ParseWithClaims(stoken, claims, func(t *jwt.Token) (interface{}, error) { return []byte(a.secret), nil })
	if err == nil && !token.Valid {
		err := errors.New("Invalid token")
		return err
	}
	return err
}

func NewAuthenticator(secret string) Authenticator {
	return Authenticator{
		secret: secret,
		Method: jwt.SigningMethodHS256,
	}
}
