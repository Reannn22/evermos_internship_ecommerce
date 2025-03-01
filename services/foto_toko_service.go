package services

import (
	"mini-project-evermos/models"
	"mini-project-evermos/models/entities"
	"mini-project-evermos/repositories"
	"time"
)

type StorePhotoService struct {
	repository repositories.StorePhotoRepository
}

func NewStorePhotoService(repository repositories.StorePhotoRepository) *StorePhotoService {
	return &StorePhotoService{repository}
}

func (service *StorePhotoService) GetAll() ([]models.StorePhotoResponse, error) {
	photos, err := service.repository.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []models.StorePhotoResponse
	for _, photo := range photos {
		response := models.StorePhotoResponse{
			ID:        photo.ID,
			IdToko:    photo.IdToko,
			URL:       photo.URL,
			Photo:     photo.Photo,
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (service *StorePhotoService) GetById(id uint) (models.StorePhotoResponse, error) {
	photo, err := service.repository.FindById(id)
	if err != nil {
		return models.StorePhotoResponse{}, err
	}

	response := models.StorePhotoResponse{
		ID:        photo.ID,
		IdToko:    photo.IdToko,
		URL:       photo.URL,
		Photo:     photo.Photo,
		CreatedAt: photo.CreatedAt,
		UpdatedAt: photo.UpdatedAt,
	}

	return response, nil
}

func (service *StorePhotoService) Create(input models.StorePhotoRequest) (*models.StorePhotoResponse, error) {
	photo := entities.FotoToko{
		IdToko:    input.IdToko,
		URL:       input.URL,
		Photo:     input.Photo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	created, err := service.repository.Insert(photo)
	if err != nil {
		return nil, err
	}

	response := &models.StorePhotoResponse{
		ID:        created.ID,
		IdToko:    created.IdToko,
		URL:       created.URL,
		Photo:     created.Photo,
		CreatedAt: created.CreatedAt,
		UpdatedAt: created.UpdatedAt,
	}

	return response, nil
}

func (service *StorePhotoService) Update(id uint, input models.StorePhotoRequest) (models.StorePhotoResponse, error) {
	// First get the existing photo to preserve its created_at timestamp
	existingPhoto, err := service.repository.FindById(id)
	if err != nil {
		return models.StorePhotoResponse{}, err
	}

	photo := entities.FotoToko{
		ID:        id,
		IdToko:    input.IdToko,
		URL:       input.URL,
		Photo:     input.Photo,
		CreatedAt: existingPhoto.CreatedAt, // Preserve original creation time
		UpdatedAt: time.Now(),
	}

	updated, err := service.repository.Update(id, photo)
	if err != nil {
		return models.StorePhotoResponse{}, err
	}

	response := models.StorePhotoResponse{
		ID:        updated.ID,
		URL:       updated.URL,
		Photo:     updated.Photo,
		CreatedAt: updated.CreatedAt, // Remove * since CreatedAt is already time.Time
		UpdatedAt: updated.UpdatedAt, // Remove * since UpdatedAt is already time.Time
	}

	return response, nil
}

// Update return type to include the deleted photo data
func (service *StorePhotoService) Delete(id uint) (models.StorePhotoResponse, error) {
	// First get the photo before deletion
	existingPhoto, err := service.repository.FindById(id)
	if err != nil {
		return models.StorePhotoResponse{}, err
	}

	// Delete the photo
	_, err = service.repository.Delete(id)
	if err != nil {
		return models.StorePhotoResponse{}, err
	}

	// Return the photo data that was retrieved before deletion
	return models.StorePhotoResponse{
		ID:        existingPhoto.ID,
		URL:       existingPhoto.URL,
		Photo:     existingPhoto.Photo,
		CreatedAt: existingPhoto.CreatedAt,
		UpdatedAt: existingPhoto.UpdatedAt,
	}, nil
}
