package repositories

import (
	"fmt"
	"math"
	"mini-project-evermos/models"
	"mini-project-evermos/models/entities"
	"mini-project-evermos/models/responder"
	"path/filepath"
	"time"

	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAllPagination(pagination responder.Pagination) (responder.Pagination, error)
	FindById(id uint) (models.ProductResponse, error)
	Insert(product models.ProductRequest) (models.ProductResponse, error)
	Update(id uint, product models.ProductRequest) (models.ProductResponse, error)
	Destroy(id uint) (bool, error)
	FindByCategory(categoryID string) ([]models.ProductResponse, error)
	SearchProducts(query string) ([]models.ProductResponse, error)
	FindRelatedProducts(id uint) ([]models.ProductResponse, error)
	SaveProductPhoto(photo entities.FotoProduk) (entities.FotoProduk, error)
	GetProductPhoto(id uint) (entities.FotoProduk, error)
	SaveProductPhotos(productID uint, photoURLs []interface{}) error
}

type productRepositoryImpl struct {
	database *gorm.DB
}

func NewProductRepository(database *gorm.DB) ProductRepository {
	return &productRepositoryImpl{database}
}

func (repository *productRepositoryImpl) FindAllPagination(request responder.Pagination) (responder.Pagination, error) {
	var products []entities.Product
	var totalRows int64
	query := repository.database.Model(&entities.Product{})
	// Add search condition if keyword is present
	if request.Keyword != "" {
		query = query.Where("nama_produk LIKE ? OR deskripsi LIKE ?",
			"%"+request.Keyword+"%", "%"+request.Keyword+"%")
	}
	// Count total rows with search condition
	query.Count(&totalRows)
	// Build main query with all relations
	query = repository.database.Model(&entities.Product{}).
		Preload("Store", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, id_user, nama_toko, deskripsi_toko, created_at, updated_at") // Added id_user
		}).
		Preload("Store.FotoToko").
		Preload("Category").
		Preload("FotoProduk").
		Preload("Reviews").
		Preload("Reviews.Store").
		Preload("Promos").
		Preload("Promos.Store")
	// Apply search condition to main query
	if request.Keyword != "" {
		query = query.Where("nama_produk LIKE ? OR deskripsi LIKE ?",
			"%"+request.Keyword+"%", "%"+request.Keyword+"%")
	}
	// Apply pagination
	err := query.
		Order("id desc").
		Limit(request.Limit).
		Offset(request.GetOffset()).
		Find(&products).Error
	if err != nil {
		return responder.Pagination{}, err
	}
	var responses []models.ProductResponse
	for _, product := range products {
		responses = append(responses, mapProductToResponse(product))
	}

	request.Rows = responses
	request.TotalRows = totalRows
	request.TotalPages = int(math.Ceil(float64(totalRows) / float64(request.Limit)))

	return request, nil
}

func (repository *productRepositoryImpl) FindById(id uint) (models.ProductResponse, error) {
	var product entities.Product

	result := repository.database.
		Preload("Store", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, id_user, nama_toko, deskripsi_toko, created_at, updated_at") // Added id_user
		}).
		Preload("Store.FotoToko").
		Preload("Category").
		Preload("FotoProduk").
		Preload("Reviews").
		Preload("Reviews.Store").
		Preload("Promos").
		Preload("Promos.Store").
		Preload("Coupons"). // Add this line
		First(&product, id)

	if result.Error != nil {
		return models.ProductResponse{}, fmt.Errorf("product with ID %d not found: %v", id, result.Error)
	}

	return mapProductToResponse(product), nil
}

func (repository *productRepositoryImpl) Insert(input models.ProductRequest) (models.ProductResponse, error) {
	now := time.Now()
	product := entities.Product{
		NamaProduk:    input.NamaProduk,
		IDToko:        input.StoreID,
		IDCategory:    input.CategoryID,
		HargaReseller: input.HargaReseller,
		HargaKonsumen: input.HargaKonsumen,
		Stok:          input.Stok,
		Deskripsi:     &input.Deskripsi,
		Slug:          slug.Make(input.NamaProduk),
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}

	err := repository.database.Create(&product).Error
	if err != nil {
		return models.ProductResponse{}, err
	}

	// Create single photo entry combining file and URL
	if len(input.PhotoURLs) > 0 {
		urlData := input.PhotoURLs[0]
		if urlMap, ok := urlData.(map[string]string); ok {
			foto := entities.FotoProduk{
				IDProduk: product.ID,
				PhotoURL: urlMap["originalName"], // Save the file path to PhotoURL (will be mapped to "photo" in response)
				Photo:    urlMap["url"],          // Save the original URL to Photo (will be mapped to "url" in response)
				FileName: filepath.Base(urlMap["originalName"]),
				FileType: filepath.Ext(urlMap["originalName"]),
				URL:      urlMap["originalName"], // Also save to URL field
			}

			if err := repository.database.Create(&foto).Error; err != nil {
				return models.ProductResponse{}, err
			}
		}
	}

	// After creating the product and photos, fetch the complete product with relationships
	var completeProduct entities.Product
	err = repository.database.
		Preload("Store", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, id_user, nama_toko, deskripsi_toko, created_at, updated_at") // Added id_user
		}).
		Preload("Store.FotoToko").
		Preload("Category").
		Preload("FotoProduk").
		First(&completeProduct, product.ID).Error

	if err != nil {
		return models.ProductResponse{}, err
	}

	return mapProductToResponse(completeProduct), nil
}

func (repository *productRepositoryImpl) Update(id uint, input models.ProductRequest) (models.ProductResponse, error) {
	// Start transaction
	tx := repository.database.Begin()

	// First verify the product exists and get its data
	var existingProduct entities.Product
	if err := tx.Preload("FotoProduk").First(&existingProduct, id).Error; err != nil {
		tx.Rollback()
		return models.ProductResponse{}, fmt.Errorf("product not found: %v", err)
	}

	// Update product fields
	now := time.Now()
	updates := map[string]interface{}{
		"nama_produk":    input.NamaProduk,
		"slug":           slug.Make(input.NamaProduk),
		"harga_reseller": input.HargaReseller,
		"harga_konsumen": input.HargaKonsumen,
		"stok":           input.Stok,
		"deskripsi":      &input.Deskripsi,
		"id_category":    input.CategoryID,
		"id_toko":        input.StoreID,
		"updated_at":     &now,
	}

	// Update the product
	if err := tx.Model(&existingProduct).Updates(updates).Error; err != nil {
		tx.Rollback()
		return models.ProductResponse{}, err
	}

	// Handle photo updates if we have new photos
	if len(input.PhotoURLs) > 0 {
		// Delete existing photos using the transaction
		if err := tx.Delete(&entities.FotoProduk{}, "id_produk = ?", id).Error; err != nil {
			tx.Rollback()
			return models.ProductResponse{}, err
		}

		// Create new photo record
		urlData := input.PhotoURLs[0]
		if urlMap, ok := urlData.(map[string]string); ok {
			foto := entities.FotoProduk{
				IDProduk:  id,
				PhotoURL:  urlMap["originalName"], // Save the file path to PhotoURL (will be mapped to "photo" in response)
				Photo:     urlMap["url"],          // Save the original URL to Photo (will be mapped to "url" in response)
				FileName:  filepath.Base(urlMap["originalName"]),
				FileType:  filepath.Ext(urlMap["originalName"]),
				URL:       urlMap["originalName"],
				CreatedAt: &now,
				UpdatedAt: &now,
			}

			if err := tx.Create(&foto).Error; err != nil {
				tx.Rollback()
				return models.ProductResponse{}, fmt.Errorf("failed to create new photo: %v", err)
			}
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return models.ProductResponse{}, err
	}

	// Fetch updated product with all relationships
	var updatedProduct entities.Product
	err := repository.database.
		Preload("Store").
		Preload("Store.FotoToko").
		Preload("Category").
		Preload("FotoProduk").
		Preload("Reviews").
		Preload("Reviews.Store").
		Preload("Promos").
		Preload("Promos.Store").
		First(&updatedProduct, id).Error

	if err != nil {
		return models.ProductResponse{}, err
	}

	return mapProductToResponse(updatedProduct), nil
}

func (repository *productRepositoryImpl) Destroy(id uint) (bool, error) {
	err := repository.database.Where("id = ?", id).Delete(&entities.Product{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repository *productRepositoryImpl) FindByCategory(categoryID string) ([]models.ProductResponse, error) {
	var products []entities.Product

	err := repository.database.
		Preload("Store").
		Preload("Store.FotoToko"). // Add this line
		Preload("Category").
		Preload("FotoProduk").
		Preload("Reviews").
		Preload("Reviews.Store").
		Preload("Promos").
		Preload("Promos.Store").
		Where("id_category = ?", categoryID).
		Order("id desc"). // Add this line to sort by newest first
		Find(&products).Error

	if err != nil {
		return nil, err
	}

	var responses []models.ProductResponse
	for _, product := range products {
		responses = append(responses, mapProductToResponse(product))
	}
	return responses, nil
}

func (repository *productRepositoryImpl) SearchProducts(query string) ([]models.ProductResponse, error) {
	var products []entities.Product

	err := repository.database.
		Preload("Store").
		Preload("Store.FotoToko").
		Preload("Category").
		Preload("FotoProduk").
		Preload("Reviews").
		Preload("Reviews.Store").
		Preload("Promos").
		Preload("Promos.Store").
		Where("LOWER(nama_produk) LIKE LOWER(?)", "%"+query+"%").
		Order("id desc"). // Add this line to sort by newest first
		Find(&products).Error

	if err != nil {
		return nil, err
	}

	var responses []models.ProductResponse
	for _, product := range products {
		responses = append(responses, mapProductToResponse(product))
	}
	return responses, nil
}

func (repository *productRepositoryImpl) FindRelatedProducts(id uint) ([]models.ProductResponse, error) {
	// First get the category of the current product
	var currentProduct entities.Product
	err := repository.database.First(&currentProduct, id).Error
	if err != nil {
		return nil, err
	}

	// Then find products in the same category, excluding the current product
	// Remove the store condition to show products from all stores
	var products []entities.Product
	err = repository.database.
		Preload("Store").
		Preload("Store.FotoToko"). // Add this line
		Preload("Category").
		Preload("FotoProduk").
		Preload("Reviews").
		Preload("Reviews.Store").
		Preload("Promos").
		Preload("Promos.Store").
		Where("id_category = ? AND id != ?", currentProduct.IDCategory, id).
		Order("id desc"). // Add this line to sort by newest first
		Limit(5).
		Find(&products).Error

	if err != nil {
		return nil, err
	}

	var responses []models.ProductResponse
	for _, product := range products {
		responses = append(responses, mapProductToResponse(product))
	}
	return responses, nil
}

func (repository *productRepositoryImpl) SaveProductPhoto(photo entities.FotoProduk) (entities.FotoProduk, error) {
	// Ensure both URL and PhotoURL are set
	if photo.PhotoURL == "" {
		photo.PhotoURL = photo.URL
	}
	if photo.URL == "" {
		photo.URL = photo.PhotoURL
	}

	result := repository.database.Create(&photo)
	if result.Error != nil {
		return entities.FotoProduk{}, result.Error
	}
	return photo, nil
}

func (repository *productRepositoryImpl) GetProductPhoto(id uint) (entities.FotoProduk, error) {
	var photo entities.FotoProduk
	err := repository.database.First(&photo, id).Error
	return photo, err
}

func (repository *productRepositoryImpl) SaveProductPhotos(productID uint, photoURLs []interface{}) error {
	for _, urlData := range photoURLs {
		if urlMap, ok := urlData.(map[string]string); ok {
			// If it's a map with URL and original name
			photo := entities.FotoProduk{
				IDProduk: productID,
				PhotoURL: urlMap["url"],
				Photo:    urlMap["originalName"],
				FileName: filepath.Base(urlMap["url"]),
				FileType: filepath.Ext(urlMap["originalName"]),
			}
			if err := repository.database.Create(&photo).Error; err != nil {
				return err
			}
		} else if urlStr, ok := urlData.(string); ok {
			// If it's just a URL string
			photo := entities.FotoProduk{
				IDProduk: productID,
				PhotoURL: urlStr,
				Photo:    filepath.Base(urlStr),
				FileName: filepath.Base(urlStr),
				FileType: filepath.Ext(urlStr),
			}
			if err := repository.database.Create(&photo).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

// Helper function to map simplified review response
func mapSimplifiedReviewResponse(review entities.ProductReview) models.SimpleProductReviewResponse {
	return models.SimpleProductReviewResponse{
		ID:        uint(review.ID),
		IDToko:    uint(review.IDToko),
		IDProduk:  uint(review.IDProduk),
		Ulasan:    review.Ulasan,
		Rating:    review.Rating,
		CreatedAt: review.CreatedAt,
		UpdatedAt: review.UpdatedAt,
	}
}

// Update the mapStorePhotosToResponse function
func mapStorePhotosToResponse(photos []entities.StorePhoto) []models.FotoTokoResponse {
	var responses []models.FotoTokoResponse
	for _, photo := range photos {
		var createdAt, updatedAt time.Time
		if photo.CreatedAt != nil {
			createdAt = *photo.CreatedAt
		}
		if photo.UpdatedAt != nil {
			updatedAt = *photo.UpdatedAt
		}
		responses = append(responses, models.FotoTokoResponse{
			ID:        photo.ID,
			URL:       photo.URL,
			Photo:     photo.Photo,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}
	return responses
}

// Update the mapProductToResponse function
func mapProductToResponse(product entities.Product) models.ProductResponse {
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
			Photo:     foto.Photo,    // Change this: map Photo to "photo"
			URL:       foto.PhotoURL, // Change this: map PhotoURL to "url"
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}

	var reviewResponses []models.SimpleProductReviewResponse
	for _, review := range product.Reviews {
		reviewResponses = append(reviewResponses, mapSimplifiedReviewResponse(review))
	}

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

	var couponResponses []models.ProductCouponResponse
	for _, coupon := range product.Coupons {
		couponResponses = append(couponResponses, mapCouponToResponse(coupon))
	}

	// Map store with description and photos
	storeResponse := models.StoreResponse{
		ID:            product.Store.ID,
		IDUser:        product.Store.IDUser, // Add this line
		NamaToko:      product.Store.NamaToko,
		DeskripsiToko: product.Store.DeskripsiToko, // Add this line
		FotoToko:      mapStorePhotosToResponse(product.Store.FotoToko),
		CreatedAt:     product.Store.CreatedAt,
		UpdatedAt:     product.Store.UpdatedAt,
	}

	return models.ProductResponse{
		ID:            product.ID,
		NamaProduk:    product.NamaProduk,
		Slug:          product.Slug,
		HargaReseller: product.HargaReseller,
		HargaKonsumen: product.HargaKonsumen,
		Stok:          product.Stok,
		Deskripsi:     product.Deskripsi,
		Store:         storeResponse, // Use the new store response with photos
		Category: models.CategoryResponse{
			ID:           product.Category.ID,
			NamaCategory: product.Category.NamaCategory,
			CreatedAt:    product.Category.CreatedAt,
			UpdatedAt:    product.Category.UpdatedAt,
		},
		FotoProduk: fotoProdukResponses,
		Reviews:    reviewResponses,
		Promos:     promoResponses,
		Coupons:    couponResponses, // Add this line
		CreatedAt:  product.CreatedAt,
		UpdatedAt:  product.UpdatedAt,
	}
}

func (repository *productRepositoryImpl) createFotoToko(storeFotos []entities.FotoToko) []models.FotoTokoResponse {
	var fotoTokoResponse []models.FotoTokoResponse
	for _, foto := range storeFotos {
		fotoTokoResponse = append(fotoTokoResponse, models.FotoTokoResponse{
			ID:        foto.ID,
			URL:       foto.URL,
			Photo:     foto.Photo,
			CreatedAt: foto.CreatedAt,
			UpdatedAt: foto.UpdatedAt,
		})
	}
	return fotoTokoResponse
}

// Update mapSavedPhotoToResponse function
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
		URL:       savedPhoto.Photo,    // Map Photo to URL
		Photo:     savedPhoto.PhotoURL, // Map PhotoURL to Photo
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
