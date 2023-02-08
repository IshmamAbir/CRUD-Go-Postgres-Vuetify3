package repository

import (
	"errors"
	"strconv"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"backend/models"
	"backend/user"
)

// postgresql repository struct
type psqlUserRepository struct {
	DB *gorm.DB
}

// NewPsqlUserRepository  will create an object that represent the user.Repository interface
// implement all methods from user.repository interface
func NewPsqlUserRepository(db *gorm.DB) user.Repository {
	return &psqlUserRepository{
		DB: db,
	}
}

// GetAllUsers implements user.Repository
func (repo *psqlUserRepository) GetAllUsers(queryParams map[string]interface{}) ([]models.User, error) {
	orderByColumns := ""
	tx := repo.DB

	if queryParams["order"] != nil {
		orderByColumns = queryParams["order"].(string)
		delete(queryParams, "order")
	}
	if queryParams["limit"] != nil {
		limit := queryParams["limit"].(string)
		limitSize, err := strconv.Atoi(limit)
		if err != nil {
			return nil, errors.New("limit error")
		}
		tx = tx.Limit(limitSize)
		delete(queryParams, "limit")
	}
	if queryParams["offset"] != nil {
		offset := queryParams["offset"].(string)
		offsetSize, err := strconv.Atoi(offset)
		if err != nil {
			return nil, errors.New("limit error")
		}
		tx = tx.Offset(offsetSize)
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
	}

	var users []models.User
	// repo.DB.Model(&models.User{}).Preload("Company").Find(&users)
	tx.Order(orderByColumns).Model(&models.User{}).Preload("Company").Find(&users)
	if users == nil {
		return nil, errors.New("no user found")
	}
	return users, nil
}

// GetUserById implements user.Repository
func (repo *psqlUserRepository) GetUserById(id uuid.UUID) (*models.User, error) {
	var user models.User
	repo.DB.Model(&models.User{}).Preload("Company").First(&user, id)
	if user.Id == uuid.Nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

// SaveUser implements user.Repository
func (repo *psqlUserRepository) SaveUser(user *models.User) error {
	err := repo.DB.Create(user).Error
	if err != nil {
		return errors.New("couldn't save to database")
	}
	return nil
}

// CountAllUsers implements user.Repository
func (repo *psqlUserRepository) CountAllUsers(map[string]interface{}) (int64, error) {
	panic("unimplemented")
}

// DeleteUer implements user.Repository
func (repo *psqlUserRepository) DeleteUer(user *models.User) error {
	err := repo.DB.Delete(user).Error
	if err != nil {
		return errors.New("unable to delete the user")
	}
	return nil
}

// UpdateUser implements user.Repository
func (repo *psqlUserRepository) UpdateUser(user *models.User) error {
	if err := repo.DB.Save(user).Error; err != nil {
		return errors.New("couldn't update user info")
	}
	return nil
}
