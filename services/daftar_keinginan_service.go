package services

import (
	"errors"
	"mini-project-evermos/models"
	"mini-project-evermos/models/entities"
	"mini-project-evermos/repositories"
	"time"
)

type WishlistService interface {
	GetAll(userID uint) ([]models.WishlistResponse, error)
	GetById(id uint, userID uint) (models.WishlistResponse, error)
	Create(input models.WishlistRequest, userID uint) (models.WishlistResponse, error)
	Delete(id uint, userID uint) error
	Update(id uint, input models.WishlistRequest, userID uint) (models.WishlistResponse, error)
	ClearAll(userID uint) ([]models.WishlistResponse, error)
}

type wishlistServiceImpl struct {
	repository        repositories.WishlistRepository
	storeRepository   repositories.StoreRepository
	productRepository repositories.ProductRepository
}

func NewWishlistService(
	repository *repositories.WishlistRepository,
	storeRepository *repositories.StoreRepository,
	productRepository *repositories.ProductRepository,
) WishlistService {
	return &wishlistServiceImpl{
		repository:        *repository,
		storeRepository:   *storeRepository,
		productRepository: *productRepository,
	}
}

func (service *wishlistServiceImpl) GetAll(userID uint) ([]models.WishlistResponse, error) {
	wishlists, err := service.repository.FindAll(userID)
	if err != nil {
		return nil, err
	}

	var responses []models.WishlistResponse
	for _, wishlist := range wishlists {
		responses = append(responses, formatWishlistResponse(wishlist))
	}
	return responses, nil
}

func (service *wishlistServiceImpl) GetById(id uint, userID uint) (models.WishlistResponse, error) {
	// Get the wishlist with all its relationships
	wishlist, err := service.repository.FindById(id)
	if err != nil {
		return models.WishlistResponse{}, err
	}

	// Check if the wishlist exists
	if wishlist.ID == 0 {
		return models.WishlistResponse{}, errors.New("wishlist not found")
	}

	// Return the formatted response directly without checking store ownership
	// This matches the behavior of GetAll which shows all wishlists
	return formatWishlistResponse(wishlist), nil
}

func (service *wishlistServiceImpl) Create(input models.WishlistRequest, userID uint) (models.WishlistResponse, error) {
	// Validate that the product exists using product repository
	product, err := service.productRepository.FindById(input.ProductID)
	if (err != nil) || (product.ID == 0) {
		return models.WishlistResponse{}, errors.New("product not found")
	}

	// Create wishlist
	wishlist := entities.Wishlist{
		IDToko:   input.StoreID,
		IDProduk: input.ProductID,
	}

	result, err := service.repository.Insert(wishlist)
	if err != nil {
		return models.WishlistResponse{}, err
	}

	wishlistWithData, err := service.repository.FindById(result.ID)
	if err != nil {
		return models.WishlistResponse{}, err
	}

	return formatWishlistResponse(wishlistWithData), nil
}

func (service *wishlistServiceImpl) Delete(id uint, userID uint) error {
	// Get wishlist data before deletion
	wishlist, err := service.repository.FindById(id)
	if err != nil {
		return err
	}

	if wishlist.ID == 0 {
		return errors.New("wishlist not found")
	}

	// Delete the wishlist
	return service.repository.Delete(id)
}

func (service *wishlistServiceImpl) Update(id uint, input models.WishlistRequest, userID uint) (models.WishlistResponse, error) {
	// Validate that the product exists
	product, err := service.productRepository.FindById(input.ProductID)
	if (err != nil) || (product.ID == 0) {
		return models.WishlistResponse{}, errors.New("product not found")
	}

	// Validate that the store exists
	store, _, err := service.storeRepository.FindById(input.StoreID)
	if (err != nil) || (store.ID == 0) {
		return models.WishlistResponse{}, errors.New("store not found")
	}

	// Update the wishlist directly
	err = service.repository.Update(id, input.StoreID, input.ProductID)
	if err != nil {
		return models.WishlistResponse{}, err
	}

	// Get updated wishlist
	updatedWishlist, err := service.repository.FindById(id)
	if err != nil {
		return models.WishlistResponse{}, err
	}

	return formatWishlistResponse(updatedWishlist), nil
}

func (service *wishlistServiceImpl) ClearAll(userID uint) ([]models.WishlistResponse, error) {
	// Get all wishlists and clear them
	wishlists, err := service.repository.ClearAll(userID)
	if err != nil {
		return nil, err
	}

	// Convert to response format
	var responses []models.WishlistResponse
	for _, wishlist := range wishlists {
		responses = append(responses, formatWishlistResponse(wishlist))
	}

	return responses, nil
}

func formatWishlistResponse(wishlist entities.Wishlist) models.WishlistResponse {
	response := models.WishlistResponse{
		ID:       wishlist.ID,
		IDToko:   wishlist.IDToko,
		IDProduk: wishlist.IDProduk,
		Store: models.StoreResponse{
			ID:            wishlist.Store.ID,
			IDUser:        wishlist.Store.IDUser,
			NamaToko:      wishlist.Store.NamaToko,
			DeskripsiToko: wishlist.Store.DeskripsiToko,
			FotoToko:      mapStorePhotosToResponse(wishlist.Store.FotoToko),
			CreatedAt:     wishlist.Store.CreatedAt,
			UpdatedAt:     wishlist.Store.UpdatedAt,
		},
		CreatedAt: wishlist.CreatedAt,
		UpdatedAt: wishlist.UpdatedAt,
	}

	// Map product photos
	var fotoProdukResponses []models.FotoProdukResponse
	for _, foto := range wishlist.Product.FotoProduk {
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

	// Set Product data including photos
	response.Product.ID = wishlist.Product.ID
	response.Product.NamaProduk = wishlist.Product.NamaProduk
	response.Product.Slug = wishlist.Product.Slug
	response.Product.HargaReseller = wishlist.Product.HargaReseller
	response.Product.HargaKonsumen = wishlist.Product.HargaKonsumen
	response.Product.Stok = wishlist.Product.Stok
	response.Product.Deskripsi = wishlist.Product.Deskripsi
	response.Product.FotoProduk = fotoProdukResponses // Add this line
	response.Product.CreatedAt = wishlist.Product.CreatedAt
	response.Product.UpdatedAt = wishlist.Product.UpdatedAt

	return response
}

func formatProductResponse(product entities.Product) models.ProductResponse {
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
			URL:       foto.PhotoURL,
			Photo:     foto.Photo,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}

	return models.ProductResponse{
		ID:            product.ID,
		NamaProduk:    product.NamaProduk,
		Slug:          product.Slug,
		HargaReseller: product.HargaReseller,
		HargaKonsumen: product.HargaKonsumen,
		Stok:          product.Stok,
		Deskripsi:     product.Deskripsi,
		Store: models.StoreResponse{
			ID:        product.Store.ID,
			NamaToko:  product.Store.NamaToko,
			FotoToko:  nil, // Initialize empty if needed
			CreatedAt: product.Store.CreatedAt,
			UpdatedAt: product.Store.UpdatedAt,
		},
		Category: models.CategoryResponse{
			ID:           product.Category.ID,
			NamaCategory: product.Category.NamaCategory,
			CreatedAt:    product.Category.CreatedAt,
			UpdatedAt:    product.Category.UpdatedAt,
		},
		FotoProduk: fotoProdukResponses,
		CreatedAt:  product.CreatedAt,
		UpdatedAt:  product.UpdatedAt,
	}
}

func mapProductToResponse(product entities.Product) models.ProductResponse {
	return formatProductResponse(product)
}

func mapWishlistToResponse(wishlist entities.Wishlist) models.WishlistResponse {
	return models.WishlistResponse{
		// ...other fields...
		CreatedAt: wishlist.CreatedAt, // Remove & operator
		UpdatedAt: wishlist.UpdatedAt, // Remove & operator
	}
}
