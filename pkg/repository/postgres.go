package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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
	query := fmt.Sprintf("user=%s port=%s password=%s host=%s dbname=%s  sslmode=%s",
		cfg.Username,
		cfg.Port,
		cfg.Password,
		cfg.Host,
		cfg.DBName,
		cfg.SSLMode)
	logrus.Debugf("%s", query)
	db, err := sqlx.Connect("postgres", query)
	if err != nil {
		return nil, err
	}
	return db, nil
}
