package repositories

import (
	"fmt"
	"mini-project-evermos/models/entities"
	"time"

	"gorm.io/gorm"
)

type WishlistRepository interface {
	FindAll(userID uint) ([]entities.Wishlist, error)
	FindById(id uint) (entities.Wishlist, error)
	Insert(wishlist entities.Wishlist) (entities.Wishlist, error)
	Delete(id uint) error
	ValidateProduct(productID uint) (entities.Product, error)
	Update(id uint, storeID uint, productID uint) error
	ClearAll(userID uint) ([]entities.Wishlist, error)
}

type wishlistRepositoryImpl struct {
	database *gorm.DB
}

func NewWishlistRepository(database *gorm.DB) WishlistRepository {
	return &wishlistRepositoryImpl{database}
}

func (repository *wishlistRepositoryImpl) FindAll(userID uint) ([]entities.Wishlist, error) {
	var wishlists []entities.Wishlist

	fmt.Printf("Finding wishlists for userID: %d\n", userID)

	// Changed the query to just fetch all wishlists without the store user filter
	err := repository.database.
		Order("daftar_keinginan_belanja.id desc").
		Preload("Store", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, id_user, nama_toko, deskripsi_toko, url_foto, created_at, updated_at")
		}).
		Preload("Store.FotoToko").
		Preload("Product").
		Preload("Product.Category").
		Preload("Product.FotoProduk").
		Preload("Product.Store").
		Find(&wishlists).Error

	if err != nil {
		fmt.Printf("Error finding wishlists: %v\n", err)
		return nil, err
	}

	// Debug print to show what we found
	fmt.Printf("Found %d wishlists\n", len(wishlists))
	for _, w := range wishlists {
		fmt.Printf("Wishlist{ID: %d, StoreID: %d, ProductID: %d}\n", w.ID, w.IDToko, w.IDProduk)
	}

	// Update Store UrlFoto for each wishlist
	for i := range wishlists {
		if len(wishlists[i].Store.FotoToko) > 0 {
			wishlists[i].Store.UrlFoto = wishlists[i].Store.FotoToko[0].URL
		}
	}

	return wishlists, nil
}

func (repository *wishlistRepositoryImpl) FindById(id uint) (entities.Wishlist, error) {
	var wishlist entities.Wishlist
	err := repository.database.
		Order("id desc"). // Add ordering here
		Preload("Store", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, id_user, nama_toko, deskripsi_toko, url_foto, created_at, updated_at") // Add id_user
		}).
		Preload("Store.FotoToko").
		Preload("Product").
		Preload("Product.Category").
		Preload("Product.FotoProduk").
		Preload("Product.Store").
		First(&wishlist, id).Error

	// Update UrlFoto if FotoToko exists
	if err == nil && len(wishlist.Store.FotoToko) > 0 {
		wishlist.Store.UrlFoto = wishlist.Store.FotoToko[0].URL
	}

	return wishlist, err
}

func (repository *wishlistRepositoryImpl) Insert(wishlist entities.Wishlist) (entities.Wishlist, error) {
	now := time.Now()
	wishlist.CreatedAt = &now
	wishlist.UpdatedAt = &now
	err := repository.database.Create(&wishlist).Error
	return wishlist, err
}

func (repository *wishlistRepositoryImpl) Delete(id uint) error {
	return repository.database.Delete(&entities.Wishlist{}, id).Error
}

func (repository *wishlistRepositoryImpl) ValidateProduct(productID uint) (entities.Product, error) {
	var product entities.Product
	err := repository.database.First(&product, productID).Error
	return product, err
}

func (repository *wishlistRepositoryImpl) Update(id uint, storeID uint, productID uint) error {
	return repository.database.Model(&entities.Wishlist{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"id_toko":    storeID,
			"id_produk":  productID,
			"updated_at": time.Now(),
		}).Error
}

func (repository *wishlistRepositoryImpl) ClearAll(userID uint) ([]entities.Wishlist, error) {
	var wishlists []entities.Wishlist

	// First get all wishlists with their relationships before deletion
	err := repository.database.
		Order("id desc").
		Preload("Store", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, id_user, nama_toko, deskripsi_toko, url_foto, created_at, updated_at")
		}).
		Preload("Store.FotoToko").
		Preload("Product").
		Preload("Product.Category").
		Preload("Product.FotoProduk").
		Preload("Product.Store").
		Find(&wishlists).Error

	if err != nil {
		return nil, err
	}

	// Delete all records from the wishlist table
	err = repository.database.Exec("DELETE FROM daftar_keinginan_belanja").Error
	if err != nil {
		return nil, err
	}

	// Update Store UrlFoto for each wishlist before returning
	for i := range wishlists {
		if len(wishlists[i].Store.FotoToko) > 0 {
			wishlists[i].Store.UrlFoto = wishlists[i].Store.FotoToko[0].URL
		}
	}

	return wishlists, nil
}
