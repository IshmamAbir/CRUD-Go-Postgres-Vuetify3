package user

import "backend/models"

// repository interface
type Repository interface {
	GetAllUsers() ([]models.User, error)

	GetUserById(id int) (*models.User, error)

	CountAllUsers(map[string]interface{}) (int64, error)

	SaveUser(user *models.User) error

	UpdateUser(user *models.User) error

	DeleteUer(user *models.User) error
}
