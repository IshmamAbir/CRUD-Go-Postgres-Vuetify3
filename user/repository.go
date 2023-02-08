package user

import (
	"backend/models"

	"github.com/google/uuid"
)

// repository interface
type Repository interface {
	GetAllUsers(queryParams map[string]interface{}) ([]models.User, error)

	GetUserById(id uuid.UUID) (*models.User, error)

	CountAllUsers(map[string]interface{}) (int64, error)

	SaveUser(user *models.User) error

	UpdateUser(user *models.User) error

	DeleteUer(user *models.User) error
}
