package repositories

import (
	"fmt"
	"mini-project-evermos/models"
	"mini-project-evermos/models/entities"
	"strconv"

	"gorm.io/gorm"
)

type DiskonProdukRepository interface {
	ApplyDiscount(productID uint, discountPercent string) (models.DiskonProdukResponse, error)
	GetById(id uint) (models.DiskonProdukResponse, error)
	GetAll() ([]models.DiskonProdukResponse, error)
	UpdateDiscount(id uint, productID uint, discountPercent string) (models.DiskonProdukResponse, error)
	DeleteDiscount(id uint) (models.DiskonProdukResponse, error) // Add this line
}

type diskonProdukRepositoryImpl struct {
	database *gorm.DB
}

func NewDiskonProdukRepository(database *gorm.DB) DiskonProdukRepository {
	return &diskonProdukRepositoryImpl{database}
}

func (repository *diskonProdukRepositoryImpl) ApplyDiscount(productID uint, discountPercent string) (models.DiskonProdukResponse, error) {
	// Start transaction
	tx := repository.database.Begin()

	// First get the product
	var product entities.Product
	if err := tx.First(&product, productID).Error; err != nil {
		tx.Rollback()
		return models.DiskonProdukResponse{}, err
	}

	// Check if this is the first discount
	if product.HargaOriginal == "" {
		// Store the original price if this is the first discount
		product.HargaOriginal = product.HargaKonsumen
		if err := tx.Save(&product).Error; err != nil {
			tx.Rollback()
			return models.DiskonProdukResponse{}, err
		}
	}

	// Use original price for discount calculation
	originalPrice, err := strconv.ParseFloat(product.HargaOriginal, 64)
	if err != nil {
		tx.Rollback()
		return models.DiskonProdukResponse{}, err
	}

	// Convert discount percentage to float
	discountPercentFloat, err := strconv.ParseFloat(discountPercent, 64)
	if err != nil {
		tx.Rollback()
		return models.DiskonProdukResponse{}, err
	}

	// Calculate new price based on original price
	discountAmount := originalPrice * (discountPercentFloat / 100)
	newPrice := originalPrice - discountAmount
	newPriceStr := fmt.Sprintf("%.0f", newPrice)

	// Create new discount record
	diskon := entities.DiskonProduk{
		ProductID:     productID,
		HargaKonsumen: discountPercent + "%", // Store the percentage
	}

	if err := tx.Create(&diskon).Error; err != nil {
		tx.Rollback()
		return models.DiskonProdukResponse{}, err
	}

	// Update product's current price with the new discounted price
	if err := tx.Model(&product).Update("harga_konsumen", newPriceStr).Error; err != nil {
		tx.Rollback()
		return models.DiskonProdukResponse{}, err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return models.DiskonProdukResponse{}, err
	}

	return models.DiskonProdukResponse{
		ID:            diskon.ID,
		ProductID:     diskon.ProductID,
		HargaKonsumen: diskon.HargaKonsumen,
		CreatedAt:     diskon.CreatedAt,
		UpdatedAt:     diskon.UpdatedAt,
	}, nil
}

func (repository *diskonProdukRepositoryImpl) GetById(id uint) (models.DiskonProdukResponse, error) {
	var diskon entities.DiskonProduk

	err := repository.database.First(&diskon, id).Error
	if err != nil {
		return models.DiskonProdukResponse{}, err
	}

	return models.DiskonProdukResponse{
		ID:            diskon.ID,
		ProductID:     diskon.ProductID,
		HargaKonsumen: diskon.HargaKonsumen,
		CreatedAt:     diskon.CreatedAt,
		UpdatedAt:     diskon.UpdatedAt,
	}, nil
}

func (repository *diskonProdukRepositoryImpl) GetAll() ([]models.DiskonProdukResponse, error) {
	var diskons []entities.DiskonProduk

	// Add Order("created_at DESC") to get newest records first
	err := repository.database.Order("created_at DESC").Find(&diskons).Error
	if err != nil {
		return nil, err
	}

	var responses []models.DiskonProdukResponse
	for _, diskon := range diskons {
		responses = append(responses, models.DiskonProdukResponse{
			ID:            diskon.ID,
			ProductID:     diskon.ProductID,
			HargaKonsumen: diskon.HargaKonsumen,
			CreatedAt:     diskon.CreatedAt,
			UpdatedAt:     diskon.UpdatedAt,
		})
	}

	return responses, nil
}

func (repository *diskonProdukRepositoryImpl) UpdateDiscount(id uint, productID uint, discountPercent string) (models.DiskonProdukResponse, error) {
	// Start transaction
	tx := repository.database.Begin()

	// First get the product
	var product entities.Product
	if err := tx.First(&product, productID).Error; err != nil {
		tx.Rollback()
		return models.DiskonProdukResponse{}, err
	}

	// Use HargaOriginal if it exists, otherwise use current HargaKonsumen
	originalPrice, err := strconv.ParseFloat(product.HargaOriginal, 64)
	if err != nil {
		// If HargaOriginal is empty or invalid, use current HargaKonsumen
		originalPrice, err = strconv.ParseFloat(product.HargaKonsumen, 64)
		if err != nil {
			tx.Rollback()
			return models.DiskonProdukResponse{}, err
		}
		// Store the original price if not already stored
		product.HargaOriginal = product.HargaKonsumen
		if err := tx.Save(&product).Error; err != nil {
			tx.Rollback()
			return models.DiskonProdukResponse{}, err
		}
	}

	// Convert discount percentage to float
	discountPercentFloat, err := strconv.ParseFloat(discountPercent, 64)
	if err != nil {
		tx.Rollback()
		return models.DiskonProdukResponse{}, err
	}

	// Calculate new price based on original price
	discountAmount := originalPrice * (discountPercentFloat / 100)
	newPrice := originalPrice - discountAmount
	newPriceStr := fmt.Sprintf("%.0f", newPrice)

	// Update existing discount record
	var diskon entities.DiskonProduk
	if err := tx.First(&diskon, id).Error; err != nil {
		tx.Rollback()
		return models.DiskonProdukResponse{}, err
	}

	diskon.ProductID = productID
	diskon.HargaKonsumen = discountPercent + "%"

	if err := tx.Save(&diskon).Error; err != nil {
		tx.Rollback()
		return models.DiskonProdukResponse{}, err
	}

	// Update product's price
	if err := tx.Model(&product).Update("harga_konsumen", newPriceStr).Error; err != nil {
		tx.Rollback()
		return models.DiskonProdukResponse{}, err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return models.DiskonProdukResponse{}, err
	}

	return models.DiskonProdukResponse{
		ID:            diskon.ID,
		ProductID:     diskon.ProductID,
		HargaKonsumen: diskon.HargaKonsumen,
		CreatedAt:     diskon.CreatedAt,
		UpdatedAt:     diskon.UpdatedAt,
	}, nil
}

func (repository *diskonProdukRepositoryImpl) DeleteDiscount(id uint) (models.DiskonProdukResponse, error) {
	// Start transaction
	tx := repository.database.Begin()

	// First get the discount record
	var diskon entities.DiskonProduk
	if err := tx.First(&diskon, id).Error; err != nil {
		tx.Rollback()
		return models.DiskonProdukResponse{}, err
	}

	// Get the product to restore original price
	var product entities.Product
	if err := tx.First(&product, diskon.ProductID).Error; err != nil {
		tx.Rollback()
		return models.DiskonProdukResponse{}, err
	}

	// Restore original price if it exists
	if product.HargaOriginal != "" {
		if err := tx.Model(&product).Update("harga_konsumen", product.HargaOriginal).Error; err != nil {
			tx.Rollback()
			return models.DiskonProdukResponse{}, err
		}
		// Clear the original price
		if err := tx.Model(&product).Update("harga_original", "").Error; err != nil {
			tx.Rollback()
			return models.DiskonProdukResponse{}, err
		}
	}

	// Create response before deleting the record
	response := models.DiskonProdukResponse{
		ID:            diskon.ID,
		ProductID:     diskon.ProductID,
		HargaKonsumen: diskon.HargaKonsumen,
		CreatedAt:     diskon.CreatedAt,
		UpdatedAt:     diskon.UpdatedAt,
	}

	// Delete the discount record
	if err := tx.Delete(&diskon).Error; err != nil {
		tx.Rollback()
		return models.DiskonProdukResponse{}, err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return models.DiskonProdukResponse{}, err
	}

	return response, nil
}
