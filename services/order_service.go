package services

import (
	"mini-project-evermos/models"
	"mini-project-evermos/models/entities"
	"mini-project-evermos/repositories"
)

type OrderService interface {
	GetAll() ([]models.OrderResponse, error)
	GetById(id uint) (models.OrderResponse, error)
	UpdateProductStatus(input models.OrderRequest) (models.OrderResponse, error)
	Update(id uint, input models.OrderRequest) (models.OrderResponse, error)
	Delete(id uint) (models.OrderResponse, error) // Change return type
}

type orderServiceImpl struct {
	repository    repositories.OrderRepository
	trxDetailRepo repositories.TransactionDetailRepository
}

func NewOrderService(repo repositories.OrderRepository, trxDetailRepo repositories.TransactionDetailRepository) OrderService {
	return &orderServiceImpl{
		repository:    repo,
		trxDetailRepo: trxDetailRepo,
	}
}

func (service *orderServiceImpl) GetAll() ([]models.OrderResponse, error) {
	orders, err := service.repository.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []models.OrderResponse
	for _, order := range orders {
		responses = append(responses, models.OrderResponse{
			ID:           order.ID,
			TrxDetailID:  order.TransactionDetailID,
			StatusProduk: order.StatusProduk,
			CreatedAt:    order.CreatedAt,
			UpdatedAt:    order.UpdatedAt,
		})
	}
	return responses, nil
}

func (service *orderServiceImpl) GetById(id uint) (models.OrderResponse, error) {
	order, err := service.repository.FindById(id)
	if err != nil {
		return models.OrderResponse{}, err
	}

	return models.OrderResponse{
		ID:           order.ID,
		TrxDetailID:  order.TransactionDetailID,
		StatusProduk: order.StatusProduk,
		CreatedAt:    order.CreatedAt,
		UpdatedAt:    order.UpdatedAt,
	}, nil
}

func (service *orderServiceImpl) UpdateProductStatus(input models.OrderRequest) (models.OrderResponse, error) {
	// Verify transaction detail exists
	_, err := service.trxDetailRepo.FindById(input.TransactionDetailID)
	if err != nil {
		return models.OrderResponse{}, err
	}

	statusProduk := "shipping..."
	if input.ProductStatus == "true" {
		statusProduk = "shipped"
	}

	// Update transaction detail status
	err = service.trxDetailRepo.UpdateProductStatus(input.TransactionDetailID, statusProduk)
	if err != nil {
		return models.OrderResponse{}, err
	}

	order := entities.Order{
		TransactionDetailID: input.TransactionDetailID,
		StatusProduk:        statusProduk,
	}

	created, err := service.repository.Create(order)
	if err != nil {
		return models.OrderResponse{}, err
	}

	return models.OrderResponse{
		ID:           created.ID,
		TrxDetailID:  created.TransactionDetailID,
		StatusProduk: created.StatusProduk,
		CreatedAt:    created.CreatedAt,
		UpdatedAt:    created.UpdatedAt,
	}, nil
}

func (service *orderServiceImpl) Update(id uint, input models.OrderRequest) (models.OrderResponse, error) {
	existing, err := service.repository.FindById(id)
	if err != nil {
		return models.OrderResponse{}, err
	}

	statusProduk := "shipping..."
	if input.ProductStatus == "true" {
		statusProduk = "shipped"
	}

	// Update transaction detail status
	err = service.trxDetailRepo.UpdateProductStatus(input.TransactionDetailID, statusProduk)
	if err != nil {
		return models.OrderResponse{}, err
	}

	existing.TransactionDetailID = input.TransactionDetailID
	existing.StatusProduk = statusProduk

	updated, err := service.repository.Update(existing)
	if err != nil {
		return models.OrderResponse{}, err
	}

	return models.OrderResponse{
		ID:           updated.ID,
		TrxDetailID:  updated.TransactionDetailID,
		StatusProduk: updated.StatusProduk,
		CreatedAt:    updated.CreatedAt,
		UpdatedAt:    updated.UpdatedAt,
	}, nil
}

func (service *orderServiceImpl) Delete(id uint) (models.OrderResponse, error) {
	// Get order before deleting
	order, err := service.repository.FindById(id)
	if err != nil {
		return models.OrderResponse{}, err
	}

	// Delete the order
	err = service.repository.Delete(id)
	if err != nil {
		return models.OrderResponse{}, err
	}

	// Return the deleted order data
	return models.OrderResponse{
		ID:           order.ID,
		TrxDetailID:  order.TransactionDetailID,
		StatusProduk: order.StatusProduk,
		CreatedAt:    order.CreatedAt,
		UpdatedAt:    order.UpdatedAt,
	}, nil
}
