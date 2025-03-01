package repositories

import (
	"fmt"
	"mini-project-evermos/models/entities"

	"gorm.io/gorm"
)

// Contract
type AddressRepository interface {
	FindByUserId(id uint) ([]entities.Address, error)
	FindById(id uint) (entities.Address, error)
	Insert(address entities.Address) (bool, error)
	Update(id uint, address entities.Address) (bool, error)
	Destroy(id uint) (bool, error)
	FindByCondition(condition map[string]interface{}) (entities.Address, error)
	FindAll() ([]entities.Address, error)
}

type addressRepositoryImpl struct {
	database *gorm.DB
}

func NewAddressRepository(database *gorm.DB) AddressRepository {
	return &addressRepositoryImpl{database}
}

func (repository *addressRepositoryImpl) FindByUserId(id uint) ([]entities.Address, error) {
	var addresses []entities.Address

	// Debug print
	fmt.Printf("Fetching addresses for user ID: %d\n", id)

	// Try both tables since we're using both 'user' and 'users' tables
	result := repository.database.Debug().
		Table("alamat").
		Where("id_user = ? AND deleted_at IS NULL", id).
		Order("id desc").
		Find(&addresses)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch addresses: %v", result.Error)
	}

	// Debug print
	fmt.Printf("Found %d addresses for user ID %d\n", len(addresses), id)
	for _, addr := range addresses {
		fmt.Printf("Address ID: %d, User ID: %d\n", addr.ID, addr.IDUser)
	}

	return addresses, nil
}

func (repository *addressRepositoryImpl) FindById(id uint) (entities.Address, error) {
	var address entities.Address

	fmt.Printf("Looking for address with ID: %d in table: %s\n", id, entities.Address{}.TableName())

	result := repository.database.Debug().
		Table(entities.Address{}.TableName()).
		Where("id = ? AND deleted_at IS NULL", id).
		Order("id desc").
		First(&address)

	if result.Error != nil {
		fmt.Printf("Database error: %v\n", result.Error)
		return entities.Address{}, fmt.Errorf("address with ID %d not found", id)
	}

	fmt.Printf("Found address: %+v\n", address)
	return address, nil
}

func (repository *addressRepositoryImpl) Insert(address entities.Address) (bool, error) {
	tx := repository.database.Begin()

	fmt.Printf("Attempting to insert address with user ID: %d\n", address.IDUser)

	// First try to find user in 'user' table
	var userCount int64
	err := tx.Table("user").Where("id = ?", address.IDUser).Count(&userCount).Error
	if err != nil {
		tx.Rollback()
		return false, fmt.Errorf("error checking user table: %v", err)
	}

	// If user doesn't exist in 'user' table but exists in 'users' table, copy them over
	if userCount == 0 {
		var usersRecord entities.User
		if err := tx.Table("users").Where("id = ?", address.IDUser).First(&usersRecord).Error; err != nil {
			tx.Rollback()
			return false, fmt.Errorf("user with ID %d not found in either table", address.IDUser)
		}

		// Copy user to 'user' table
		if err := tx.Exec(`
            INSERT INTO user (
                id, nama, kata_sandi, notelp, tanggal_lahir, 
                jenis_kelamin, tentang, pekerjaan, email, 
                id_provinsi, id_kota, is_admin, created_at, updated_at
            ) 
            SELECT 
                id, nama, kata_sandi, notelp, tanggal_lahir,
                jenis_kelamin, tentang, pekerjaan, email,
                id_provinsi, id_kota, is_admin, created_at, updated_at
            FROM users WHERE id = ?
        `, address.IDUser).Error; err != nil {
			tx.Rollback()
			return false, fmt.Errorf("failed to copy user data: %v", err)
		}
	}

	// Now insert the address
	if err := tx.Table("alamat").Create(&address).Error; err != nil {
		tx.Rollback()
		fmt.Printf("Failed to insert address: %v\n", err)
		return false, fmt.Errorf("failed to create address: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return false, fmt.Errorf("failed to commit transaction: %v", err)
	}

	fmt.Printf("Successfully created address for user ID: %d\n", address.IDUser)
	return true, nil
}

func (repository *addressRepositoryImpl) Update(id uint, address entities.Address) (bool, error) {
	err := repository.database.Model(&address).Where("id = ?", id).Updates(address).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func (repository *addressRepositoryImpl) Destroy(id uint) (bool, error) {
	var address entities.Address
	err := repository.database.Where("id = ?", id).Delete(&address).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func (repository *addressRepositoryImpl) FindByCondition(condition map[string]interface{}) (entities.Address, error) {
	var address entities.Address
	result := repository.database.Where(condition).Order("created_at DESC").First(&address)
	if result.Error != nil {
		return entities.Address{}, result.Error
	}

	return address, nil
}

func (repository *addressRepositoryImpl) FindAll() ([]entities.Address, error) {
	var addresses []entities.Address

	result := repository.database.Debug().
		Table("alamat").
		Where("deleted_at IS NULL").
		Order("id desc").
		Find(&addresses)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch addresses: %v", result.Error)
	}

	fmt.Printf("Found %d total addresses\n", len(addresses))
	return addresses, nil
}
