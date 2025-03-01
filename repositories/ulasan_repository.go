package repositories

import (
	"mini-project-evermos/models"
	"mini-project-evermos/models/entities"
	"time"

	"gorm.io/gorm"
)

type ProductReviewRepository interface {
	FindAll() ([]entities.ProductReview, error)
	FindById(id uint64) (entities.ProductReview, error)
	FindByProductId(productId uint32) ([]entities.ProductReview, error)
	Insert(review models.ProductReviewRequest, storeId uint64) (entities.ProductReview, error)
	Update(review models.ProductReviewRequest, id uint64) (entities.ProductReview, error)
	Delete(id uint64) error
}

type productReviewRepositoryImpl struct {
	db *gorm.DB
}

func NewProductReviewRepository(db *gorm.DB) ProductReviewRepository {
	return &productReviewRepositoryImpl{db}
}

func (r *productReviewRepositoryImpl) FindAll() ([]entities.ProductReview, error) {
	var reviews []entities.ProductReview
	err := r.db.
		Order("id desc"). // Add ordering here
		Preload("Store", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, id_user, nama_toko, deskripsi_toko, created_at, updated_at")
		}).
		Preload("Store.FotoToko").
		Preload("Product.FotoProduk").
		Preload("Product", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, nama_produk, slug, harga_reseller, harga_konsumen, stok, deskripsi, created_at, updated_at")
		}).
		Find(&reviews).Error

	// Handle potential nil pointers and set default values
	for i := range reviews {
		if reviews[i].Store.UrlFoto == "" {
			reviews[i].Store.UrlFoto = "h22222teeeeeetps://i2.pickpik.com/photos/997/209/75/retail-grocery-supermarket-store-preview32222.jpppppg"
		}

		// Ensure timestamps aren't nil
		if reviews[i].CreatedAt == nil {
			now := time.Now()
			reviews[i].CreatedAt = &now
		}
		if reviews[i].UpdatedAt == nil {
			now := time.Now()
			reviews[i].UpdatedAt = &now
		}
		if reviews[i].Store.CreatedAt == nil {
			now := time.Now()
			reviews[i].Store.CreatedAt = &now
		}
		if reviews[i].Store.UpdatedAt == nil {
			now := time.Now()
			reviews[i].Store.UpdatedAt = &now
		}
	}

	return reviews, err
}

func (r *productReviewRepositoryImpl) FindById(id uint64) (entities.ProductReview, error) {
	var review entities.ProductReview
	err := r.db.Table("product_reviews").
		Order("id desc"). // Add ordering here
		Preload("Store", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, id_user, nama_toko, deskripsi_toko, created_at, updated_at")
		}).
		Preload("Store.FotoToko").
		Preload("Product").
		Preload("Product.FotoProduk").
		First(&review, id).Error

	if err == nil {
		// Set default URL if none exists
		if review.Store.UrlFoto == "" {
			review.Store.UrlFoto = "h22222teeeeeetps://i2.pickpik.com/photos/997/209/75/retail-grocery-supermarket-store-preview32222.jpppppg"
		}
	}

	return review, err
}

func (r *productReviewRepositoryImpl) FindByProductId(productId uint32) ([]entities.ProductReview, error) {
	var reviews []entities.ProductReview
	err := r.db.
		Order("id desc"). // Add ordering here
		Preload("Store", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, id_user, nama_toko, deskripsi_toko, created_at, updated_at")
		}).
		Preload("Store.FotoToko").
		Preload("Product").
		Preload("Product.FotoProduk").
		Where("id_produk = ?", productId).
		Find(&reviews).Error

	if err == nil {
		// Set default URL for all stores if needed
		for i := range reviews {
			if reviews[i].Store.UrlFoto == "" {
				reviews[i].Store.UrlFoto = "h22222teeeeeetps://i2.pickpik.com/photos/997/209/75/retail-grocery-supermarket-store-preview32222.jpppppg"
			}
		}
	}

	return reviews, err
}

func (r *productReviewRepositoryImpl) Insert(review models.ProductReviewRequest, storeId uint64) (entities.ProductReview, error) {
	now := time.Now()
	newReview := entities.ProductReview{
		IDToko:    storeId,
		IDProduk:  review.IDProduk,
		Ulasan:    review.Ulasan,
		Rating:    review.Rating,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	err := r.db.Create(&newReview).Error
	if err != nil {
		return entities.ProductReview{}, err
	}

	// Get complete review with relationships, explicitly selecting store fields
	var completeReview entities.ProductReview
	err = r.db.
		Preload("Store", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, id_user, nama_toko, deskripsi_toko, created_at, updated_at") // Added id_user here
		}).
		Preload("Store.FotoToko").
		Preload("Product").
		Preload("Product.FotoProduk").
		First(&completeReview, newReview.ID).Error

	// Set default URL if none exists
	if err == nil && completeReview.Store.UrlFoto == "" {
		completeReview.Store.UrlFoto = "https://i2.pickpik.com/photos/997/209/75/retail-grocery-supermarket-store-preview32222.jpg"
	}

	return completeReview, err
}

func (r *productReviewRepositoryImpl) Update(review models.ProductReviewRequest, id uint64) (entities.ProductReview, error) {
	var existingReview entities.ProductReview
	now := time.Now()

	err := r.db.First(&existingReview, id).Error
	if err != nil {
		return entities.ProductReview{}, err
	}

	// Update only the fields that should be updatable
	existingReview.Ulasan = review.Ulasan
	existingReview.Rating = review.Rating
	existingReview.IDToko = uint64(review.IDToko)
	existingReview.IDProduk = uint(review.IDProduk)
	existingReview.UpdatedAt = &now

	err = r.db.Save(&existingReview).Error
	if err != nil {
		return entities.ProductReview{}, err
	}

	// Fetch the updated review with all relations
	var updatedReview entities.ProductReview
	err = r.db.Model(&entities.ProductReview{}).
		Preload("Store", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, id_user, nama_toko, deskripsi_toko, created_at, updated_at") // Added id_user here
		}).
		Preload("Store.FotoToko").
		Preload("Product").
		Preload("Product.FotoProduk").
		First(&updatedReview, id).Error
	return updatedReview, err
}

func (r *productReviewRepositoryImpl) Delete(id uint64) error {
	return r.db.Delete(&entities.ProductReview{}, id).Error
}
