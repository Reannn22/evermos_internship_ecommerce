package services

import (
	"errors"
	"mini-project-evermos/models"
	"mini-project-evermos/models/entities"
	"mini-project-evermos/repositories"
	"time"
)

type ProductReviewService interface {
	GetAll() ([]models.ProductReviewResponse, error)
	GetById(id uint64) (models.ProductReviewResponse, error)
	GetByProductId(productId uint32) ([]models.ProductReviewResponse, error)
	Create(input models.ProductReviewRequest) (models.ProductReviewResponse, error)
	Update(input models.ProductReviewRequest, id uint64, userId uint64) (models.ProductReviewResponse, error)
	Delete(id uint64, userId uint64) error
}

type productReviewServiceImpl struct {
	reviewRepository repositories.ProductReviewRepository
	storeRepository  repositories.StoreRepository
}

func NewProductReviewService(
	reviewRepository repositories.ProductReviewRepository,
	storeRepository repositories.StoreRepository,
) ProductReviewService {
	return &productReviewServiceImpl{
		reviewRepository: reviewRepository,
		storeRepository:  storeRepository,
	}
}

// Update the mapReviewResponse method to remove UrlFoto
func (s *productReviewServiceImpl) mapReviewResponse(review entities.ProductReview) models.ProductReviewResponse {
	// Handle nil Product.Deskripsi
	deskripsi := ""
	if review.Product.Deskripsi != nil {
		deskripsi = *review.Product.Deskripsi
	}

	// Update store response to include IDUser
	storeResponse := models.StoreResponse{
		ID:            review.Store.ID,
		IDUser:        review.Store.IDUser, // Add this line
		NamaToko:      review.Store.NamaToko,
		DeskripsiToko: review.Store.DeskripsiToko, // Add this line
		FotoToko:      mapStorePhotosToResponse(review.Store.FotoToko),
		CreatedAt:     review.Store.CreatedAt,
		UpdatedAt:     review.Store.UpdatedAt,
	}

	// Map product photos
	var fotoProdukResponses []models.FotoProdukResponse
	for _, foto := range review.Product.FotoProduk {
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

	return models.ProductReviewResponse{
		ID:       uint(review.ID),
		IDToko:   uint(review.IDToko),
		IDProduk: uint(review.IDProduk),
		Ulasan:   review.Ulasan,
		Rating:   review.Rating,
		Store:    storeResponse,
		Product: models.ProductDetail{
			ID:            review.Product.ID,
			NamaProduk:    review.Product.NamaProduk,
			Slug:          review.Product.Slug,
			HargaReseller: review.Product.HargaReseller,
			HargaKonsumen: review.Product.HargaKonsumen,
			Stok:          review.Product.Stok,
			Deskripsi:     deskripsi,
			FotoProduk:    fotoProdukResponses, // Add this field
			CreatedAt:     review.Product.CreatedAt,
			UpdatedAt:     review.Product.UpdatedAt,
		},
		CreatedAt: review.CreatedAt,
		UpdatedAt: review.UpdatedAt,
	}
}

// Update mapReviewStoreToResponse
func mapReviewStoreToResponse(store entities.Store) models.StoreResponse {
	return models.StoreResponse{
		ID:            store.ID,
		NamaToko:      store.NamaToko,
		DeskripsiToko: store.DeskripsiToko, // Make sure this field is included
		FotoToko:      mapStorePhotosToResponse(store.FotoToko),
		CreatedAt:     store.CreatedAt,
		UpdatedAt:     store.UpdatedAt,
	}
}

func (s *productReviewServiceImpl) GetAll() ([]models.ProductReviewResponse, error) {
	reviews, err := s.reviewRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []models.ProductReviewResponse
	for _, review := range reviews {
		responses = append(responses, s.mapReviewResponse(review))
	}

	return responses, nil
}

func (s *productReviewServiceImpl) GetById(id uint64) (models.ProductReviewResponse, error) {
	review, err := s.reviewRepository.FindById(id)
	if err != nil {
		return models.ProductReviewResponse{}, err
	}
	return s.mapReviewResponse(review), nil
}

func (s *productReviewServiceImpl) GetByProductId(productId uint32) ([]models.ProductReviewResponse, error) {
	reviews, err := s.reviewRepository.FindByProductId(productId)
	if err != nil {
		return nil, err
	}

	var responses []models.ProductReviewResponse
	for _, review := range reviews {
		responses = append(responses, s.mapReviewResponse(review))
	}

	return responses, nil
}

func (s *productReviewServiceImpl) Create(input models.ProductReviewRequest) (models.ProductReviewResponse, error) {
	review, err := s.reviewRepository.Insert(input, uint64(input.IDToko))
	if err != nil {
		return models.ProductReviewResponse{}, err
	}

	// Set default store URL if it's empty
	if review.Store.UrlFoto == "" {
		review.Store.UrlFoto = "https://i2.pickpik.com/photos/997/209/75/retail-grocery-supermarket-store-preview32222.jpg"
	}

	return s.mapReviewResponse(review), nil
}

func (s *productReviewServiceImpl) Update(input models.ProductReviewRequest, id uint64, userId uint64) (models.ProductReviewResponse, error) {
	// Get the existing review for validation
	existingReview, err := s.reviewRepository.FindById(id)
	if err != nil {
		return models.ProductReviewResponse{}, err
	}

	// Validate if the review exists
	if existingReview.ID == 0 {
		return models.ProductReviewResponse{}, errors.New("review not found")
	}

	// Update the review
	updatedReview, err := s.reviewRepository.Update(input, id)
	if err != nil {
		return models.ProductReviewResponse{}, err
	}

	return s.GetById(updatedReview.ID)
}

func (s *productReviewServiceImpl) Delete(id uint64, userId uint64) error {
	// Delete the review directly
	return s.reviewRepository.Delete(id)
}
