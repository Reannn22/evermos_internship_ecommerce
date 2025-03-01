package repositories

import (
	"fmt"
	"mini-project-evermos/models"
	"mini-project-evermos/models/entities"
	"mini-project-evermos/models/responder"
	"time"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindAllPagination(pagination responder.Pagination) (responder.Pagination, error)
	FindById(id uint) (entities.Trx, error)
	Insert(transaction models.TransactionProcessData) (uint, error)
	Update(transaction entities.Trx) (entities.Trx, error)
	Delete(id uint) error
}

type transactionRepositoryImpl struct {
	database *gorm.DB
}

func NewTransactionRepository(database *gorm.DB) TransactionRepository {
	return &transactionRepositoryImpl{database}
}

func (repository *transactionRepositoryImpl) FindAllPagination(pagination responder.Pagination) (responder.Pagination, error) {
	var transactions []entities.Trx
	var totalRows int64

	query := repository.database.Model(&entities.Trx{})
	query.Count(&totalRows)

	err := query.
		Order("id desc").
		Preload("Address").
		Select("trx.*, trx.kode_invoice as kode_invoice").
		Limit(pagination.Limit).
		Offset(pagination.GetOffset()).
		Find(&transactions).Error

	if err != nil {
		return responder.Pagination{}, err
	}

	// Remove RegionService usage, just populate basic data
	var responses []models.TransactionResponse
	for _, trx := range transactions {
		// Handle nullable timestamps
		var createdAt, updatedAt time.Time
		if trx.CreatedAt != nil {
			createdAt = *trx.CreatedAt
		}
		if trx.UpdatedAt != nil {
			updatedAt = *trx.UpdatedAt
		}

		response := models.TransactionResponse{
			ID:          trx.ID,
			UserID:      trx.IDUser,
			HargaTotal:  trx.HargaTotal,
			KodeInvoice: trx.KodeInvoice,
			MethodBayar: trx.MethodBayar,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			Address: models.AddressResponse{
				ID:           trx.Address.ID,
				IDUser:       trx.IDUser,
				JudulAlamat:  trx.Address.JudulAlamat,
				NamaPenerima: trx.Address.NamaPenerima,
				NoTelp:       trx.Address.NoTelp,
				DetailAlamat: trx.Address.DetailAlamat,
				Province: models.LocationResponse{
					ID:   trx.Address.IDProvinsi,
					Name: "", // Will be populated by service layer
				},
				City: models.CityResponse{
					ID:         trx.Address.IDKota,
					ProvinceID: trx.Address.IDProvinsi,
					Name:       "", // Will be populated by service layer
				},
				CreatedAt: trx.Address.CreatedAt,
				UpdatedAt: trx.Address.UpdatedAt,
			},
		}
		responses = append(responses, response)
	}

	pagination.Rows = responses
	pagination.TotalRows = totalRows
	pagination.TotalPages = int((totalRows + int64(pagination.Limit) - 1) / int64(pagination.Limit))

	return pagination, nil
}

func (repository *transactionRepositoryImpl) FindById(id uint) (entities.Trx, error) {
	var transaction entities.Trx

	err := repository.database.
		Order("id desc").   // Add ordering here
		Preload("Address"). // Make sure we're preloading the Address
		Where("id = ?", id).
		First(&transaction).Error

	// Map the IDUser to the response
	if err == nil {
		transaction.IDUser = transaction.IDUser // Ensure IDUser is properly set
	}

	fmt.Printf("Retrieved transaction with address: %+v\n", transaction)

	return transaction, err
}

func (repository *transactionRepositoryImpl) Insert(transaction models.TransactionProcessData) (uint, error) {
	tx := repository.database.Begin()

	fmt.Printf("Repository: Creating transaction with UserID: %d\n", transaction.Transaction.UserID)

	transaction_insert := &entities.Trx{
		IDUser:           transaction.Transaction.UserID, // Make sure this matches your DB column
		AlamatPengiriman: transaction.Transaction.AlamatPengiriman,
		HargaTotal:       float64(transaction.Transaction.HargaTotal),
		KodeInvoice:      transaction.Transaction.KodeInvoice,
		MethodBayar:      transaction.Transaction.MethodBayar,
	}

	// Debug the entity before saving
	fmt.Printf("Repository: Transaction entity before create: %+v\n", transaction_insert)

	if err := tx.Create(transaction_insert).Error; err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to create transaction: %w", err)
	}

	// Immediately load the address after creation
	if err := tx.Model(&transaction_insert).Preload("Address").First(&transaction_insert, transaction_insert.ID).Error; err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to load address: %w", err)
	}

	for _, v := range transaction.LogProduct {
		log_product := &entities.ProductLog{
			IDProduk:      v.ProductID,
			NamaProduk:    v.NamaProduk,
			Slug:          v.Slug,
			HargaReseller: fmt.Sprintf("%.2f", v.HargaReseller),
			HargaKonsumen: fmt.Sprintf("%.2f", v.HargaKonsumen),
			Deskripsi:     &v.Deskripsi,
			IDToko:        v.StoreID,
			IDCategory:    v.CategoryID,
		}
		if err := tx.Create(log_product).Error; err != nil {
			tx.Rollback()
			return 0, err
		}

		if err := tx.Create(&entities.TrxDetail{
			IDTrx:       transaction_insert.ID,
			IDLogProduk: log_product.ID,
			IDToko:      v.StoreID,
			Kuantitas:   v.Kuantitas,
			HargaTotal:  float64(v.HargaTotal),
		}).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return transaction_insert.ID, nil
}

func (repository *transactionRepositoryImpl) Update(transaction entities.Trx) (entities.Trx, error) {
	err := repository.database.Save(&transaction).Error
	return transaction, err
}

func (repository *transactionRepositoryImpl) Delete(id uint) error {
	tx := repository.database.Begin()

	// First delete related records in trx_detail
	if err := tx.Where("id_trx = ?", id).Delete(&entities.TrxDetail{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Then delete the transaction
	if err := tx.Delete(&entities.Trx{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *transactionRepositoryImpl) Create(input models.TransactionRequest) (entities.Trx, error) {
	tx := r.database.Begin()

	// Verify if the address exists
	var addressCount int64
	if err := tx.Model(&entities.Address{}).Where("id = ?", input.AlamatPengiriman).Count(&addressCount).Error; err != nil {
		tx.Rollback()
		return entities.Trx{}, err
	}

	if addressCount == 0 {
		tx.Rollback()
		return entities.Trx{}, fmt.Errorf("address with ID %d not found", input.AlamatPengiriman)
	}

	// Verify if the user exists
	var userCount int64
	if err := tx.Table("user").Where("id = ?", input.UserID).Count(&userCount).Error; err != nil {
		tx.Rollback()
		return entities.Trx{}, err
	}

	if userCount == 0 {
		tx.Rollback()
		return entities.Trx{}, fmt.Errorf("user with ID %d not found", input.UserID)
	}

	// Generate invoice code with timestamp
	currentTime := time.Now()
	kodeInvoice := fmt.Sprintf("INV%s", currentTime.Format("2006-01-02-15-04-05"))

	trx := entities.Trx{
		IDUser:           input.UserID,
		AlamatPengiriman: input.AlamatPengiriman,
		HargaTotal:       input.HargaTotal,
		KodeInvoice:      kodeInvoice,
		MethodBayar:      input.MethodBayar,
	}

	if err := tx.Create(&trx).Error; err != nil {
		tx.Rollback()
		return entities.Trx{}, err
	}

	// Load the complete transaction with address
	if err := tx.Preload("Address").First(&trx, trx.ID).Error; err != nil {
		tx.Rollback()
		return entities.Trx{}, err
	}

	tx.Commit()
	return trx, nil
}
