package services

import (
	"mini-project-evermos/models"
	"mini-project-evermos/models/entities" // Add this import
	"mini-project-evermos/repositories"
	"strconv"
)

type ProductLogService interface {
	Create(input models.ProductLogProcess) (models.ProductLogResponse, error)
	GetAll() ([]models.ProductLogResponse, error)
	GetById(id uint) (models.ProductLogResponse, error)
	Update(id uint, input models.ProductLogProcess) (models.ProductLogResponse, error)
	Delete(id uint) (models.ProductLogResponse, error)
}

type productLogServiceImpl struct {
	repository repositories.ProductLogRepository
}

func NewProductLogService(productLogRepository *repositories.ProductLogRepository) ProductLogService {
	return &productLogServiceImpl{
		repository: *productLogRepository,
	}
}

func mapStoreToProductLogResponse(store entities.Store) models.StoreResponse {
	return models.StoreResponse{
		ID:            store.ID,
		IDUser:        store.IDUser, // Add this line to include IDUser
		NamaToko:      store.NamaToko,
		DeskripsiToko: store.DeskripsiToko, // Add this field
		FotoToko:      mapStorePhotosToResponse(store.FotoToko),
		CreatedAt:     store.CreatedAt,
		UpdatedAt:     store.UpdatedAt,
	}
}

// Add this function near the other mapping functions at the top of the file
func mapCategoryToResponse(category entities.Category) models.CategoryResponse {
	return models.CategoryResponse{
		ID:           category.ID,
		NamaCategory: category.NamaCategory,
		CreatedAt:    category.CreatedAt,
		UpdatedAt:    category.UpdatedAt,
	}
}

func mapProductLogToResponse(productLog entities.ProductLog) models.ProductLogResponse {
	hargaReseller, _ := strconv.ParseFloat(productLog.HargaReseller, 64)
	hargaKonsumen, _ := strconv.ParseFloat(productLog.HargaKonsumen, 64)

	var fotoProdukResponses []models.FotoProdukResponse
	for _, foto := range productLog.Product.FotoProduk {
		fotoProdukResponses = append(fotoProdukResponses, models.FotoProdukResponse{
			ID:        foto.ID,
			Photo:     foto.Photo,
			URL:       foto.PhotoURL,
			CreatedAt: *foto.CreatedAt,
			UpdatedAt: *foto.UpdatedAt,
		})
	}

	// Handle nullable Deskripsi field
	var deskripsi string
	if productLog.Product.Deskripsi != nil {
		deskripsi = *productLog.Product.Deskripsi
	}

	return models.ProductLogResponse{
		ID:            productLog.ID,
		StoreID:       productLog.IDToko,
		ProductID:     productLog.IDProduk,
		CategoryID:    productLog.IDCategory,
		NamaProduk:    productLog.NamaProduk,
		Slug:          productLog.Slug,
		HargaReseller: hargaReseller,
		HargaKonsumen: hargaKonsumen,
		Deskripsi:     *productLog.Deskripsi,
		Store:         mapStoreToProductLogResponse(productLog.Store),
		Produk: models.ProductLogDetailResponse{
			ID:            productLog.Product.ID,
			NamaProduk:    productLog.Product.NamaProduk,
			Slug:          productLog.Product.Slug,
			HargaReseller: productLog.Product.HargaReseller,
			HargaKonsumen: productLog.Product.HargaKonsumen,
			Stok:          productLog.Product.Stok,
			Deskripsi:     deskripsi, // Use the handled deskripsi value
			FotoProduk:    fotoProdukResponses,
			CreatedAt:     productLog.Product.CreatedAt,
			UpdatedAt:     productLog.Product.UpdatedAt,
		},
		Category:  mapCategoryToResponse(productLog.Category),
		CreatedAt: *productLog.CreatedAt,
		UpdatedAt: *productLog.UpdatedAt,
	}
}

func (service *productLogServiceImpl) Create(input models.ProductLogProcess) (models.ProductLogResponse, error) {
	productLog, err := service.repository.Insert(input)
	if err != nil {
		return models.ProductLogResponse{}, err
	}

	response := mapProductLogToResponse(productLog)

	return response, nil
}

func (service *productLogServiceImpl) GetAll() ([]models.ProductLogResponse, error) {
	productLogs, err := service.repository.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []models.ProductLogResponse
	for _, log := range productLogs {
		response := mapProductLogToResponse(log)
		responses = append(responses, response)
	}

	return responses, nil
}

func (service *productLogServiceImpl) GetById(id uint) (models.ProductLogResponse, error) {
	productLog, err := service.repository.FindById(id)
	if err != nil {
		return models.ProductLogResponse{}, err
	}

	response := mapProductLogToResponse(productLog)

	return response, nil
}

func (service *productLogServiceImpl) Update(id uint, input models.ProductLogProcess) (models.ProductLogResponse, error) {
	productLog, err := service.repository.Update(id, input)
	if err != nil {
		return models.ProductLogResponse{}, err
	}

	response := mapProductLogToResponse(productLog)

	return response, nil
}

func (service *productLogServiceImpl) Delete(id uint) (models.ProductLogResponse, error) {
	// Get the product log before deletion
	productLog, err := service.repository.FindById(id)
	if err != nil {
		return models.ProductLogResponse{}, err
	}

	// Create response before deleting
	response := mapProductLogToResponse(productLog)

	// Perform deletion
	err = service.repository.Delete(id)
	if err != nil {
		return models.ProductLogResponse{}, err
	}

	return response, nil
}
