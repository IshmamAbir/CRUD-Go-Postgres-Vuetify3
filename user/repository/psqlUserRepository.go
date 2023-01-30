package repository

import (
	"errors"

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
func (repo *psqlUserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	repo.DB.Model(&models.User{}).Preload("Company").Find(&users)
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
