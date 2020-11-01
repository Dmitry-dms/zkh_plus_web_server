package service

import (
	"github.com/dmitry-dms/rest-gin/models"
	"github.com/dmitry-dms/rest-gin/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)                    //вернёт id или ошибку
	GenerateToken(email string, password string) (string, error) // вернет токен
}
type UserList interface {
}

type Service struct {
	Authorization
	UserList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
