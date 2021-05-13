package service

import (
	"github.com/Dmitry-dms/zkh-plus/models"
	"github.com/Dmitry-dms/zkh-plus/pkg/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

type Authorization interface {
	CreateUser(user models.User, companyId primitive.ObjectID) (interface{}, error)     //вернёт id или ошибку
	GenerateToken(email string, password string) (string, error) // вернет токен
	ParseToken(token string) (interface{}, error)                        //вернёт id при успешном парсинге
	CreateCompany(owner models.Company) (interface{}, error)
	GenerateCompanyOwnerToken(email, password string) (string, error)
}
type CompanyList interface {
	GetAllCompanies() ([]models.Company, error)
	GetCompanyById(companyId int) (models.Company, error)
	CreateNotification(companyId int, notification models.Notification) error
}
 type UserRequest interface {
// 	UpdateUserCompany(userId, companyId int) error
// 	CreateUserAddress(userId int, address models.UserAddress) (int, error)
// 	GetAllUserAddress(userId int) ([]models.UserAddress, error)
// 	InputVolumes(userId int, volume models.DataVolume) (models.VolumeResponse, error)
// 	GetUsersValuesByYearAndMonth(userId, year, month int) ([]models.DataVolume, error)
// 	GetAllUserValues(userId int) ([]models.DataVolume, error)
// 	GetNotifications(companyId int) ([]models.Notification, error)
 	GetUserInfo(userId primitive.ObjectID) (models.User, models.Company, error)
// 	GetUsersLastVolume(userId int) ([]models.DataVolume, error)
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
