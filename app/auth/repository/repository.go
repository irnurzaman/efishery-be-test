package repository

import (
	"efishery-be-test/app/auth/entities"
	"efishery-be-test/pkg/logging"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db     *sqlx.DB
	logger *logging.Logger
}

func (r *Repository) CreateUser(user entities.User) (err error) {
	SQL := "INSERT INTO users VALUES (:phone, :name, :role, :password);"
	_, err = r.db.NamedExec(SQL, user)
	if err != nil {
		r.logger.Error("repo.CreateUser", err)
	}
	return
}

func (r *Repository) GetUser(phone string) (user entities.User, err error) {
	SQL := "SELECT phone, name, role FROM users WHERE phone = $1;"
	err = r.db.Get(&user, SQL, phone)
	if err != nil {
		r.logger.Error("repo.GetUser", err)
	}
	return
}

func NewRepository(db *sqlx.DB, logger *logging.Logger) (*Repository) {
	return &Repository{
		db: db,
		logger: logger,
	}
}
