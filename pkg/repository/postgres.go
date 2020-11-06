package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const (
	usersTable        = "users"
	usersAddressTable = "address"
	electricityTable  = "electricity_lists"
	companyTable      = "company"
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
	query := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)
	logrus.Printf("URL to database %s", query)
	db, err := sqlx.Connect("postgres", query)
	if err != nil {
		return nil, err
	}
	return db, nil
}
