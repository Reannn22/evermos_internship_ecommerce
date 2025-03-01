package services

import (
	"mini-project-evermos/models"
	"mini-project-evermos/models/entities"
	"mini-project-evermos/repositories"
	"time" // Add this import
)

type TransactionDetailService interface {
	GetAll() ([]models.TransactionDetailResponse, error)
	GetById(id uint) (models.TransactionDetailResponse, error)
	GetByTrxId(trxId uint) ([]models.TransactionDetailResponse, error)
	Create(input models.TransactionDetailProcess) (models.TransactionDetailResponse, error)
	Update(id uint, input models.TransactionDetailProcess) (models.TransactionDetailResponse, error)
	Delete(id uint) (models.TransactionDetailResponse, error)
}

type transactionDetailServiceImpl struct {
	repository repositories.TransactionDetailRepository
}

func NewTransactionDetailService(repo repositories.TransactionDetailRepository) TransactionDetailService {
	return &transactionDetailServiceImpl{
		repository: repo,
	}
}

func (service *transactionDetailServiceImpl) GetAll() ([]models.TransactionDetailResponse, error) {
	details, err := service.repository.FindAll()
	if err != nil {
		return nil, err
	}

	var detailResponses []models.TransactionDetailResponse
	for _, detail := range details {
		// Handle nil CreatedAt and UpdatedAt
		createdAt := time.Now()
		updatedAt := time.Now()
		if detail.CreatedAt != nil {
			createdAt = *detail.CreatedAt
		}
		if detail.UpdatedAt != nil {
			updatedAt = *detail.UpdatedAt
		}

		detailResponse := models.TransactionDetailResponse{
			ID:            detail.ID,
			IDTransaksi:   detail.IDTrx,
			IDLogProduk:   detail.IDLogProduk,
			IDToko:        detail.IDToko,
			Kuantitas:     detail.Kuantitas,
			HargaTotal:    detail.HargaTotal,
			ProductStatus: "shipping...", // Add this line
			Store:         mapTransactionStoreToResponse(detail.Store),
			Produk:        mapToSimpleProductResponse(detail.ProductLog.Product),
			CreatedAt:     createdAt,
			UpdatedAt:     updatedAt,
		}
		detailResponses = append(detailResponses, detailResponse)
	}

	return detailResponses, nil
}

func (service *transactionDetailServiceImpl) GetById(id uint) (models.TransactionDetailResponse, error) {
	detail, err := service.repository.FindById(id)
	if err != nil {
		return models.TransactionDetailResponse{}, err
	}

	response := models.TransactionDetailResponse{
		ID:            detail.ID,
		IDTransaksi:   detail.IDTrx,
		IDLogProduk:   detail.IDLogProduk,
		IDToko:        detail.IDToko,
		Kuantitas:     detail.Kuantitas,
		HargaTotal:    detail.HargaTotal,
		ProductStatus: detail.ProductStatus, // Add this line
		Store:         mapTransactionStoreToResponse(detail.Store),
		Produk:        mapToSimpleProductResponse(detail.ProductLog.Product),
		CreatedAt:     *detail.CreatedAt,
		UpdatedAt:     *detail.UpdatedAt,
	}

	return response, nil
}

func mapToSimpleProductResponse(product entities.Product) models.SimpleProductResponse {
	// Map product photos
	var fotoProdukResponses []models.FotoProdukResponse
	for _, foto := range product.FotoProduk {
		var createdAt, updatedAt time.Time
		if foto.CreatedAt != nil {
			createdAt = *foto.CreatedAt
		}
		if foto.UpdatedAt != nil {
			updatedAt = *foto.UpdatedAt
		}

		fotoProdukResponses = append(fotoProdukResponses, models.FotoProdukResponse{
			ID:        foto.ID,
			Photo:     foto.Photo,
			URL:       foto.PhotoURL,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}

	// Map reviews
	var reviewResponses []models.SimpleProductReviewResponse
	for _, review := range product.Reviews {
		reviewResponses = append(reviewResponses, models.SimpleProductReviewResponse{
			ID:        uint(review.ID),
			IDToko:    uint(review.IDToko),
			IDProduk:  uint(review.IDProduk),
			Ulasan:    review.Ulasan,
			Rating:    review.Rating,
			CreatedAt: review.CreatedAt,
			UpdatedAt: review.UpdatedAt,
		})
	}

	// Map promos
	var promoResponses []models.ProductPromoResponse
	for _, promo := range product.Promos {
		promoResponses = append(promoResponses, models.ProductPromoResponse{
			ID:        uint(promo.ID),
			IDToko:    uint(promo.IDToko),
			IDProduk:  uint(promo.IDProduk),
			Promo:     promo.Promo,
			CreatedAt: promo.CreatedAt,
			UpdatedAt: promo.UpdatedAt,
		})
	}

	return models.SimpleProductResponse{
		ID:            product.ID,
		NamaProduk:    product.NamaProduk,
		Slug:          product.Slug,
		HargaReseller: product.HargaReseller,
		HargaKonsumen: product.HargaKonsumen,
		Stok:          product.Stok,
		Deskripsi:     *product.Deskripsi,
		FotoProduk:    fotoProdukResponses, // Add this field
		Reviews:       reviewResponses,     // Update this line
		Promos:        promoResponses,      // Update this line
		CreatedAt:     product.CreatedAt,
		UpdatedAt:     product.UpdatedAt,
	}
}

func (service *transactionDetailServiceImpl) Create(input models.TransactionDetailProcess) (models.TransactionDetailResponse, error) {
	detail, err := service.repository.Create(input)
	if err != nil {
		return models.TransactionDetailResponse{}, err
	}

	// Get complete detail with relationships
	detailWithRelations, err := service.repository.FindById(detail.ID)
	if err != nil {
		return models.TransactionDetailResponse{}, err
	}

	response := models.TransactionDetailResponse{
		ID:            detailWithRelations.ID,
		IDTransaksi:   detailWithRelations.IDTrx,
		IDLogProduk:   detailWithRelations.IDLogProduk,
		IDToko:        detailWithRelations.IDToko,
		Kuantitas:     detailWithRelations.Kuantitas,
		HargaTotal:    detailWithRelations.HargaTotal,
		ProductStatus: "shipping...", // Add this line
		Store:         mapTransactionStoreToResponse(detailWithRelations.Store),
		Produk:        mapToSimpleProductResponse(detailWithRelations.ProductLog.Product),
		CreatedAt:     *detailWithRelations.CreatedAt,
		UpdatedAt:     *detailWithRelations.UpdatedAt,
	}

	return response, nil
}

func (service *transactionDetailServiceImpl) Update(id uint, input models.TransactionDetailProcess) (models.TransactionDetailResponse, error) {
	// First check if detail exists
	existing, err := service.repository.FindById(id)
	if err != nil {
		return models.TransactionDetailResponse{}, err
	}

	// Update fields
	existing.IDTrx = input.TrxID
	existing.IDLogProduk = input.LogProductID
	existing.IDToko = input.StoreID
	existing.Kuantitas = input.Kuantitas
	existing.HargaTotal = input.HargaTotal

	// Save updates
	_, err = service.repository.Update(existing)
	if err != nil {
		return models.TransactionDetailResponse{}, err
	}

	// Get updated data with all relationships
	updatedDetail, err := service.repository.FindById(id)
	if err != nil {
		return models.TransactionDetailResponse{}, err
	}

	// Map complete response including Store and Product data
	response := models.TransactionDetailResponse{
		ID:            updatedDetail.ID,
		IDTransaksi:   updatedDetail.IDTrx,
		IDLogProduk:   updatedDetail.IDLogProduk,
		IDToko:        updatedDetail.IDToko,
		Kuantitas:     updatedDetail.Kuantitas,
		HargaTotal:    updatedDetail.HargaTotal,
		ProductStatus: "shipping...", // Add this line
		Store:         mapTransactionStoreToResponse(updatedDetail.Store),
		Produk:        mapToSimpleProductResponse(updatedDetail.ProductLog.Product),
		CreatedAt:     *updatedDetail.CreatedAt,
		UpdatedAt:     *updatedDetail.UpdatedAt,
	}

	return response, nil
}

func (service *transactionDetailServiceImpl) GetByTrxId(trxId uint) ([]models.TransactionDetailResponse, error) {
	details, err := service.repository.FindByTrxId(trxId)
	if err != nil {
		return nil, err
	}

	var responses []models.TransactionDetailResponse
	for _, detail := range details {
		// Handle nil timestamps
		createdAt := time.Now()
		updatedAt := time.Now()
		if detail.CreatedAt != nil {
			createdAt = *detail.CreatedAt
		}
		if detail.UpdatedAt != nil {
			updatedAt = *detail.UpdatedAt
		}

		response := models.TransactionDetailResponse{
			ID:            detail.ID,
			IDTransaksi:   detail.IDTrx,
			IDLogProduk:   detail.IDLogProduk,
			IDToko:        detail.IDToko,
			Kuantitas:     detail.Kuantitas,
			HargaTotal:    detail.HargaTotal,
			ProductStatus: "shipping...", // Add this line
			Store:         mapTransactionStoreToResponse(detail.Store),
			Produk:        mapToSimpleProductResponse(detail.ProductLog.Product),
			CreatedAt:     createdAt,
			UpdatedAt:     updatedAt,
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (service *transactionDetailServiceImpl) Delete(id uint) (models.TransactionDetailResponse, error) {
	// Get the detail first for the response
	detail, err := service.repository.FindById(id)
	if err != nil {
		return models.TransactionDetailResponse{}, err
	}

	// Prepare the response before deletion
	response := models.TransactionDetailResponse{
		ID:            detail.ID,
		IDTransaksi:   detail.IDTrx,
		IDLogProduk:   detail.IDLogProduk,
		IDToko:        detail.IDToko,
		Kuantitas:     detail.Kuantitas,
		HargaTotal:    detail.HargaTotal,
		ProductStatus: detail.ProductStatus, // Add this line
		Store:         mapTransactionStoreToResponse(detail.Store),
		Produk:        mapToSimpleProductResponse(detail.ProductLog.Product),
		CreatedAt:     *detail.CreatedAt,
		UpdatedAt:     *detail.UpdatedAt,
	}

	// Then delete
	err = service.repository.Delete(id)
	if err != nil {
		return models.TransactionDetailResponse{}, err
	}

	return response, nil
}

// Update getFotoProdukResponses function
func getFotoProdukResponses(fotos []entities.FotoProduk) []models.FotoProdukResponse {
	var responses []models.FotoProdukResponse
	for _, foto := range fotos {
		var createdAt, updatedAt time.Time
		if foto.CreatedAt != nil {
			createdAt = *foto.CreatedAt
		}
		if foto.UpdatedAt != nil {
			updatedAt = *foto.UpdatedAt
		}

		responses = append(responses, models.FotoProdukResponse{
			ID:        foto.ID,
			URL:       foto.PhotoURL,
			Photo:     foto.Photo,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}
	return responses
}

// Update store response creation
func mapTransactionStoreToResponse(store entities.Store) models.StoreResponse {
	return models.StoreResponse{
		ID:            store.ID,
		IDUser:        store.IDUser, // Add this line
		NamaToko:      store.NamaToko,
		DeskripsiToko: store.DeskripsiToko, // Add this line to include deskripsi_toko
		FotoToko:      mapStorePhotosToResponse(store.FotoToko),
		CreatedAt:     store.CreatedAt,
		UpdatedAt:     store.UpdatedAt,
	}
}

func mapDetailToResponse(detail entities.TransactionDetail) models.TransactionDetailResponse {
	if detail.CreatedAt == nil {
		detail.CreatedAt = &time.Time{}
	}
	if detail.UpdatedAt == nil {
		detail.UpdatedAt = &time.Time{}
	}

	return models.TransactionDetailResponse{
		CreatedAt: *detail.CreatedAt, // Dereference pointer
		UpdatedAt: *detail.UpdatedAt, // Dereference pointer
	}
}

func formatDetailResponse(detail entities.TransactionDetail) models.TransactionDetailResponse {
	if detail.CreatedAt == nil {
		detail.CreatedAt = &time.Time{}
	}
	if detail.UpdatedAt == nil {
		detail.UpdatedAt = &time.Time{}
	}

	return models.TransactionDetailResponse{
		CreatedAt: *detail.CreatedAt, // Dereference pointer
		UpdatedAt: *detail.UpdatedAt, // Dereference pointer
	}
}

// Update the photo mapping in any function that creates FotoProdukResponse
func mapPhotoToResponse(photo entities.FotoProduk) models.FotoProdukResponse {
	var createdAt, updatedAt time.Time
	if photo.CreatedAt != nil {
		createdAt = *photo.CreatedAt
	}
	if photo.UpdatedAt != nil {
		updatedAt = *photo.UpdatedAt
	}

	return models.FotoProdukResponse{
		ID:        photo.ID,
		URL:       photo.URL,
		Photo:     photo.Photo,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
