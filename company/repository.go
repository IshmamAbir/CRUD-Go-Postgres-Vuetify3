package company

import (
	"backend/models"

	"github.com/google/uuid"
)

// repository interface
type Repository interface {
	GetAllCompanies() ([]models.Company, error)

	GetCompanyById(id uuid.UUID) (*models.Company, error)

	SaveCompany(company *models.Company) error

	UpdateCompany(company *models.Company) error

	DeleteCompany(company *models.Company) error
}
