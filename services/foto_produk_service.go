package services

import (
	"fmt"
	"mini-project-evermos/models"
	"mini-project-evermos/models/entities"
	"mini-project-evermos/repositories"
	"time"
)

type FotoProdukService interface {
	Create(request models.FotoProdukRequest) (models.FotoProdukResponse, error)
	FindById(id uint) (models.FotoProdukResponse, error)
	Update(id uint, request models.FotoProdukRequest) (models.FotoProdukResponse, error)
	Delete(id uint) (models.FotoProdukResponse, error)
	FindAll() ([]models.FotoProdukResponse, error)
	FindByProductId(productId uint) ([]models.FotoProdukResponse, error)
	Upload(request models.FotoProdukRequest) (models.FileUploadResponse, error)
}

type fotoProdukServiceImpl struct {
	repository        repositories.FotoProdukRepository
	productRepository repositories.ProductRepository
}

func NewFotoProdukService(
	repository repositories.FotoProdukRepository,
	productRepository repositories.ProductRepository,
) FotoProdukService {
	return &fotoProdukServiceImpl{
		repository:        repository,
		productRepository: productRepository,
	}
}

func (service *fotoProdukServiceImpl) Create(request models.FotoProdukRequest) (models.FotoProdukResponse, error) {
	// First validate that the product exists
	exists, err := service.repository.CheckProductExists(request.ProductID)
	if err != nil {
		return models.FotoProdukResponse{}, err
	}
	if !exists {
		return models.FotoProdukResponse{}, fmt.Errorf("product with ID %d does not exist", request.ProductID)
	}

	// Then create the photo
	foto, err := service.repository.Create(request)
	if err != nil {
		return models.FotoProdukResponse{}, err
	}

	return mapToFotoProdukResponse(foto), nil
}

func (service *fotoProdukServiceImpl) FindById(id uint) (models.FotoProdukResponse, error) {
	foto, err := service.repository.FindById(id)
	if err != nil {
		return models.FotoProdukResponse{}, err
	}

	return mapToFotoProdukResponse(foto), nil
}

func (service *fotoProdukServiceImpl) Update(id uint, request models.FotoProdukRequest) (models.FotoProdukResponse, error) {
	foto, err := service.repository.Update(id, request)
	if err != nil {
		return models.FotoProdukResponse{}, err
	}

	return mapToFotoProdukResponse(foto), nil
}

func (service *fotoProdukServiceImpl) Delete(id uint) (models.FotoProdukResponse, error) {
	// Get the photo before deleting
	foto, err := service.repository.FindById(id)
	if err != nil {
		return models.FotoProdukResponse{}, err
	}

	// Delete the photo
	err = service.repository.Delete(id)
	if err != nil {
		return models.FotoProdukResponse{}, err
	}

	return mapToFotoProdukResponse(foto), nil
}

func (service *fotoProdukServiceImpl) FindAll() ([]models.FotoProdukResponse, error) {
	photos, err := service.repository.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []models.FotoProdukResponse
	for _, foto := range photos {
		responses = append(responses, mapToFotoProdukResponse(foto))
	}
	return responses, nil
}

func (service *fotoProdukServiceImpl) FindByProductId(productId uint) ([]models.FotoProdukResponse, error) {
	photos, err := service.repository.FindByProductId(productId)
	if err != nil {
		return nil, err
	}

	var responses []models.FotoProdukResponse
	for _, foto := range photos {
		responses = append(responses, mapToFotoProdukResponse(foto))
	}
	return responses, nil
}

// Update the photo mapping function
func mapToFotoProdukResponse(foto entities.FotoProduk) models.FotoProdukResponse {
	var createdAt, updatedAt time.Time
	if foto.CreatedAt != nil {
		createdAt = *foto.CreatedAt
	}
	if foto.UpdatedAt != nil {
		updatedAt = *foto.UpdatedAt
	}

	return models.FotoProdukResponse{
		ID:        foto.ID,
		URL:       foto.URL,
		Photo:     foto.Photo,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func (service *fotoProdukServiceImpl) Upload(request models.FotoProdukRequest) (models.FileUploadResponse, error) {
	// Check if the product exists
	_, err := service.productRepository.FindById(request.ProductID)
	if err != nil {
		return models.FileUploadResponse{}, err
	}

	// Create new photo record using only available fields
	result, err := service.repository.Create(models.FotoProdukRequest{
		ProductID: request.ProductID,
		PhotoURL:  request.PhotoURL,
		File:      request.File,
	})
	if err != nil {
		return models.FileUploadResponse{}, err
	}

	return models.FileUploadResponse{
		URL:      result.URL,
		PhotoID:  result.ID,
		Filename: request.File.Filename,
	}, nil
}
