package repository

import (
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

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	fullName := user.Surname + " " + user.Name + " " + user.Patronymic
	query := fmt.Sprintf("INSERT INTO %s (name, surname, patronymic, full_name, email, password_hash) values ($1, $2, $3, $4, $5, $6) RETURNING user_id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Surname, user.Patronymic, fullName, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, email, password)
	return user, err
}
