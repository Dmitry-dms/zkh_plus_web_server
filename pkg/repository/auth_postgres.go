package repository

import (
	"errors"
	"fmt"
	"github.com/dmitry-dms/rest-gin/models"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User, companyId int) (int, error) {
	var id int
	fullName := user.Surname + " " + user.Name + " " + user.Patronymic
	query := fmt.Sprintf("INSERT INTO %s (name, surname, patronymic, full_name, email, password_hash, company_id) values ($1, $2, $3, $4, $5, $6, $7) RETURNING user_id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Surname, user.Patronymic, fullName, user.Email, user.Password, companyId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE email=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, email, password)
	return user, err
}
func (r *AuthPostgres) CreateCompany(company models.Company) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, password_hash, company_name, director_full_name, company_phone, company_city, company_street, company_home_number, company_flat) values ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING company_id", companyTable)
	row := r.db.QueryRow(query, company.Email, company.Password, company.Name, company.DirectorName, company.Phone, company.City, company.Street, company.HomeNumber, company.Flat)
	if err := row.Scan(&id); err != nil {
		return 0, errors.New("account with this credentials already exists")
	}
	return id, nil
}

func (r *AuthPostgres) GetCompany(email, password string) (models.Company, error) {
	var user models.Company
	query := fmt.Sprintf("SELECT company_id FROM %s WHERE email=$1 AND password_hash=$2", companyTable)
	err := r.db.Get(&user, query, email, password)
	return user, err
}
