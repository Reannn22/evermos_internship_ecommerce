package repositories

import (
	"fmt"
	"mini-project-evermos/models/entities"

	"gorm.io/gorm"
)

// Contract
type UserRepository interface {
	FindByNoTelp(no_telp string) (entities.User, error)
	FindById(id uint) (entities.User, error)
	Update(id uint, user entities.User) (bool, error)
	FindByEmail(email string) (entities.User, error)
	Delete(id uint) error
	FindLastUser() (entities.User, error)
	FindAll() ([]entities.User, error)
}

type userRepositoryImpl struct {
	database *gorm.DB
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepositoryImpl{database}
}

func (repository *userRepositoryImpl) FindByNoTelp(no_telp string) (entities.User, error) {
	var user entities.User
	err := repository.database.Where("notelp = ?", no_telp).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *userRepositoryImpl) FindById(id uint) (entities.User, error) {
	var user entities.User
	fmt.Printf("Looking for user with ID: %d\n", id)

	result := repository.database.Debug().
		Table("users"). // Change from 'user' to 'users'
		Where("id = ? AND deleted_at IS NULL", id).
		First(&user)

	if result.Error != nil {
		// Try the 'user' table if 'users' fails
		result = repository.database.Debug().
			Table("user").
			Where("id = ? AND deleted_at IS NULL", id).
			First(&user)

		if result.Error != nil {
			fmt.Printf("Error finding user: %v\n", result.Error)
			return user, result.Error
		}
	}

	fmt.Printf("Found user: %+v\n", user)
	return user, nil
}

func (repository *userRepositoryImpl) Update(id uint, user entities.User) (bool, error) {
	err := repository.database.Model(&user).Where("id = ?", id).Updates(user).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func (repository *userRepositoryImpl) FindByEmail(email string) (entities.User, error) {
	var user entities.User

	err := repository.database.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *userRepositoryImpl) Delete(id uint) error {
	var user entities.User
	err := repository.database.Where("id = ?", id).Delete(&user).Error
	return err
}

func (repository *userRepositoryImpl) FindLastUser() (entities.User, error) {
	var user entities.User
	err := repository.database.Last(&user).Error
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (repository *userRepositoryImpl) FindAll() ([]entities.User, error) {
	var users []entities.User
	fmt.Println("Fetching all users from database")

	// Modified query to sort by ID in descending order
	err := repository.database.
		Unscoped().
		Where("deleted_at IS NULL").
		Order("id desc"). // Changed from "id asc" to "id desc"
		Find(&users).Error

	fmt.Printf("Found %d users\n", len(users))

	if err != nil {
		fmt.Printf("Error fetching users: %v\n", err)
		return nil, err
	}
	return users, nil
}
