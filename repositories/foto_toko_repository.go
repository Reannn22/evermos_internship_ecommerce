package repositories

import (
	"fmt"
	"mini-project-evermos/models"
	"mini-project-evermos/models/entities"

	"gorm.io/gorm"
)

type StorePhotoRepository interface {
	FindAll() ([]entities.FotoToko, error)
	FindById(id uint) (entities.FotoToko, error)
	FindByToko(idToko uint) ([]entities.FotoToko, error)
	Insert(photo entities.FotoToko) (entities.FotoToko, error)
	Update(id uint, photo entities.FotoToko) (entities.FotoToko, error)
	Delete(id uint) (entities.FotoToko, error)
	Create(input models.StorePhotoRequest) (*entities.FotoToko, error)
}

type storePhotoRepositoryImpl struct {
	database *gorm.DB
}

func NewStorePhotoRepository(database *gorm.DB) StorePhotoRepository {
	return &storePhotoRepositoryImpl{database}
}

func (repository *storePhotoRepositoryImpl) FindAll() ([]entities.FotoToko, error) {
	var photos []entities.FotoToko
	result := repository.database.Debug().Order("id desc").Find(&photos)
	if result.Error != nil {
		return nil, result.Error
	}
	return photos, nil
}

func (repository *storePhotoRepositoryImpl) FindByToko(idToko uint) ([]entities.FotoToko, error) {
	var photos []entities.FotoToko

	// Print debug info
	fmt.Printf("Debug: Finding photos for store ID = %d\n", idToko)

	// Use explicit join query to ensure we get the right photos
	query := repository.database.Debug().
		Table("foto_toko").
		Select("foto_toko.*").
		Where("foto_toko.id_toko = ?", idToko).
		Order("id desc, created_at DESC") // Modified ordering here

	result := query.Find(&photos)

	// Print debug info about found photos
	if result.Error == nil {
		fmt.Printf("Debug: Found %d photos\n", len(photos))
		for _, p := range photos {
			fmt.Printf("Debug: Photo{ID: %d, IdToko: %d, URL: %s, Photo: %s}\n",
				p.ID, p.IdToko, p.URL, p.Photo)
		}
	} else {
		fmt.Printf("Debug: Error finding photos: %v\n", result.Error)
	}

	return photos, result.Error
}

func (repository *storePhotoRepositoryImpl) FindById(id uint) (entities.FotoToko, error) {
	var photo entities.FotoToko
	err := repository.database.Debug().
		Order("id desc"). // Add ordering here
		Where("id = ?", id).
		First(&photo).Error
	return photo, err
}

func (repository *storePhotoRepositoryImpl) Insert(photo entities.FotoToko) (entities.FotoToko, error) {
	err := repository.database.Create(&photo).Error
	if err != nil {
		return entities.FotoToko{}, err
	}

	// Fetch the created photo with all fields
	var createdPhoto entities.FotoToko
	err = repository.database.First(&createdPhoto, photo.ID).Error
	if err != nil {
		return entities.FotoToko{}, err
	}

	return createdPhoto, nil
}

func (repository *storePhotoRepositoryImpl) Update(id uint, photo entities.FotoToko) (entities.FotoToko, error) {
	var existingPhoto entities.FotoToko

	// First get the existing photo
	if err := repository.database.First(&existingPhoto, id).Error; err != nil {
		return entities.FotoToko{}, err
	}

	// Update only the fields that should change
	existingPhoto.URL = photo.URL
	if photo.Photo != "" {
		existingPhoto.Photo = photo.Photo
	}
	existingPhoto.UpdatedAt = photo.UpdatedAt

	// Save the updated photo
	if err := repository.database.Save(&existingPhoto).Error; err != nil {
		return entities.FotoToko{}, err
	}

	// Fetch the updated record to ensure we have all fields
	var updatedPhoto entities.FotoToko
	err := repository.database.First(&updatedPhoto, id).Error
	if err != nil {
		return entities.FotoToko{}, err
	}

	return updatedPhoto, nil
}

func (repository *storePhotoRepositoryImpl) Delete(id uint) (entities.FotoToko, error) {
	var photo entities.FotoToko

	// First get the photo details
	err := repository.database.First(&photo, id).Error
	if err != nil {
		return entities.FotoToko{}, err
	}

	// Then delete it
	err = repository.database.Delete(&photo).Error
	if err != nil {
		return entities.FotoToko{}, err
	}

	return photo, nil
}

func (repository *storePhotoRepositoryImpl) Create(input models.StorePhotoRequest) (*entities.FotoToko, error) {
	photo := entities.FotoToko{
		IdToko: input.IdToko,
		URL:    input.URL,
		Photo:  input.Photo,
	}

	err := repository.database.Create(&photo).Error
	if err != nil {
		return nil, err
	}

	return &photo, nil
}
