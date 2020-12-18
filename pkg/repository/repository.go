package repository

import (
	"github.com/dmitry-dms/rest-gin/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User, companyId int) (int, error)
	GetUser(email, password string) (models.User, error)
	CreateCompany(owner models.Company) (int, error)
	GetCompany(email, password string) (models.Company, error)
}
type CompanyList interface {
	GetAllCompanies() ([]models.Company, error)
	GetCompanyById(companyId int) (models.Company, error)
	CreateNotification(companyId int, notification models.Notification) error
}
type UserRequest interface {
	UpdateUserCompany(userId, companyId int) error
	CreateUserAddress(userId int, address models.UserAddress) (int, error)
	GetAllUserAddress(userId int) ([]models.UserAddress, error)
	InputVolumes(userId int, volume models.DataVolume) (float32, error)
	GetUsersValuesByYearAndMonth(userId, year, month int) ([]models.DataVolume, error)
	GetAllUserValues(userId int) ([]models.DataVolume, error)
	GetNotifications(companyId int) ([]models.Notification, error)
	GetUserInfo(userId int) (models.User, error)
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
