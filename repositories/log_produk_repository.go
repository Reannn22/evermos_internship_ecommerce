package repositories

import (
	"mini-project-evermos/models"
	"mini-project-evermos/models/entities"

	"gorm.io/gorm"
)

type ProductLogRepository interface {
	Insert(input models.ProductLogProcess) (entities.ProductLog, error)
	FindAll() ([]entities.ProductLog, error)
	FindById(id uint) (entities.ProductLog, error)
	Update(id uint, input models.ProductLogProcess) (entities.ProductLog, error)
	Delete(id uint) error
}

type productLogRepositoryImpl struct {
	db *gorm.DB
}

func NewProductLogRepository(db *gorm.DB) ProductLogRepository {
	return &productLogRepositoryImpl{db}
}

func (repository *productLogRepositoryImpl) Insert(input models.ProductLogProcess) (entities.ProductLog, error) {
	productLog := entities.ProductLog{
		IDProduk:      input.ProductID,
		NamaProduk:    input.NamaProduk,
		Slug:          input.Slug,
		HargaReseller: input.HargaReseller, // Use string directly
		HargaKonsumen: input.HargaKonsumen, // Use string directly
		Deskripsi:     &input.Deskripsi,
		IDToko:        input.StoreID,
		IDCategory:    input.CategoryID,
	}

	err := repository.db.
		Preload("Store.FotoToko"). // Add this line
		Preload("Category").
		Preload("Product").
		Preload("Product.FotoProduk").
		Create(&productLog).Error
	if err != nil {
		return entities.ProductLog{}, err
	}

	// Fetch the complete data with relationships after creation
	var completeLog entities.ProductLog
	err = repository.db.
		Preload("Store", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, id_user, nama_toko, deskripsi_toko, url_foto, created_at, updated_at") // Add id_user
		}).
		Preload("Store.FotoToko").
		Preload("Category").
		Preload("Product").
		Preload("Product.FotoProduk").
		First(&completeLog, productLog.ID).Error
	if err != nil {
		return entities.ProductLog{}, err
	}

	// Update UrlFoto if FotoToko exists
	if len(completeLog.Store.FotoToko) > 0 {
		completeLog.Store.UrlFoto = completeLog.Store.FotoToko[0].URL
	}

	return completeLog, nil
}

func (repository *productLogRepositoryImpl) FindAll() ([]entities.ProductLog, error) {
	var productLogs []entities.ProductLog
	err := repository.db.
		Preload("Store", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, id_user, nama_toko, deskripsi_toko, url_foto, created_at, updated_at") // Add id_user
		}).
		Preload("Store.FotoToko").
		Preload("Category").
		Preload("Product").            // Add this
		Preload("Product.FotoProduk"). // Add this
		Order("id desc").              // Add this line to order by newest first
		Find(&productLogs).Error

	if err != nil {
		return nil, err
	}

	// Set default URL for stores if needed
	for i := range productLogs {
		if productLogs[i].Store.UrlFoto == "" {
			productLogs[i].Store.UrlFoto = "h22222teeeeeetps://i2.pickpik.com/photos/997/209/75/retail-grocery-supermarket-store-preview32222.jpppppg"
		}
	}

	return productLogs, nil
}

func (repository *productLogRepositoryImpl) FindById(id uint) (entities.ProductLog, error) {
	var productLog entities.ProductLog
	err := repository.db.
		Preload("Store", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, id_user, nama_toko, deskripsi_toko, url_foto, created_at, updated_at") // Add id_user
		}).
		Preload("Store.FotoToko").
		Preload("Category").
		Preload("Product").            // Add this
		Preload("Product.FotoProduk"). // Add this
		First(&productLog, id).Error

	if err != nil {
		return productLog, err
	}

	// If you have the URL stored directly in Store, use it
	if productLog.Store.UrlFoto == "" {
		// Set a default URL if none exists
		productLog.Store.UrlFoto = "h22222teeeeeetps://i2.pickpik.com/photos/997/209/75/retail-grocery-supermarket-store-preview32222.jpppppg"
	}

	return productLog, nil
}

func (repository *productLogRepositoryImpl) Update(id uint, input models.ProductLogProcess) (entities.ProductLog, error) {
	// Try to find the record first
	var productLog entities.ProductLog
	if err := repository.db.Where("id = ?", id).First(&productLog).Error; err != nil {
		return entities.ProductLog{}, err
	}

	// Prepare update data
	updates := entities.ProductLog{
		IDProduk:      input.ProductID,
		NamaProduk:    input.NamaProduk,
		Slug:          input.Slug,
		HargaReseller: input.HargaReseller,
		HargaKonsumen: input.HargaKonsumen,
		Deskripsi:     &input.Deskripsi,
		IDToko:        input.StoreID,
		IDCategory:    input.CategoryID,
	}

	// Perform update
	if err := repository.db.Model(&productLog).Updates(updates).Error; err != nil {
		return entities.ProductLog{}, err
	}

	// Fetch the updated record with all relationships
	var updatedLog entities.ProductLog
	err := repository.db.
		Preload("Store", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, id_user, nama_toko, deskripsi_toko, url_foto, created_at, updated_at") // Add id_user
		}).
		Preload("Store.FotoToko").
		Preload("Category").
		Preload("Product").            // Add this
		Preload("Product.FotoProduk"). // Add this
		First(&updatedLog, id).Error

	if err != nil {
		return entities.ProductLog{}, err
	}

	// Set default URL for store if needed
	if updatedLog.Store.UrlFoto == "" {
		updatedLog.Store.UrlFoto = "h22222teeeeeetps://i2.pickpik.com/photos/997/209/75/retail-grocery-supermarket-store-preview32222.jpppppg"
	}

	return updatedLog, nil
}

func (repository *productLogRepositoryImpl) Delete(id uint) error {
	result := repository.db.Table("log_produk").Where("id = ?", id).Delete(&entities.ProductLog{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
