package repositories

import (
	"fmt"
	"mini-project-evermos/models/entities"
	"mini-project-evermos/models/responder"
	"path/filepath"
	"strconv"

	"gorm.io/gorm"
)

// Contract
type StoreRepository interface {
	FindAllPagination(pagination responder.Pagination) ([]entities.Store, int, error)
	FindById(id uint) (entities.Store, []entities.StorePhoto, error)
	FindByUserId(id uint) (entities.Store, error)
	Update(id uint, store entities.Store, photoURLs []interface{}) (bool, error)
	Insert(store entities.Store, photoURLs []interface{}) (entities.Store, error)
	Delete(id uint) (bool, error)
}

type storeRepositoryImpl struct {
	database *gorm.DB
}

func NewStoreRepository(database *gorm.DB) StoreRepository {
	return &storeRepositoryImpl{database}
}

func (repository *storeRepositoryImpl) FindAllPagination(pagination responder.Pagination) ([]entities.Store, int, error) {
	var stores []entities.Store
	var total int64

	query := repository.database.Model(&entities.Store{})

	if pagination.Keyword != "" {
		query = query.Where("nama_toko LIKE ?", "%"+pagination.Keyword+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (pagination.Page - 1) * pagination.Limit

	// Add Preload for FotoToko and ensure proper ordering
	err = query.
		Select("id, id_user, nama_toko, deskripsi_toko, url_foto, created_at, updated_at"). // Add this line to ensure id_user is selected
		Preload("FotoToko", func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at DESC")
		}).
		Order("id DESC").
		Offset(offset).
		Limit(pagination.Limit).
		Find(&stores).Error

	if err != nil {
		return nil, 0, err
	}

	return stores, int(total), nil
}

func (repository *storeRepositoryImpl) FindById(id uint) (entities.Store, []entities.StorePhoto, error) {
	var store entities.Store
	var photos []entities.StorePhoto

	// Get the store with eager loading of photos
	err := repository.database.
		Select("id, id_user, nama_toko, deskripsi_toko, url_foto, created_at, updated_at"). // Make sure id_user is selected
		Preload("FotoToko", func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at DESC").Limit(1) // Only get the latest photo
		}).
		Where("id = ?", id).
		First(&store).Error
	if err != nil {
		return store, nil, err
	}

	// Get only the latest photo for this store
	err = repository.database.
		Where("id_toko = ?", id).
		Order("created_at DESC").
		Limit(1). // Add this line to get only one photo
		Find(&photos).Error
	if err != nil {
		return store, nil, err
	}

	return store, photos, nil
}

func (repository *storeRepositoryImpl) FindByUserId(id uint) (entities.Store, error) {
	var store entities.Store
	err := repository.database.Where("id_user = ?", id).First(&store).Error

	if err != nil {
		return store, err
	}

	return store, nil
}

func (repository *storeRepositoryImpl) Update(id uint, store entities.Store, photoURLs []interface{}) (bool, error) {
	tx := repository.database.Begin()

	// Update store fields
	updates := map[string]interface{}{
		"nama_toko":      store.NamaToko,
		"deskripsi_toko": store.DeskripsiToko,
		"url_foto":       store.UrlFoto,
	}

	if err := tx.Model(&entities.Store{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	// Delete existing photos
	if err := tx.Where("id_toko = ?", id).Delete(&entities.StorePhoto{}).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	// Create new photos if provided
	if len(photoURLs) > 0 {
		for _, urlData := range photoURLs {
			if urlMap, ok := urlData.(map[string]string); ok {
				photo := entities.StorePhoto{
					IdToko:   id,
					URL:      urlMap["url"],
					Photo:    fmt.Sprintf("/uploads/%s", urlMap["originalName"]),
					FileName: urlMap["originalName"],
					FileType: filepath.Ext(urlMap["originalName"]),
				}
				if err := tx.Create(&photo).Error; err != nil {
					tx.Rollback()
					return false, err
				}
			}
		}
	}

	tx.Commit()
	return true, nil
}

func (repository *storeRepositoryImpl) Insert(store entities.Store, photoURLs []interface{}) (entities.Store, error) {
	tx := repository.database.Begin()

	storeToCreate := entities.Store{
		IDUser:        store.IDUser,
		NamaToko:      store.NamaToko,
		DeskripsiToko: store.DeskripsiToko,
		UrlFoto:       store.UrlFoto,
	}

	if err := tx.Create(&storeToCreate).Error; err != nil {
		tx.Rollback()
		return entities.Store{}, err
	}

	// Create store photos
	for _, urlData := range photoURLs {
		if urlMap, ok := urlData.(map[string]string); ok {
			// Get the id_foto from the map
			idFoto := uint(1) // default value
			if idFotoStr, exists := urlMap["id_foto"]; exists {
				if parsed, err := strconv.ParseUint(idFotoStr, 10, 32); err == nil {
					idFoto = uint(parsed)
				}
			}

			photo := entities.StorePhoto{
				IdToko:   storeToCreate.ID,
				IdFoto:   idFoto, // Use the parsed id_foto value
				URL:      urlMap["url"],
				Photo:    fmt.Sprintf("/uploads/%s", urlMap["originalName"]),
				FileName: urlMap["originalName"],
				FileType: filepath.Ext(urlMap["originalName"]),
			}
			if err := tx.Create(&photo).Error; err != nil {
				tx.Rollback()
				return entities.Store{}, err
			}
		}
	}

	tx.Commit()

	// Fetch complete store with photos
	var completeStore entities.Store
	err := repository.database.
		Preload("FotoToko").
		First(&completeStore, storeToCreate.ID).Error

	if err != nil {
		return entities.Store{}, err
	}

	return completeStore, nil
}

func (repository *storeRepositoryImpl) Delete(id uint) (bool, error) {
	err := repository.database.Delete(&entities.Store{}, id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
