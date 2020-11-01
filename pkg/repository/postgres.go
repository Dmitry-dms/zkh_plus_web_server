package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable       = "users"
	usersListTable   = "users_list"
	electricityTable = "electricity_lists"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s port=%s password=%s host=%s dbname=%s  sslmode=%s",
		cfg.Username,
		cfg.Port,
		cfg.Password,
		cfg.Host,
		cfg.DBName,
		cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	return db, nil
}
