package repository

import (
	"github.com/Dmitry-dms/zkh-plus/models"
	//"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	CreateUser(user models.User, companyId primitive.ObjectID) (interface{}, error)
	GetUser(email, password string) (models.User, error)
	CreateCompany(owner models.Company) (interface{}, error)
	GetCompany(email, password string) (models.Company, error)
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
 	GetUserInfo(userId primitive.ObjectID) (models.User,models.Company, error)
// 	GetUsersLastVolume(userId int) ([]models.DataVolume, error)
 }

type Repository struct {
	Authorization
	CompanyList
	UserRequest
}


// func NewRepository(db *sqlx.DB) *Repository {
// 	return &Repository{
// 		Authorization: NewAuthPostgres(db),
// 		CompanyList:   NewCompanyListPostgres(db),
// 		UserRequest:   NewUserRequestsPostgres(db),
// 	}
// }

func NewMongoRepository(db *mongo.Database) *Repository {
	return &Repository{
		Authorization: NewAuthMongo(db),
		CompanyList:   NewCompanyListMongo(db),
		UserRequest:   NewUserRequestsMongo(db),
	}
}
