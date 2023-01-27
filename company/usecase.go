package company

import (
	"backend/models"

	"github.com/google/uuid"
)

// usecase interface
type Usecase interface {
	GetAllCompanies() ([]models.Company, error)

	GetCompanyById(id uuid.UUID) (*models.Company, error)

	SaveCompany(company *models.Company) error

	UpdateCompany(id uuid.UUID, company *models.Company) error

	DeleteCompany(id uuid.UUID) error
}
