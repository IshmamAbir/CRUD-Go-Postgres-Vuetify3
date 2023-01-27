package repository

import (
	"backend/company"
	"backend/models"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PsqlCompanyRepository struct {
	DB *gorm.DB
}

// DeleteCompany implements company.Repository
func (repo *PsqlCompanyRepository) DeleteCompany(company *models.Company) error {
	err := repo.DB.Delete(company).Error
	if err != nil {
		return errors.New("unable to delete the company")
	}
	return nil
}

// GetAllCompany implements company.Repository
func (repo *PsqlCompanyRepository) GetAllCompanies() ([]models.Company, error) {
	var companies []models.Company
	repo.DB.Find(&companies)
	if companies == nil {
		return nil, errors.New("no company found")
	}
	return companies, nil
}

// GetCompanyById implements company.Repository
func (repo *PsqlCompanyRepository) GetCompanyById(id uuid.UUID) (*models.Company, error) {
	var company models.Company
	repo.DB.First(&company, id)
	if company.Id == uuid.Nil {
		return nil, errors.New("company not found")
	}
	return &company, nil
}

// SaveCompany implements company.Repository
func (repo *PsqlCompanyRepository) SaveCompany(company *models.Company) error {
	err := repo.DB.Create(company).Error
	if err != nil {
		return errors.New("couldn't save to database")
	}
	return nil
}

// UpdateCompany implements company.Repository
func (repo *PsqlCompanyRepository) UpdateCompany(company *models.Company) error {
	if err := repo.DB.Save(company).Error; err != nil {
		return errors.New("couldn't update company info")
	}
	return nil
}

func NewPsqlCompanyRepository(db *gorm.DB) company.Repository {
	return &PsqlCompanyRepository{
		DB: db,
	}
}
