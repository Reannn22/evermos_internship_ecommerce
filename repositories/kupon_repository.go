package repositories

import (
	"fmt"
	"mini-project-evermos/models"
	"mini-project-evermos/models/entities"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type ProductCouponRepository interface {
	FindAll() ([]models.ProductCouponResponse, error)
	FindById(id uint) (models.ProductCouponResponse, error)
	FindByCode(code string) (models.ProductCouponResponse, error)
	FindByProduct(productId uint) ([]models.ProductCouponResponse, error)
	Create(coupon models.ProductCouponRequest) (models.ProductCouponResponse, error)
	Update(id uint, coupon models.ProductCouponRequest) (models.ProductCouponResponse, error)
	Delete(id uint) (models.ProductCouponResponse, error) // Changed return type
	ValidateCoupon(code string, productId uint) (models.ProductCouponResponse, error)
}

type productCouponRepositoryImpl struct {
	db *gorm.DB
}

func NewProductCouponRepository(db *gorm.DB) ProductCouponRepository {
	return &productCouponRepositoryImpl{db}
}

func (r *productCouponRepositoryImpl) FindAll() ([]models.ProductCouponResponse, error) {
	var coupons []entities.ProductCoupon
	if err := r.db.
		Order("id DESC"). // Add this line to order by ID descending
		Preload("Product").
		Find(&coupons).Error; err != nil {
		return nil, err
	}

	var responses []models.ProductCouponResponse
	for _, coupon := range coupons {
		responses = append(responses, mapCouponToResponse(coupon))
	}
	return responses, nil
}

func (r *productCouponRepositoryImpl) FindById(id uint) (models.ProductCouponResponse, error) {
	var coupon entities.ProductCoupon
	if err := r.db.
		Preload("Product").
		Preload("Product.Store").
		Preload("Product.Store.FotoToko").
		Preload("Product.Category").
		Preload("Product.FotoProduk").
		Preload("Product.Reviews").
		Preload("Product.Promos").
		First(&coupon, id).Error; err != nil {
		return models.ProductCouponResponse{}, err
	}
	return mapCouponToResponse(coupon), nil
}

func (r *productCouponRepositoryImpl) FindByCode(code string) (models.ProductCouponResponse, error) {
	var coupon entities.ProductCoupon
	if err := r.db.
		Preload("Product").
		Preload("Product.Store").
		Preload("Product.Store.FotoToko").
		Preload("Product.Category").
		Preload("Product.FotoProduk").
		Preload("Product.Reviews").
		Preload("Product.Promos").
		Where("kode_kupon = ?", code).
		First(&coupon).Error; err != nil {
		return models.ProductCouponResponse{}, err
	}
	return mapCouponToResponse(coupon), nil
}

func (r *productCouponRepositoryImpl) FindByProduct(productId uint) ([]models.ProductCouponResponse, error) {
	var coupons []entities.ProductCoupon
	if err := r.db.
		Preload("Product").
		Preload("Product.Store").
		Preload("Product.Store.FotoToko").
		Preload("Product.Category").
		Preload("Product.FotoProduk").
		Preload("Product.Reviews").
		Preload("Product.Promos").
		Where("id_produk = ?", productId).
		Find(&coupons).Error; err != nil {
		return nil, err
	}

	var responses []models.ProductCouponResponse
	for _, coupon := range coupons {
		responses = append(responses, mapCouponToResponse(coupon))
	}
	return responses, nil
}

func (r *productCouponRepositoryImpl) Create(request models.ProductCouponRequest) (models.ProductCouponResponse, error) {
	tx := r.db.Begin()

	// Get product to update price
	var product entities.Product
	if err := tx.First(&product, request.ProductID).Error; err != nil {
		tx.Rollback()
		return models.ProductCouponResponse{}, err
	}

	// Store original price if not already stored
	if product.HargaOriginal == "" {
		// Update the product with the original price
		if err := tx.Model(&product).Updates(map[string]interface{}{
			"harga_original": product.HargaKonsumen,
		}).Error; err != nil {
			tx.Rollback()
			return models.ProductCouponResponse{}, err
		}
	}

	// Delete all existing coupons for this product
	if err := tx.Where("id_produk = ?", request.ProductID).Delete(&entities.ProductCoupon{}).Error; err != nil {
		tx.Rollback()
		return models.ProductCouponResponse{}, err
	}

	// Calculate new discounted price based on original price or current price
	originalPrice := product.HargaOriginal
	if originalPrice == "" {
		originalPrice = product.HargaKonsumen
	}
	price, _ := strconv.ParseFloat(originalPrice, 64)
	discountedPrice := price - (price * request.Discount / 100)

	// Create new coupon
	now := time.Now()
	coupon := entities.ProductCoupon{
		IDProduk:  request.ProductID,
		Code:      request.Code,
		Discount:  request.Discount,
		ValidFrom: request.ValidFrom,
		ValidTo:   request.ValidTo,
		IsActive:  true,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	if err := tx.Create(&coupon).Error; err != nil {
		tx.Rollback()
		return models.ProductCouponResponse{}, err
	}

	// Update product with discounted price
	if err := tx.Model(&product).Updates(map[string]interface{}{
		"harga_konsumen": fmt.Sprintf("%.0f", discountedPrice),
		"updated_at":     now,
	}).Error; err != nil {
		tx.Rollback()
		return models.ProductCouponResponse{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return models.ProductCouponResponse{}, err
	}

	return r.FindById(coupon.ID)
}

func (r *productCouponRepositoryImpl) Update(id uint, request models.ProductCouponRequest) (models.ProductCouponResponse, error) {
	// Start transaction
	tx := r.db.Begin()

	// Get existing coupon
	var existingCoupon entities.ProductCoupon
	if err := tx.First(&existingCoupon, id).Error; err != nil {
		tx.Rollback()
		return models.ProductCouponResponse{}, err
	}

	// Get product to update price
	var product entities.Product
	if err := tx.First(&product, request.ProductID).Error; err != nil {
		tx.Rollback()
		return models.ProductCouponResponse{}, err
	}

	// Calculate new discounted price
	originalPriceKonsumen, _ := strconv.ParseFloat(product.HargaKonsumen, 64)
	discountedPriceKonsumen := originalPriceKonsumen - (originalPriceKonsumen * request.Discount / 100)

	now := time.Now()
	// Use correct column names that match the database
	updates := map[string]interface{}{
		"id_produk":      request.ProductID,
		"kode_kupon":     request.Code,
		"diskon":         request.Discount,
		"berlaku_dari":   request.ValidFrom,
		"berlaku_sampai": request.ValidTo,
		"updated_at":     now,
	}

	// Update the coupon
	if err := tx.Model(&existingCoupon).Updates(updates).Error; err != nil {
		tx.Rollback()
		return models.ProductCouponResponse{}, err
	}

	// Update product price
	if err := tx.Model(&product).Updates(map[string]interface{}{
		"harga_konsumen": fmt.Sprintf("%.0f", discountedPriceKonsumen),
		"updated_at":     now,
	}).Error; err != nil {
		tx.Rollback()
		return models.ProductCouponResponse{}, err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return models.ProductCouponResponse{}, err
	}

	return r.FindById(id)
}

func (r *productCouponRepositoryImpl) Delete(id uint) (models.ProductCouponResponse, error) {
	tx := r.db.Begin()

	// Get the coupon with product information
	var coupon entities.ProductCoupon
	if err := tx.Preload("Product").First(&coupon, id).Error; err != nil {
		tx.Rollback()
		return models.ProductCouponResponse{}, err
	}

	// Store the response before deletion
	response := mapCouponToResponse(coupon)

	// Get the complete product information
	var product entities.Product
	if err := tx.First(&product, coupon.IDProduk).Error; err != nil {
		tx.Rollback()
		return models.ProductCouponResponse{}, err
	}

	// Delete the coupon
	if err := tx.Delete(&coupon).Error; err != nil {
		tx.Rollback()
		return models.ProductCouponResponse{}, err
	}

	// Count remaining active coupons
	var remainingCoupons int64
	if err := tx.Model(&entities.ProductCoupon{}).
		Where("id_produk = ? AND aktif = ? AND id != ?", coupon.IDProduk, true, id).
		Count(&remainingCoupons).Error; err != nil {
		tx.Rollback()
		return models.ProductCouponResponse{}, err
	}

	// If no more active coupons exist, restore original price
	if remainingCoupons == 0 && product.HargaOriginal != "" {
		// Update product with original price
		if err := tx.Model(&product).Updates(map[string]interface{}{
			"harga_konsumen": product.HargaOriginal,
			"updated_at":     time.Now(),
		}).Error; err != nil {
			tx.Rollback()
			return models.ProductCouponResponse{}, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return models.ProductCouponResponse{}, err
	}

	return response, nil
}

func (r *productCouponRepositoryImpl) ValidateCoupon(code string, productId uint) (models.ProductCouponResponse, error) {
	var coupon entities.ProductCoupon
	now := time.Now()

	err := r.db.
		Preload("Product").
		Preload("Product.Store").
		Preload("Product.Store.FotoToko").
		Preload("Product.Category").
		Preload("Product.FotoProduk").
		Preload("Product.Reviews").
		Preload("Product.Promos").
		Where("kode_kupon = ? AND id_produk = ? AND aktif = ? AND berlaku_dari <= ? AND berlaku_sampai >= ?",
			code, productId, true, now, now).
		First(&coupon).Error
	if err != nil {
		return models.ProductCouponResponse{}, err
	}

	return mapCouponToResponse(coupon), nil
}

func mapCouponToResponse(coupon entities.ProductCoupon) models.ProductCouponResponse {
	var createdAt, updatedAt time.Time
	if coupon.CreatedAt != nil {
		createdAt = *coupon.CreatedAt
	}
	if coupon.UpdatedAt != nil {
		updatedAt = *coupon.UpdatedAt
	}

	return models.ProductCouponResponse{
		ID:        coupon.ID,
		ProductID: coupon.IDProduk,
		Code:      coupon.Code,
		Discount:  coupon.Discount,
		ValidFrom: coupon.ValidFrom,
		ValidTo:   coupon.ValidTo,
		IsActive:  coupon.IsActive,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
