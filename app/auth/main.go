package main

import (
	"efishery-be-test/pkg/logging"
	"efishery-be-test/pkg/security"
	"efishery-be-test/app/auth/models"
	"efishery-be-test/app/auth/repository"
	"efishery-be-test/app/auth/service"
	"efishery-be-test/app/auth/api"

	"github.com/jessevdk/go-flags"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	schema = `DROP TABLE IF EXISTS users;
	CREATE TABLE users (
		phone TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		role TEXT NOT NULL,
		password TEXT NOT NULL
	);`
	cfg models.Configuration
)

// @title Auth Service API
// @version 1.0
// @description This is a simple auth service for generating and verifying JWT.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	_, err := flags.Parse(&cfg)
	if err != nil {
		panic(err)
	}
	logger := logging.NewLogger(cfg.Service)
	authenticator := security.NewAuthenticator(cfg.Secret)
	db := initSchema(cfg.DB, logger)
	repo := repository.NewRepository(db, logger)
	service := service.NewService(repo, &authenticator, logger)
	rest := api.NewRESTAPI(cfg.APIHost, cfg.APIPort, service)
	rest.Run()
}

func initSchema(file string, logger *logging.Logger) (db *sqlx.DB) {
	db, err := sqlx.Connect("sqlite3", file)
	if err != nil {
		logger.Panic("db.Connect", err)
	}
	db.MustExec(schema)
	return db
}
