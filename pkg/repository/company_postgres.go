package repository

import (
	"fmt"
	"github.com/dmitry-dms/rest-gin/models"
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
