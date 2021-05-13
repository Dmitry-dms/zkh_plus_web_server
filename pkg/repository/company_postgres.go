package repository

import (
	"errors"
	"fmt"
	"github.com/Dmitry-dms/zkh-plus/models"
	"github.com/jmoiron/sqlx"
)

type CompanyListPostgres struct {
	db *sqlx.DB
}

func NewCompanyListPostgres(db *sqlx.DB) *CompanyListPostgres {
	return &CompanyListPostgres{db: db}
}

func (r *CompanyListPostgres) GetAllCompanies() ([]models.Company, error) {
	var lists []models.Company
	query := fmt.Sprintf("SELECT * FROM %s", companyTable)
	err := r.db.Select(&lists, query)

	return lists, err
}

func (r *CompanyListPostgres) GetCompanyById(companyId int) (models.Company, error) {
	var company models.Company
	query := fmt.Sprintf("SELECT * FROM %s WHERE company_id=$1", companyTable)
	err := r.db.Get(&company, query, companyId)
	return company, err
}

func (r *CompanyListPostgres) CreateNotification(companyId int, notification models.Notification) error {
	query := fmt.Sprintf("INSERT INTO %s (company_id, article, description, date_full) values ($1, $2, $3, $4)", notificationsTable)
	row := r.db.QueryRow(query, companyId, notification.Article, notification.Description, notification.FullDate)
	if row == nil {
		return errors.New("failed to create notification")
	}
	return nil
}
