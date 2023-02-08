package usecase

import (
	"backend/company"
	"backend/models"

	"github.com/google/uuid"
)

type CompanyUsecase struct {
	companyRepo company.Repository
}

// DeleteCompany implements company.Usecase
func (cu *CompanyUsecase) DeleteCompany(id uuid.UUID) error {
	company, err := cu.companyRepo.GetCompanyById(id)
	if err != nil {
		return err
	}
	err = cu.companyRepo.DeleteCompany(company)
	if err != nil {
		return err
	}
	return nil
}

// GetAllCompany implements company.Usecase
func (cu *CompanyUsecase) GetAllCompanies(queryParams map[string]interface{}) ([]models.Company, error) {
	companies, err := cu.companyRepo.GetAllCompanies(queryParams)
	if err != nil {
		return nil, err
	}
	return companies, nil
}

// GetCompanyById implements company.Usecase
func (cu *CompanyUsecase) GetCompanyById(id uuid.UUID) (*models.Company, error) {
	company, err := cu.companyRepo.GetCompanyById(id)
	if err != nil {
		return nil, err
	}
	return company, nil
}

// SaveCompany implements company.Usecase
func (cu *CompanyUsecase) SaveCompany(company *models.Company) error {
	err := cu.companyRepo.SaveCompany(company)
	return err
}

// UpdateCompany implements company.Usecase
func (cu *CompanyUsecase) UpdateCompany(id uuid.UUID, company *models.Company) error {
	savedCompany, err := cu.companyRepo.GetCompanyById(id)
	if err != nil {
		return err
	}
	savedCompany.Name = company.Name
	savedCompany.Location = company.Location

	savedErr := cu.companyRepo.UpdateCompany(savedCompany)
	if savedErr != nil {
		return savedErr
	}
	company.Id = savedCompany.Id
	return nil
}

func NewCompanyUsecase(cr company.Repository) company.Usecase {
	return &CompanyUsecase{
		companyRepo: cr,
	}
}
