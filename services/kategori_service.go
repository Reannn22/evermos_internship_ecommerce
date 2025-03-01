package services

import (
	"mini-project-evermos/models"
	"mini-project-evermos/models/entities"
	"mini-project-evermos/repositories"
)

// Contract
type CategoryService interface {
	GetAll() ([]models.CategoryResponse, error)
	GetById(id uint) (models.CategoryResponse, error)
	Create(payload models.CategoryRequest) (models.CategoryResponse, error)
	Edit(id uint, payload models.CategoryRequest) (models.CategoryResponse, error)
	Delete(id uint) (models.CategoryResponse, error)
}

type categoryServiceImpl struct {
	repository repositories.CategoryRepository
}

func NewCategoryService(categoryRepository *repositories.CategoryRepository) CategoryService {
	return &categoryServiceImpl{
		repository: *categoryRepository,
	}
}

func (service *categoryServiceImpl) GetAll() ([]models.CategoryResponse, error) {
	categories, err := service.repository.FindAll()

	if err != nil {
		return nil, err
	}

	// mapping response
	responses := []models.CategoryResponse{}

	for _, category := range categories {
		response := models.CategoryResponse{
			ID:           category.ID,
			NamaCategory: category.NamaCategory,
			CreatedAt:    category.CreatedAt,
			UpdatedAt:    category.UpdatedAt,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (service *categoryServiceImpl) GetById(id uint) (models.CategoryResponse, error) {
	category, err := service.repository.FindById(id)

	if err != nil {
		return models.CategoryResponse{}, err
	}

	var response = models.CategoryResponse{
		ID:           category.ID,
		NamaCategory: category.NamaCategory,
		CreatedAt:    category.CreatedAt,
		UpdatedAt:    category.UpdatedAt,
	}

	return response, nil
}

func (service *categoryServiceImpl) Create(payload models.CategoryRequest) (models.CategoryResponse, error) {
	category := entities.Category{}
	category.NamaCategory = payload.NamaCategory

	result, err := service.repository.Insert(category)
	if err != nil {
		return models.CategoryResponse{}, err
	}

	response := models.CategoryResponse{
		ID:           result.ID,
		NamaCategory: result.NamaCategory,
		CreatedAt:    result.CreatedAt,
		UpdatedAt:    result.UpdatedAt,
	}
	return response, nil
}

func (service *categoryServiceImpl) Edit(id uint, payload models.CategoryRequest) (models.CategoryResponse, error) {
	//check
	_, err := service.repository.FindById(id)
	if err != nil {
		return models.CategoryResponse{}, err
	}

	category := entities.Category{}
	category.NamaCategory = payload.NamaCategory

	result, err := service.repository.Update(id, category)
	if err != nil {
		return models.CategoryResponse{}, err
	}

	response := models.CategoryResponse{
		ID:           result.ID,
		NamaCategory: result.NamaCategory,
		CreatedAt:    result.CreatedAt,
		UpdatedAt:    result.UpdatedAt,
	}
	return response, nil
}

func (service *categoryServiceImpl) Delete(id uint) (models.CategoryResponse, error) {
	//check first to get the data
	category, err := service.repository.FindById(id)
	if err != nil {
		return models.CategoryResponse{}, err
	}

	_, err = service.repository.Destroy(id)
	if err != nil {
		return models.CategoryResponse{}, err
	}

	response := models.CategoryResponse{
		ID:           category.ID,
		NamaCategory: category.NamaCategory,
		CreatedAt:    category.CreatedAt,
		UpdatedAt:    category.UpdatedAt,
	}
	return response, nil
}
