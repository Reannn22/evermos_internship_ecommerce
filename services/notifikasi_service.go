package services

import (
	"mini-project-evermos/models"
	"mini-project-evermos/models/entities"
	"mini-project-evermos/repositories"
)

type NotificationService interface {
	GetAll() ([]models.NotificationResponse, error)
	GetById(id uint) (models.NotificationResponse, error)
	Create(input models.NotificationRequest) (models.NotificationResponse, error)
	Update(id uint, input models.NotificationRequest) (models.NotificationResponse, error)
	Delete(id uint) (models.NotificationResponse, error)
}

type notificationServiceImpl struct {
	repository repositories.NotificationRepository
}

func NewNotificationService(repository repositories.NotificationRepository) NotificationService {
	return &notificationServiceImpl{repository}
}

func (s *notificationServiceImpl) GetAll() ([]models.NotificationResponse, error) {
	notifications, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []models.NotificationResponse
	for _, notification := range notifications {
		responses = append(responses, toNotificationResponse(notification))
	}
	return responses, nil
}

func (s *notificationServiceImpl) GetById(id uint) (models.NotificationResponse, error) {
	notification, err := s.repository.FindById(id)
	if err != nil {
		return models.NotificationResponse{}, err
	}
	return toNotificationResponse(notification), nil
}

func (s *notificationServiceImpl) Create(input models.NotificationRequest) (models.NotificationResponse, error) {
	notification := entities.Notification{
		Pesan: input.Pesan,
	}

	result, err := s.repository.Create(notification)
	if err != nil {
		return models.NotificationResponse{}, err
	}

	return toNotificationResponse(result), nil
}

func (s *notificationServiceImpl) Update(id uint, input models.NotificationRequest) (models.NotificationResponse, error) {
	notification := entities.Notification{
		Pesan: input.Pesan,
	}

	result, err := s.repository.Update(id, notification)
	if err != nil {
		return models.NotificationResponse{}, err
	}

	return toNotificationResponse(result), nil
}

func (s *notificationServiceImpl) Delete(id uint) (models.NotificationResponse, error) {
	// Get notification before deleting
	notification, err := s.repository.FindById(id)
	if err != nil {
		return models.NotificationResponse{}, err
	}

	// Delete the notification
	err = s.repository.Delete(id)
	if err != nil {
		return models.NotificationResponse{}, err
	}

	return toNotificationResponse(notification), nil
}

func toNotificationResponse(notification entities.Notification) models.NotificationResponse {
	return models.NotificationResponse{
		ID:        notification.ID,
		Pesan:     notification.Pesan,
		CreatedAt: notification.CreatedAt,
		UpdatedAt: notification.UpdatedAt,
	}
}
