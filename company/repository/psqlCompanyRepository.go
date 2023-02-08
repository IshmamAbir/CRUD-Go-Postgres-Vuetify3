package repository

import (
	"backend/company"
	"backend/models"
	"errors"
	"strconv"

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
func (repo *PsqlCompanyRepository) GetAllCompanies(queryParams map[string]interface{}) ([]models.Company, error) {
	orderByColumns := ""
	tx := repo.DB

	if queryParams["order"] != nil {
		orderByColumns = queryParams["order"].(string)
		delete(queryParams, "order")
	}
	if queryParams["limit"] != nil {
		limit := queryParams["limit"].(string)
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			return nil, errors.New("limit error")
		}
		tx = tx.Limit(limitInt)
		delete(queryParams, "limit")
	}
	if queryParams["offset"] != nil {
		offset := queryParams["offset"].(string)
		offsetInt, err := strconv.Atoi(offset)
		if err != nil {
			return nil, errors.New("offset error")
		}
		tx = tx.Offset(offsetInt)
		delete(queryParams, "offset")
	}
	// it will be added when 'created_at' column is introduced to table
	// if queryParams["start"] != nil && queryParams["end"] != nil {
	// 	tx = tx.Where("created_at BETWEEN ? AND ?", queryParams["start"], queryParams["end"])
	// 	delete(queryParams, "start")
	// 	delete(queryParams, "end")
	// }
	for k, v := range queryParams {
		tx = tx.Where(k+"=?", v.(string))
		// tx = tx.Where(k + "=" + v.(string))
	}

	var companies []models.Company
	// repo.DB.Preload("Users").Find(&companies)
	tx.Order(orderByColumns).Preload("Users").Find(&companies)
	if companies == nil {
		return nil, errors.New("no company found")
	}
	return companies, nil
}

// GetCompanyById implements company.Repository
func (repo *PsqlCompanyRepository) GetCompanyById(id uuid.UUID) (*models.Company, error) {
	var company models.Company
	repo.DB.Preload("User").First(&company, id)
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
