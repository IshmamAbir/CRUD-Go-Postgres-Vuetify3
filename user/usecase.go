package user

import "backend/models"

// usecase interface
type Usecase interface {
	GetAllUsers() ([]models.User, error)

	GetUserById(id int) (*models.User, error)

	CountAllUsers(map[string]interface{}) (int64, error)

	SaveUser(user *models.User) error

	UpdateUser(id int, user *models.User) error

	DeleteUer(id int) error
}
