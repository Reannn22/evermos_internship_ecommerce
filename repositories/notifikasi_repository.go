package repositories

import (
	"mini-project-evermos/models/entities"
	"time"

	"gorm.io/gorm"
)

type NotificationRepository interface {
	FindAll() ([]entities.Notification, error)
	FindById(id uint) (entities.Notification, error)
	Create(notification entities.Notification) (entities.Notification, error)
	Update(id uint, notification entities.Notification) (entities.Notification, error)
	Delete(id uint) error
}

type notificationRepositoryImpl struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepositoryImpl{db}
}

func (r *notificationRepositoryImpl) FindAll() ([]entities.Notification, error) {
	var notifications []entities.Notification
	err := r.db.Order("created_at desc").Find(&notifications).Error
	return notifications, err
}

func (r *notificationRepositoryImpl) FindById(id uint) (entities.Notification, error) {
	var notification entities.Notification
	err := r.db.First(&notification, id).Error
	return notification, err
}

func (r *notificationRepositoryImpl) Create(notification entities.Notification) (entities.Notification, error) {
	now := time.Now()
	notification.CreatedAt = &now
	notification.UpdatedAt = &now
	err := r.db.Create(&notification).Error
	return notification, err
}

func (r *notificationRepositoryImpl) Update(id uint, notification entities.Notification) (entities.Notification, error) {
	now := time.Now()
	notification.UpdatedAt = &now
	err := r.db.Model(&entities.Notification{}).Where("id = ?", id).Updates(notification).Error
	if err != nil {
		return entities.Notification{}, err
	}
	return r.FindById(id)
}

func (r *notificationRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&entities.Notification{}, id).Error
}
