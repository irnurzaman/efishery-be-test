package service

import (
	"crypto/subtle"
	"efishery-be-test/app/auth/entities"
	"efishery-be-test/app/auth/interfaces"
	"efishery-be-test/app/auth/models"
	"efishery-be-test/pkg/logging"
	"efishery-be-test/pkg/security"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type Service struct {
	repo   interfaces.Repository
	auth   *security.Authenticator
	logger *logging.Logger
}

func (s *Service) RegisterUser(req models.ReqRegisterUser) (pwd string, err error) {
	pwd = generatePassword()
	var user entities.User
	copier.Copy(&user, req)
	user.Password = pwd
	err = s.repo.CreateUser(user)
	if err != nil {
		err = fmt.Errorf("Phone number has been registered")
	}
	return
}

func (s *Service) VerifyUser(req models.ReqLoginUser) (token string, err error) {
	user, err := s.repo.GetUser(req.Phone)
	if err != nil {
		err = fmt.Errorf("Invalid authentication for phone %s", req.Phone)
		return
	}
	valid := subtle.ConstantTimeCompare([]byte(user.Password), []byte(req.Password))
	if valid != 1 {
		err = fmt.Errorf("Invalid authentication for phone %s", req.Phone)
		return
	}
	claims := models.Token{
		user.Phone,
		user.Name,
		user.Role,
		time.Now().String(),
		jwt.StandardClaims{},
	}
	token, err = s.auth.GenerateToken(claims)
	if err != nil {
		s.logger.Error("service.VerifyUser(GenerateToken)", err)
		err = fmt.Errorf("Invalid authentication for phone %s", req.Phone)
		return
	}
	return
}

func (s *Service) VerifyToken(req models.ReqVerifyToken) (claims models.Token, err error) {
	err = s.auth.ParseToken(req.Token, &claims)
	if err != nil {
		s.logger.Error("service.VerifyUser(GenerateToken)", err)
		err = fmt.Errorf("Invalid token verification")
	}
	return
}

func generatePassword() string {
	return uuid.New().String()[:4]
}

func NewService(repo interfaces.Repository, auth *security.Authenticator, logger *logging.Logger) (*Service) {
	return &Service{
		repo: repo,
		auth: auth,
		logger: logger,
	}
}