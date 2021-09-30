package services

import (
	"go-fx-test/lib"
	"go-fx-test/models"
	"go-fx-test/repository"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(
	repository repository.UserRepository,
) UserService {
	return UserService{
		repository: repository,
	}
}

func (s UserService) GetAllUser() (users []models.User, err error) {
	return users, s.repository.Find(&users).Error
}

func (s UserService) Create(user models.User) error {
	return s.repository.Create(&user).Error
}

func (s UserService) GetOneUser(userID lib.BinaryUUID) (user models.User, err error) {
	return user, s.repository.First(&user, "id = ?", userID).Error
}

func (s UserService) UpdateUser(user *models.User) (*models.User, error) {
	if err := s.repository.Model(&models.User{}).Where("id = ?", user.ID).Updates(map[string]interface{}{
		"fname": user.Fname,
		"lname": user.Lname,
		"admin": user.Admin,
	}).Error; err != nil {
		return nil, err
	}
	return user, nil
}
