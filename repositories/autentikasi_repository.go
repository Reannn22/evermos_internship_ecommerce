package repositories

import (
	"fmt"
	"mini-project-evermos/models/entities"

	"gorm.io/gorm"
)

// Contract
type AuthRepository interface {
	Register(user entities.User) (entities.User, error)
}

type authRepositoryImpl struct {
	database *gorm.DB
}

func NewAuthRepository(database *gorm.DB) AuthRepository {
	return &authRepositoryImpl{database}
}

func (repository *authRepositoryImpl) Register(user entities.User) (entities.User, error) {
	// Check for existing user with same email or phone
	var existingUser entities.User
	result := repository.database.Where("notelp = ? OR email = ?", user.Notelp, user.Email).First(&existingUser)

	if result.Error == nil {
		// If user exists, return error
		return entities.User{}, fmt.Errorf("user with email %s or phone %s already exists", user.Email, user.Notelp)
	}

	// If no existing user found, create new user
	err := repository.database.Create(&user).Error
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}
