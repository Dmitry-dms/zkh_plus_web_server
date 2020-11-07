package service

import (
	"github.com/dmitry-dms/rest-gin/models"
	"github.com/dmitry-dms/rest-gin/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)                    //вернёт id или ошибку
	GenerateToken(email string, password string) (string, error) // вернет токен
	ParseToken(token string) (int, error)                        //вернёт id при успешном парсинге
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

type Service struct {
	Authorization
	CompanyList
	UserRequest
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		CompanyList:   NewCompanyListService(repos.CompanyList),
		UserRequest:   NewUserRequestService(repos.UserRequest),
	}
}
