package usecase

import (
	"backend/models"
	"backend/user"

	"github.com/google/uuid"
)

type userUsecase struct {
	userRepo user.Repository
}

func NewUserUsecase(ur user.Repository) user.Usecase {
	return &userUsecase{
		userRepo: ur,
	}
}

// GetAllUsers implements user.Usecase
func (uu *userUsecase) GetAllUsers() ([]models.User, error) {
	users, err := uu.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserById implements user.Usecase
func (uu *userUsecase) GetUserById(id uuid.UUID) (*models.User, error) {
	user, err := uu.userRepo.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser implements user.Usecase
func (uu *userUsecase) UpdateUser(id uuid.UUID, user *models.User) error {
	savedUser, err := uu.userRepo.GetUserById(id)
	if err != nil {
		return err
	}
	savedUser.FirstName = user.FirstName
	savedUser.LastName = user.LastName
	savedUser.Email = user.Email

	savedErr := uu.userRepo.UpdateUser(savedUser)
	if savedErr != nil {
		return savedErr
	}
	user.Id = savedUser.Id
	return nil
}

// SaveUser implements user.Usecase
func (uu *userUsecase) SaveUser(user *models.User) error {
	user.Id = uuid.New()
	err := uu.userRepo.SaveUser(user)
	return err
}

// CountAllUsers implements user.Usecase
func (*userUsecase) CountAllUsers(map[string]interface{}) (int64, error) {
	panic("unimplemented")
}

// DeleteUer implements user.Usecase
func (uu *userUsecase) DeleteUer(id uuid.UUID) error {
	user, err := uu.GetUserById(id)
	if err != nil {
		return err
	}
	err = uu.userRepo.DeleteUer(user)
	if err != nil {
		return err
	}
	return nil
}
