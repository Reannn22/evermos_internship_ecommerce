package repositories

import (
	"fmt"
	"mini-project-evermos/models"
	"mini-project-evermos/models/entities"
	"time"

	"gorm.io/gorm"
)

type TransactionDetailRepository interface {
	FindAll() ([]entities.TrxDetail, error)
	FindById(id uint) (entities.TrxDetail, error)
	FindByTrxId(trxId uint) ([]entities.TrxDetail, error)
	Create(detail models.TransactionDetailProcess) (entities.TrxDetail, error)
	Update(detail entities.TrxDetail) (entities.TrxDetail, error)
	Delete(id uint) error
	VerifyTrxExists(trxID uint) (bool, error)
	VerifyProductLogExists(logID uint) (bool, error)
	VerifyStoreExists(storeID uint) (bool, error)
	UpdateProductStatus(id uint, status string) error
}

type transactionDetailRepositoryImpl struct {
	db *gorm.DB
}

func NewTransactionDetailRepository(db *gorm.DB) TransactionDetailRepository {
	return &transactionDetailRepositoryImpl{db: db}
}

func (repo *transactionDetailRepositoryImpl) FindAll() ([]entities.TrxDetail, error) {
	var details []entities.TrxDetail
	err := repo.db.
		Preload("Transaction").
		Preload("Store", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, id_user, nama_toko, deskripsi_toko, created_at, updated_at") // Add id_user
		}).
		Preload("Store.FotoToko").
		Preload("ProductLog").
		Preload("ProductLog.Product").
		Preload("ProductLog.Product.Store").
		Preload("ProductLog.Product.Category").
		Preload("ProductLog.Product.FotoProduk").
		Preload("ProductLog.Product.Reviews").
		Preload("ProductLog.Product.Promos").
		Order("id desc").
		Find(&details).Error

	if err != nil {
		fmt.Printf("Error in FindAll: %v\n", err)
		return nil, err
	}

	// Initialize all nested structs if they're nil
	now := time.Now()
	for i := range details {
		if details[i].CreatedAt == nil {
			details[i].CreatedAt = &now
		}
		if details[i].UpdatedAt == nil {
			details[i].UpdatedAt = &now
		}
		if details[i].ProductLog.Product.Deskripsi == nil {
			empty := ""
			details[i].ProductLog.Product.Deskripsi = &empty
		}
		// Initialize empty slices if nil
		if details[i].ProductLog.Product.FotoProduk == nil {
			details[i].ProductLog.Product.FotoProduk = []entities.FotoProduk{}
		}
		if details[i].ProductLog.Product.Reviews == nil {
			details[i].ProductLog.Product.Reviews = []entities.ProductReview{}
		}
		if details[i].ProductLog.Product.Promos == nil {
			details[i].ProductLog.Product.Promos = []entities.ProductPromo{}
		}
	}

	return details, nil
}

func (repo *transactionDetailRepositoryImpl) FindById(id uint) (entities.TrxDetail, error) {
	var detail entities.TrxDetail
	err := repo.db.
		Preload("Transaction").
		Preload("Store", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, id_user, nama_toko, deskripsi_toko, created_at, updated_at") // Add id_user
		}).
		Preload("Store.FotoToko").
		Preload("ProductLog").
		Preload("ProductLog.Product").
		Preload("ProductLog.Product.FotoProduk").
		Preload("ProductLog.Product.Reviews").
		Preload("ProductLog.Product.Promos").
		Preload("ProductLog.Product.Store").
		Preload("ProductLog.Product.Category").
		First(&detail, id).Error
	return detail, err
}

func (repo *transactionDetailRepositoryImpl) FindByTrxId(trxId uint) ([]entities.TrxDetail, error) {
	var details []entities.TrxDetail
	err := repo.db.Where("id_trx = ?", trxId).
		Preload("Transaction").
		Preload("Store").
		Preload("Store.FotoToko").
		Preload("ProductLog").
		Preload("ProductLog.Product").
		Preload("ProductLog.Product.FotoProduk").
		Preload("ProductLog.Product.Reviews").
		Preload("ProductLog.Product.Promos").
		Preload("ProductLog.Product.Store").
		Preload("ProductLog.Product.Category").
		Find(&details).Error

	// Initialize nil timestamps with current time
	now := time.Now()
	for i := range details {
		if details[i].CreatedAt == nil {
			details[i].CreatedAt = &now
		}
		if details[i].UpdatedAt == nil {
			details[i].UpdatedAt = &now
		}
	}

	return details, err
}

func (repo *transactionDetailRepositoryImpl) Create(detail models.TransactionDetailProcess) (entities.TrxDetail, error) {
	newDetail := entities.TrxDetail{
		IDTrx:         detail.TrxID,
		IDLogProduk:   detail.LogProductID,
		IDToko:        detail.StoreID,
		Kuantitas:     detail.Kuantitas,
		HargaTotal:    float64(detail.HargaTotal),
		ProductStatus: "shipping...", // Add this line
	}

	err := repo.db.Create(&newDetail).Error
	if err != nil {
		return newDetail, err
	}

	return newDetail, nil
}

func (repo *transactionDetailRepositoryImpl) Update(detail entities.TrxDetail) (entities.TrxDetail, error) {
	// Remove this line that was forcing the status to "shipping..."
	// detail.ProductStatus = "shipping..."

	err := repo.db.Save(&detail).Error
	return detail, err
}

func (repo *transactionDetailRepositoryImpl) Delete(id uint) error {
	return repo.db.Delete(&entities.TrxDetail{}, id).Error
}

func (repo *transactionDetailRepositoryImpl) VerifyTrxExists(trxID uint) (bool, error) {
	var count int64
	err := repo.db.Table("trx").Where("id = ?", trxID).Count(&count).Error
	return count > 0, err
}

func (repo *transactionDetailRepositoryImpl) VerifyProductLogExists(logID uint) (bool, error) {
	var count int64
	err := repo.db.Table("log_produk").Where("id = ?", logID).Count(&count).Error
	return count > 0, err
}

func (repo *transactionDetailRepositoryImpl) VerifyStoreExists(storeID uint) (bool, error) {
	var count int64
	err := repo.db.Table("toko").Where("id = ?", storeID).Count(&count).Error
	return count > 0, err
}

func (repo *transactionDetailRepositoryImpl) UpdateProductStatus(id uint, status string) error {
	// Use the correct table name "trx_detail" and column name "product_status"
	result := repo.db.Exec("UPDATE trx_detail SET product_status = ? WHERE id = ?", status, id)
	if result.Error != nil {
		return result.Error
	}

	// Check if any row was affected
	if result.RowsAffected == 0 {
		return fmt.Errorf("no record found with id %d", id)
	}

	return nil
}
