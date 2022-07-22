package user

import "github.com/golang-api-user/user/models"

// UserUsecase is a interface for function business logic
type UserUsecase interface {
	Login(user *models.User) (*models.User, error)
	CheckMail(user *models.User) bool
	Create(user *models.User) (*models.User, error)
	FindAll() ([]*models.UserWrapper, error)
	FindByID(id int) (*models.UserWrapper, error)
	Update(user *models.User) (*models.User, error)
	Delete(id int) error
	UpdatePhoto(user *models.User) error
}
