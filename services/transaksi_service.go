package services

import (
	"errors"
	"fmt"
	"mini-project-evermos/models"
	"mini-project-evermos/models/responder"
	"mini-project-evermos/repositories"
	"time"
)

type TransactionService interface {
	GetAll(limit int, page int, keyword string) (responder.Pagination, error)
	GetById(id uint, user_id uint) (models.TransactionResponse, error)
	Create(input models.TransactionRequest, user_id uint) (models.TransactionResponse, error)
	Update(id uint, user_id uint, input models.TransactionUpdateRequest) (models.TransactionResponse, error)
	Delete(id uint, user_id uint) error
}

type transactionServiceImpl struct {
	repository        repositories.TransactionRepository
	repositoryProduct repositories.ProductRepository
	repositoryAddress repositories.AddressRepository
	userRepository    repositories.UserRepository
	regionService     RegionService
	productLogRepo    repositories.ProductLogRepository
}

func NewTransactionService(
	transactionRepository *repositories.TransactionRepository,
	productRepository *repositories.ProductRepository,
	addressRepository *repositories.AddressRepository,
	userRepository *repositories.UserRepository,
	regionService *RegionService,
	productLogRepo *repositories.ProductLogRepository,
) TransactionService {
	return &transactionServiceImpl{
		repository:        *transactionRepository,
		repositoryProduct: *productRepository,
		repositoryAddress: *addressRepository,
		userRepository:    *userRepository,
		regionService:     *regionService,
		productLogRepo:    *productLogRepo,
	}
}

func (service *transactionServiceImpl) Create(input models.TransactionRequest, user_id uint) (models.TransactionResponse, error) {
	// Create transaction process data
	transactionProcess := models.TransactionProcessData{
		Transaction: models.Transaction{
			UserID:           input.UserID,
			AlamatPengiriman: input.AlamatPengiriman,
			HargaTotal:       int(input.HargaTotal),
			KodeInvoice:      fmt.Sprintf("INV-%d-%s", input.UserID, time.Now().Format("20060102150405")),
			MethodBayar:      input.MethodBayar,
		},
	}

	// Get product details and create log entries for each product
	var logProducts []models.ProductLogProcess
	for _, item := range input.Products {
		product, err := service.repositoryProduct.FindById(item.ProductID)
		if err != nil {
			return models.TransactionResponse{}, fmt.Errorf("failed to get product details: %v", err)
		}

		// Create product log entry
		logProduct := models.ProductLogProcess{
			ProductID:     product.ID,
			NamaProduk:    product.NamaProduk,
			Slug:          product.Slug,
			HargaReseller: fmt.Sprintf("%.2f", product.HargaReseller),
			HargaKonsumen: fmt.Sprintf("%.2f", product.HargaKonsumen),
			Deskripsi:     *product.Deskripsi,
			StoreID:       product.Store.ID,
			CategoryID:    product.Category.ID,
			Kuantitas:     item.Quantity,
			HargaTotal:    item.Price * float64(item.Quantity),
		}
		logProducts = append(logProducts, logProduct)
	}

	// Add log products to transaction process
	transactionProcess.LogProduct = logProducts

	// Create transaction and process logs
	trxID, err := service.repository.Insert(transactionProcess)
	if err != nil {
		return models.TransactionResponse{}, err
	}

	// Get complete transaction data
	transaction, err := service.repository.FindById(trxID)
	if err != nil {
		return models.TransactionResponse{}, err
	}

	// Get address details
	address, err := service.repositoryAddress.FindById(input.AlamatPengiriman)
	if err != nil {
		return models.TransactionResponse{}, fmt.Errorf("failed to get address details: %v", err)
	}

	// Handle nullable timestamps
	var createdAt, updatedAt time.Time
	if transaction.CreatedAt != nil {
		createdAt = *transaction.CreatedAt
	}
	if transaction.UpdatedAt != nil {
		updatedAt = *transaction.UpdatedAt
	}

	// Return response
	return models.TransactionResponse{
		ID:          transaction.ID,
		UserID:      input.UserID,
		HargaTotal:  transaction.HargaTotal,
		KodeInvoice: transaction.KodeInvoice,
		MethodBayar: transaction.MethodBayar,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		Address: models.AddressResponse{
			ID:           address.ID,
			IDUser:       address.IDUser,
			JudulAlamat:  address.JudulAlamat,
			NamaPenerima: address.NamaPenerima,
			NoTelp:       address.NoTelp,
			DetailAlamat: address.DetailAlamat,
			Province: models.LocationResponse{
				ID:   address.IDProvinsi,
				Name: service.getProvinceName(address.IDProvinsi),
			},
			City: models.CityResponse{
				ID:         address.IDKota,
				ProvinceID: address.IDProvinsi,
				Name:       service.getCityName(address.IDKota),
			},
			CreatedAt: address.CreatedAt,
			UpdatedAt: address.UpdatedAt,
		},
	}, nil
}

func (service *transactionServiceImpl) getProvinceName(id string) string {
	province, err := service.regionService.GetProvince(id)
	if err != nil {
		return ""
	}
	return province.Name
}

func (service *transactionServiceImpl) getCityName(id string) string {
	city, err := service.regionService.GetCity(id)
	if err != nil {
		return ""
	}
	return city.Name
}

func (service *transactionServiceImpl) GetById(id uint, user_id uint) (models.TransactionResponse, error) {
	transaction, err := service.repository.FindById(id)
	if err != nil {
		return models.TransactionResponse{}, err
	}

	provinceData, err := service.regionService.GetProvince(transaction.Address.IDProvinsi)
	if err != nil {
		return models.TransactionResponse{}, fmt.Errorf("failed to get province data: %v", err)
	}

	cityData, err := service.regionService.GetCity(transaction.Address.IDKota)
	if err != nil {
		return models.TransactionResponse{}, fmt.Errorf("failed to get city data: %v", err)
	}

	// Handle nullable timestamps
	var createdAt, updatedAt time.Time
	if transaction.CreatedAt != nil {
		createdAt = *transaction.CreatedAt
	}
	if transaction.UpdatedAt != nil {
		updatedAt = *transaction.UpdatedAt
	}

	response := models.TransactionResponse{
		ID:          transaction.ID,
		UserID:      transaction.IDUser,
		HargaTotal:  transaction.HargaTotal,
		KodeInvoice: transaction.KodeInvoice,
		MethodBayar: transaction.MethodBayar,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		Address: models.AddressResponse{
			ID:           transaction.Address.ID,
			IDUser:       transaction.IDUser,
			JudulAlamat:  transaction.Address.JudulAlamat,
			NamaPenerima: transaction.Address.NamaPenerima,
			NoTelp:       transaction.Address.NoTelp,
			DetailAlamat: transaction.Address.DetailAlamat,
			Province: models.LocationResponse{
				ID:   provinceData.ID,
				Name: provinceData.Name,
			},
			City: models.CityResponse{
				ID:         cityData.ID,
				ProvinceID: cityData.ProvinceID,
				Name:       cityData.Name,
			},
			CreatedAt: transaction.Address.CreatedAt,
			UpdatedAt: transaction.Address.UpdatedAt,
		},
	}
	return response, nil
}

func (service *transactionServiceImpl) GetAll(limit int, page int, keyword string) (responder.Pagination, error) {
	request := responder.Pagination{}
	request.Limit = limit
	request.Page = page
	request.Keyword = keyword

	pagination, err := service.repository.FindAllPagination(request)
	if err != nil {
		return responder.Pagination{}, err
	}

	if responses, ok := pagination.Rows.([]models.TransactionResponse); ok {
		for i := range responses {
			if provinceData, err := service.regionService.GetProvince(responses[i].Address.Province.ID); err == nil {
				responses[i].Address.Province.Name = provinceData.Name
			}
			if cityData, err := service.regionService.GetCity(responses[i].Address.City.ID); err == nil {
				responses[i].Address.City.Name = cityData.Name
			}
		}
		pagination.Rows = responses
	}

	return pagination, nil
}

func (service *transactionServiceImpl) Update(id uint, user_id uint, input models.TransactionUpdateRequest) (models.TransactionResponse, error) {
	// Get existing transaction
	transaction, err := service.repository.FindById(id)
	if err != nil {
		return models.TransactionResponse{}, err
	}

	// Check if user is admin or owns the transaction
	user, err := service.userRepository.FindById(user_id)
	if err != nil {
		return models.TransactionResponse{}, err
	}
	if !user.IsAdmin && transaction.IDUser != user_id {
		return models.TransactionResponse{}, errors.New("forbidden")
	}

	// Update transaction details
	transaction.MethodBayar = input.MethodBayar

	// Save updated transaction
	updated, err := service.repository.Update(transaction)
	if err != nil {
		return models.TransactionResponse{}, err
	}

	// Get updated address details
	provinceData, err := service.regionService.GetProvince(updated.Address.IDProvinsi)
	if err != nil {
		return models.TransactionResponse{}, fmt.Errorf("failed to get province: %v", err)
	}
	cityData, err := service.regionService.GetCity(updated.Address.IDKota)
	if err != nil {
		return models.TransactionResponse{}, fmt.Errorf("failed to get city: %v", err)
	}

	// Handle nullable timestamps
	var createdAt, updatedAt time.Time
	if updated.CreatedAt != nil {
		createdAt = *updated.CreatedAt
	}
	if updated.UpdatedAt != nil {
		updatedAt = *updated.UpdatedAt
	}

	// Return updated response
	response := models.TransactionResponse{
		ID:          updated.ID,
		UserID:      updated.IDUser,
		HargaTotal:  updated.HargaTotal,
		KodeInvoice: updated.KodeInvoice,
		MethodBayar: updated.MethodBayar,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		Address: models.AddressResponse{
			ID:           updated.Address.ID,
			IDUser:       updated.IDUser,
			JudulAlamat:  updated.Address.JudulAlamat,
			NamaPenerima: updated.Address.NamaPenerima,
			NoTelp:       updated.Address.NoTelp,
			DetailAlamat: updated.Address.DetailAlamat,
			Province: models.LocationResponse{
				ID:   provinceData.ID,
				Name: provinceData.Name,
			},
			City: models.CityResponse{
				ID:         cityData.ID,
				ProvinceID: cityData.ProvinceID,
				Name:       cityData.Name,
			},
			CreatedAt: updated.Address.CreatedAt,
			UpdatedAt: updated.Address.UpdatedAt,
		},
	}

	return response, nil
}

func (service *transactionServiceImpl) Delete(id uint, user_id uint) error {
	// Get transaction data before deletion
	transaction, err := service.repository.FindById(id)
	if err != nil {
		return err
	}

	// Check if user is admin
	user, err := service.userRepository.FindById(user_id)
	if err != nil {
		return err
	}

	// Allow access if user is admin or owns the transaction
	if !user.IsAdmin && transaction.IDUser != user_id {
		return errors.New("forbidden")
	}

	// Delete the transaction
	return service.repository.Delete(id)
}
