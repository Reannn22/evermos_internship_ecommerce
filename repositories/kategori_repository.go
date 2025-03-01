package repositories

import (
	"mini-project-evermos/models/entities"

	"gorm.io/gorm"
)

// Contract
type CategoryRepository interface {
	FindAll() ([]entities.Category, error)
	FindById(id uint) (entities.Category, error)
	Insert(category entities.Category) (entities.Category, error)
	Update(id uint, category entities.Category) (entities.Category, error)
	Destroy(id uint) (bool, error)
}

type categoryRepositoryImpl struct {
	database *gorm.DB
}

func NewCategoryRepository(database *gorm.DB) CategoryRepository {
	return &categoryRepositoryImpl{database}
}

func (repository *categoryRepositoryImpl) FindAll() ([]entities.Category, error) {
	var categories []entities.Category
	err := repository.database.Order("id desc").Find(&categories).Error
	if err != nil {
		return categories, err
	}
	return categories, nil
}

func (repository *categoryRepositoryImpl) FindById(id uint) (entities.Category, error) {
	var category entities.Category
	err := repository.database.Where("id = ?", id).Order("id desc").First(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (repository *categoryRepositoryImpl) Insert(category entities.Category) (entities.Category, error) {
	err := repository.database.Create(&category).Error
	return category, err
}

func (repository *categoryRepositoryImpl) Update(id uint, category entities.Category) (entities.Category, error) {
	category.ID = id
	err := repository.database.Save(&category).Error
	return category, err
}

func (repository *categoryRepositoryImpl) Destroy(id uint) (bool, error) {
	var category entities.Category
	err := repository.database.Where("id = ?", id).Delete(&category).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
