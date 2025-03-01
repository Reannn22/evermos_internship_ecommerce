package repositories

import (
	"mini-project-evermos/models/entities"

	"gorm.io/gorm"
)

type OrderRepository interface {
	FindAll() ([]entities.Order, error)
	FindById(id uint) (entities.Order, error)
	Create(order entities.Order) (entities.Order, error)
	Update(order entities.Order) (entities.Order, error)
	Delete(id uint) error
}

type orderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepositoryImpl{db: db}
}

func (repo *orderRepositoryImpl) FindAll() ([]entities.Order, error) {
	var orders []entities.Order
	err := repo.db.Order("id desc").Find(&orders).Error // Added Order("id desc")
	return orders, err
}

func (repo *orderRepositoryImpl) FindById(id uint) (entities.Order, error) {
	var order entities.Order
	err := repo.db.First(&order, id).Error
	return order, err
}

func (repo *orderRepositoryImpl) Create(order entities.Order) (entities.Order, error) {
	err := repo.db.Create(&order).Error
	return order, err
}

func (repo *orderRepositoryImpl) Update(order entities.Order) (entities.Order, error) {
	err := repo.db.Save(&order).Error
	return order, err
}

func (repo *orderRepositoryImpl) Delete(id uint) error {
	return repo.db.Delete(&entities.Order{}, id).Error
}
