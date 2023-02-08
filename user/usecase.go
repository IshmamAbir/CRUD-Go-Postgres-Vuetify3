package user

import (
	"backend/models"

	"github.com/google/uuid"
)

// usecase interface
type Usecase interface {
	GetAllUsers(queryParams map[string]interface{}) ([]models.User, error)

	GetUserById(id uuid.UUID) (*models.User, error)

	CountAllUsers(map[string]interface{}) (int64, error)

	SaveUser(user *models.User) error

	UpdateUser(id uuid.UUID, user *models.User) error

	DeleteUer(id uuid.UUID) error
}
