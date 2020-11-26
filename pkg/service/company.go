package service

import (
	"fmt"
	"github.com/dmitry-dms/rest-gin/models"
	"github.com/dmitry-dms/rest-gin/pkg/repository"
	"time"
)

type CompanyListService struct {
	repo repository.CompanyList
}

func NewCompanyListService(repo repository.CompanyList) *CompanyListService {
	return &CompanyListService{repo: repo}
}

func (s *CompanyListService) GetAllCompanies() ([]models.Company, error) {
	return s.repo.GetAllCompanies()
}

func (s *CompanyListService) GetCompanyById(companyId int) (models.Company, error) {
	return s.repo.GetCompanyById(companyId)
}

func (s *CompanyListService) CreateNotification(companyId int, notification models.Notification) error {
	t := time.Now()
	fullDate := fmt.Sprintf(t.Format("2006-01-02"))
	notification.FullDate = fullDate
	return s.repo.CreateNotification(companyId, notification)
}
