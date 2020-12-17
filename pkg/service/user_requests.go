package service

import (
	"fmt"
	"github.com/dmitry-dms/rest-gin/models"
	"github.com/dmitry-dms/rest-gin/pkg/repository"
	"strings"
	"time"
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
func (s *UserRequestService) InputVolumes(userId int, volume models.DataVolume) error {
	if err := volume.Validate(); err != nil {
		return err
	}
	t := time.Now()
	fullDate := fmt.Sprintf(t.Format("2006-01-02"))
	result := strings.Split(fullDate, "-")
	volume.FullDate = fullDate
	volume.Year = result[0]
	volume.Month = result[1]
	volume.Day = result[2]
	return s.repo.InputVolumes(userId, volume)
}
func (s *UserRequestService) GetUsersValuesByYearAndMonth(userId, year, month int) ([]models.DataVolume, error) {
	return s.repo.GetUsersValuesByYearAndMonth(userId, year, month)
}

func (s *UserRequestService) GetAllUserValues(userId int) ([]models.DataVolume, error) {
	return s.repo.GetAllUserValues(userId)
}
func (s *UserRequestService) GetNotifications(companyId int) ([]models.Notification, error) {
	return s.repo.GetNotifications(companyId)
}
func (s *UserRequestService) GetUserInfo(userId int) (models.User, error) {
	return s.repo.GetUserInfo(userId)
}
