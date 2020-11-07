package repository

import (
	"github.com/dmitry-dms/rest-gin/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(email, password string) (models.User, error)
}
type CompanyList interface {
	GetAllCompanies() ([]models.Company, error)
	GetCompanyById(companyId int) (models.Company, error)
}
type UserRequest interface {
	UpdateUserCompany(userId, companyId int) error
	CreateUserAddress(userId int, address models.UserAddress) (int, error)
	GetAllUserAddress(userId int) ([]models.UserAddress, error)
	InputVolumes(userId int, volume models.DataVolume) error
	GetUsersValuesByYearAndMonth(userId, year, month int) ([]models.DataVolume, error)
	GetAllUserValues(userId int) ([]models.DataVolume, error)
}
type Repository struct {
	Authorization
	CompanyList
	UserRequest
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		CompanyList:   NewCompanyListPostgres(db),
		UserRequest:   NewUserRequestsPostgres(db),
	}
}
