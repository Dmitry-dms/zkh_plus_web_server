package service

import (
	"github.com/dmitry-dms/rest-gin/models"
	"github.com/dmitry-dms/rest-gin/pkg/repository"
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
