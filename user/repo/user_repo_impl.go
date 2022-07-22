package repo

import (
	"github.com/golang-api-user/user"
	"github.com/golang-api-user/user/common"
	"github.com/golang-api-user/user/models"

	"github.com/jinzhu/gorm"
)

// UserRepositoryImpl is use sharing connection
type UserRepositoryImpl struct {
	database *gorm.DB
}

// CreateRepo use for getting connection
func CreateRepo(db *gorm.DB) user.UserRepo {
	return &UserRepositoryImpl{db}
}

// Login for check a valid email when user login in
func (call *UserRepositoryImpl) Login(user *models.User) (*models.User, error) {
	users := new(models.User)

	err := call.database.Table("user").Where("email = ?", user.Email).Take(&users).Error
	if err != nil {
		// common.LogError("Login", "Error when trying geeting email address, error is =>", err)
		return nil, err
	}
	return users, nil
}

// CheckMail use for validate email is already use or nah
func (call *UserRepositoryImpl) CheckMail(user *models.User) bool {
	users := new(models.UserWrapper)
	err := call.database.Raw("SELECT * FROM \"user\" WHERE  email = ? LIMIT 1", user.Email).Scan(users).Error
	if err != nil {
		return true // email already registered
	}
	return false // email not registered
}

// Create use for create a new account user
func (call *UserRepositoryImpl) Create(user *models.User) (*models.User, error) {
	err := call.database.Table("user").Create(&user).Error
	if err != nil {
		common.LogError("Create", "Error wehn trying to create use error is =>", err)
		return nil, err
	}
	return user, nil
}

// FindByID use for search user by id
func (call *UserRepositoryImpl) FindByID(id int) (*models.UserWrapper, error) {
	user := new(models.UserWrapper)

	err := call.database.Table("user").Where("id = ?", id).First(&user).Error
	if err != nil {
		common.LogError("FindByID", "Error when trying to find data by id, error is =>", err)
		return nil, err
	}

	return user, nil
}

// FindAll use when you want to show all user data
func (call *UserRepositoryImpl) FindAll() ([]*models.UserWrapper, error) {
	userList := make([]*models.UserWrapper, 0)

	err := call.database.Table("user").Find(&userList).Error
	if err != nil {
		common.LogError("FindAll", "Error when trying to get all ada, error is =>", err)
		return nil, err
	}
	return userList, nil
}

// Update use for update user data
func (call *UserRepositoryImpl) Update(user *models.User) (*models.User, error) {

	users := new(models.User)

	err := call.database.Table("user").Where("id = ?", user.ID).First(&users).Update(&user).Error
	if err != nil {
		common.LogError("Update", "Error when trying to update user data, error is =>", err)
		return nil, err
	}
	return users, nil
}

// Delete use for delete account
func (call *UserRepositoryImpl) Delete(id int) error {
	err := call.database.Where("id = ?", id).Delete(&models.User{}).Error
	if err != nil {
		common.LogError("Delete", "Error when trying to delete user, error is =>", err)
		return err
	}
	return nil
}

// UpdatePhoto use for update a picture
func (call *UserRepositoryImpl) UpdatePhoto(user *models.User) error {
	users := new(models.Photo)
	err := call.database.Table("user").Where(" id = ?", user.ID).First(&users).Update(user).Error
	if err != nil {
		common.LogError("UpdatePhoto", "Error when trying to update photo profile, error is =>", err)
		return err
	}
	return nil
}
