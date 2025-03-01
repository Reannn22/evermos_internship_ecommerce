package services

import (
	"mini-project-evermos/models"
	"mini-project-evermos/repositories"
)

type DiskonProdukService interface {
	ApplyDiscount(input models.DiskonProdukRequest) (models.DiskonProdukResponse, error)
	GetById(id uint) (models.DiskonProdukResponse, error)
	GetAll() ([]models.DiskonProdukResponse, error)
	UpdateDiscount(id uint, input models.DiskonProdukRequest) (models.DiskonProdukResponse, error)
	DeleteDiscount(id uint) (models.DiskonProdukResponse, error)
}

type diskonProdukServiceImpl struct {
	repository repositories.DiskonProdukRepository
}

func NewDiskonProdukService(repository repositories.DiskonProdukRepository) DiskonProdukService {
	return &diskonProdukServiceImpl{repository}
}

func (service *diskonProdukServiceImpl) ApplyDiscount(input models.DiskonProdukRequest) (models.DiskonProdukResponse, error) {
	return service.repository.ApplyDiscount(input.ProductID, input.HargaKonsumen)
}

func (service *diskonProdukServiceImpl) GetById(id uint) (models.DiskonProdukResponse, error) {
	return service.repository.GetById(id)
}

func (service *diskonProdukServiceImpl) GetAll() ([]models.DiskonProdukResponse, error) {
	return service.repository.GetAll()
}

func (service *diskonProdukServiceImpl) UpdateDiscount(id uint, input models.DiskonProdukRequest) (models.DiskonProdukResponse, error) {
	return service.repository.UpdateDiscount(id, input.ProductID, input.HargaKonsumen)
}

func (service *diskonProdukServiceImpl) DeleteDiscount(id uint) (models.DiskonProdukResponse, error) {
	return service.repository.DeleteDiscount(id)
}
