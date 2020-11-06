package service

import (
	"github.com/dmitry-dms/rest-gin/models"
	"github.com/dmitry-dms/rest-gin/pkg/repository"
)

type UserRequestService struct {
	repo repository.UserRequest
}

func NewUserRequestService(repo repository.UserRequest) *UserRequestService {
	return &UserRequestService{repo: repo}
}

func (s *UserRequestService) UpdateUserCompany(userId, companyId int) error {
	return s.repo.UpdateUserCompany(userId, companyId)
}

func (s *UserRequestService) CreateUserAddress(userId int, address models.UserAddress) (int, error) {
	return s.repo.CreateUserAddress(userId, address)
}

func (s *UserRequestService) GetAllUserAddress(userId int) ([]models.UserAddress, error) {
	return s.repo.GetAllUserAddress(userId)
}
