package repository

import (
	"errors"
	"fmt"
	"github.com/dmitry-dms/rest-gin/models"
	"github.com/jmoiron/sqlx"
)

type UserRequestsPostgres struct {
	db *sqlx.DB
}

func NewUserRequestsPostgres(db *sqlx.DB) *UserRequestsPostgres {
	return &UserRequestsPostgres{db: db}
}

func (r *UserRequestsPostgres) UpdateUserCompany(userId, companyId int) error {
	query := fmt.Sprintf("UPDATE %s SET company_id=$1 WHERE user_id = $2", usersTable)
	row := r.db.QueryRow(query, companyId, userId)
	if row == nil {
		return errors.New("failed to update company_id")
	}
	return nil
}

func (r *UserRequestsPostgres) CreateUserAddress(userId int, address models.UserAddress) (int, error) {
	var addressId int
	query := fmt.Sprintf("INSERT INTO %s (user_id, city, street, home_number, flat) values ($1, $2, $3, $4, $5) RETURNING address_id", usersAddressTable)
	row := r.db.QueryRow(query, userId, address.City, address.Street, address.HomeNumber, address.Flat)
	if err := row.Scan(&addressId); err != nil {
		return 0, err
	}
	return addressId, nil
}

func (r *UserRequestsPostgres) GetAllUserAddress(userId int) ([]models.UserAddress, error) {
	var lists []models.UserAddress
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", usersAddressTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}
