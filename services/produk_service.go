package services

import (
	"errors"
	"fmt"
	"mini-project-evermos/models"
	"mini-project-evermos/models/entities"
	"mini-project-evermos/models/responder" // Updated import path
	"mini-project-evermos/repositories"
	"time"
)

type ProductService interface {
	FindAllPagination(limit int, page int, keyword string) (models.Pagination, error)
	FindById(id uint) (models.ProductResponse, error)
	Create(input models.ProductRequest, userId uint) (models.ProductResponse, error)
	Update(id uint, input models.ProductRequest, userId uint) (models.ProductResponse, error)
	Delete(id uint, userId uint) (models.ProductResponse, error) // Changed return type
	FindByCategory(categoryID string) ([]models.ProductResponse, error)
	SearchProducts(query string) ([]models.ProductResponse, error)
	GetRelatedProducts(id uint) ([]models.ProductResponse, error)
	SaveProductPhoto(photo entities.FotoProduk) (models.FotoProdukResponse, error) // Changed return type
	GetProductPhoto(id uint) (models.FotoProdukResponse, error)                    // Changed return type
}

type productServiceImpl struct {
	repository   repositories.ProductRepository
	storeRepo    repositories.StoreRepository
	categoryRepo repositories.CategoryRepository
}

func NewProductService(
	repository repositories.ProductRepository,
	storeRepo repositories.StoreRepository,
	categoryRepo repositories.CategoryRepository,
) ProductService {
	return &productServiceImpl{
		repository:   repository,
		storeRepo:    storeRepo,
		categoryRepo: categoryRepo,
	}
}

func (service *productServiceImpl) FindAllPagination(limit int, page int, keyword string) (models.Pagination, error) {
	request := responder.Pagination{}
	request.Limit = limit
	request.Page = page
	request.Keyword = keyword

	response, err := service.repository.FindAllPagination(request)
	if err != nil {
		return models.Pagination{}, err
	}

	return models.Pagination{
		Limit:      response.Limit,
		Page:       response.Page,
		TotalRows:  response.TotalRows,
		TotalPages: response.TotalPages,
		Rows:       response.Rows,
		Keyword:    response.Keyword,
	}, nil
}

func (service *productServiceImpl) FindById(id uint) (models.ProductResponse, error) {
	return service.repository.FindById(id)
}

func (service *productServiceImpl) Create(input models.ProductRequest, userId uint) (models.ProductResponse, error) {
	// Instead of finding store by user ID, verify the store exists
	store, _, err := service.storeRepo.FindById(input.StoreID) // Add _ to handle photos return value
	if err != nil {
		return models.ProductResponse{}, err
	}

	category, err := service.categoryRepo.FindById(input.CategoryID)
	if err != nil {
		return models.ProductResponse{}, err
	}
	if category.ID == 0 {
		return models.ProductResponse{}, errors.New("category with ID " + fmt.Sprint(input.CategoryID) + " not found")
	}

	// Use the store ID from the request
	input.StoreID = store.ID

	return service.repository.Insert(input)
}

func (service *productServiceImpl) Update(id uint, request models.ProductRequest, userId uint) (models.ProductResponse, error) {
	// Remove store lookup by userId and use the store ID from the request
	store, _, err := service.storeRepo.FindById(request.StoreID)
	if err != nil {
		return models.ProductResponse{}, err
	}
	if store.ID == 0 {
		return models.ProductResponse{}, errors.New("store not found")
	}

	// Verify category exists
	category, err := service.categoryRepo.FindById(request.CategoryID)
	if err != nil {
		return models.ProductResponse{}, err
	}
	if category.ID == 0 {
		return models.ProductResponse{}, errors.New("category with ID " + fmt.Sprint(request.CategoryID) + " not found")
	}

	// Add store ID to request
	request.StoreID = store.ID

	// Update product
	response, err := service.repository.Update(id, request)
	if err != nil {
		return models.ProductResponse{}, err
	}

	return response, nil
}

func (service *productServiceImpl) Delete(id uint, userId uint) (models.ProductResponse, error) {
	// Get product data before deletion
	product, err := service.repository.FindById(id)
	if err != nil {
		return models.ProductResponse{}, err
	}

	// Attempt to delete the product
	success, err := service.repository.Destroy(id)
	if err != nil {
		return models.ProductResponse{}, err
	}

	if !success {
		return models.ProductResponse{}, errors.New("failed to delete product")
	}

	return product, nil
}

func (service *productServiceImpl) FindByCategory(categoryID string) ([]models.ProductResponse, error) {
	return service.repository.FindByCategory(categoryID)
}

func (service *productServiceImpl) SearchProducts(query string) ([]models.ProductResponse, error) {
	return service.repository.SearchProducts(query)
}

func (service *productServiceImpl) GetRelatedProducts(id uint) ([]models.ProductResponse, error) {
	return service.repository.FindRelatedProducts(id)
}

func (service *productServiceImpl) SaveProductPhoto(photo entities.FotoProduk) (models.FotoProdukResponse, error) {
	// Ensure the PhotoURL is set from the request
	if photo.PhotoURL == "" {
		photo.PhotoURL = photo.URL // Fallback to URL if PhotoURL is not set
	}

	savedPhoto, err := service.repository.SaveProductPhoto(photo)
	if err != nil {
		return models.FotoProdukResponse{}, err
	}
	return mapSavedPhotoToResponse(savedPhoto), nil
}

func mapSavedPhotoToResponse(savedPhoto entities.FotoProduk) models.FotoProdukResponse {
	var createdAt, updatedAt time.Time
	if savedPhoto.CreatedAt != nil {
		createdAt = *savedPhoto.CreatedAt
	}
	if savedPhoto.UpdatedAt != nil {
		updatedAt = *savedPhoto.UpdatedAt
	}

	return models.FotoProdukResponse{
		ID:        savedPhoto.ID,
		URL:       savedPhoto.URL,
		Photo:     savedPhoto.Photo,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func (service *productServiceImpl) GetProductPhoto(id uint) (models.FotoProdukResponse, error) {
	photo, err := service.repository.GetProductPhoto(id)
	if err != nil {
		return models.FotoProdukResponse{}, err
	}

	return mapSavedPhotoToResponse(photo), nil
}

func mapProductPhotoToResponse(photo entities.FotoProduk) models.FotoProdukResponse {
	var createdAt, updatedAt time.Time
	if photo.CreatedAt != nil {
		createdAt = *photo.CreatedAt
	}
	if photo.UpdatedAt != nil {
		updatedAt = *photo.UpdatedAt
	}

	return models.FotoProdukResponse{
		ID:        photo.ID,
		URL:       photo.PhotoURL,
		Photo:     photo.Photo,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
