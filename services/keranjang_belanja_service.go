package services

import (
	"mini-project-evermos/models"
	"mini-project-evermos/models/entities"
	"mini-project-evermos/repositories"
	"time" // Add this import
)

type KeranjangBelanjaService interface {
	GetAll() ([]models.KeranjangBelanjaResponse, error)
	GetById(id uint) (models.KeranjangBelanjaResponse, error)
	Create(input models.KeranjangBelanjaRequest) (models.KeranjangBelanjaResponse, error)
	Update(id uint, input models.KeranjangBelanjaRequest) (models.KeranjangBelanjaResponse, error)
	Delete(id uint) (models.KeranjangBelanjaResponse, error)
	ClearAll() ([]models.KeranjangBelanjaResponse, error)
}

type keranjangBelanjaServiceImpl struct {
	repository repositories.KeranjangBelanjaRepository
}

func NewKeranjangBelanjaService(repository *repositories.KeranjangBelanjaRepository) KeranjangBelanjaService {
	return &keranjangBelanjaServiceImpl{
		repository: *repository,
	}
}

func (service *keranjangBelanjaServiceImpl) toResponse(kb entities.KeranjangBelanja) models.KeranjangBelanjaResponse {
	response := models.KeranjangBelanjaResponse{
		ID:           kb.ID,
		IDToko:       kb.IDToko,
		IDProduk:     kb.IDProduk,
		JumlahProduk: kb.JumlahProduk,
		CreatedAt:    kb.CreatedAt,
		UpdatedAt:    kb.UpdatedAt,
	}

	// Set Store data with IDUser
	response.Store.ID = kb.Store.ID
	response.Store.IDUser = kb.Store.IDUser // Add this line
	response.Store.NamaToko = kb.Store.NamaToko
	response.Store.DeskripsiToko = kb.Store.DeskripsiToko // Add this line
	response.Store.CreatedAt = kb.Store.CreatedAt
	response.Store.UpdatedAt = kb.Store.UpdatedAt

	// Set FotoToko data
	for _, foto := range kb.Store.FotoToko {
		var createdAt, updatedAt *time.Time // Changed to pointer type
		if foto.CreatedAt != nil {
			createdAtValue := *foto.CreatedAt
			createdAt = &createdAtValue
		}
		if foto.UpdatedAt != nil {
			updatedAtValue := *foto.UpdatedAt
			updatedAt = &updatedAtValue
		}

		fotoData := models.FotoToko{
			ID:        foto.ID,
			IdFoto:    foto.IdToko,
			URL:       foto.URL,
			Foto:      foto.Photo,
			CreatedAt: createdAt, // Now using *time.Time
			UpdatedAt: updatedAt, // Now using *time.Time
		}
		response.Store.FotoToko = append(response.Store.FotoToko, fotoData)
	}

	// Map Product data with photos
	response.Product.ID = kb.Product.ID
	response.Product.NamaProduk = kb.Product.NamaProduk
	response.Product.Slug = kb.Product.Slug
	response.Product.HargaReseller = kb.Product.HargaReseller
	response.Product.HargaKonsumen = kb.Product.HargaKonsumen
	response.Product.Stok = kb.Product.Stok
	response.Product.Deskripsi = kb.Product.Deskripsi

	// Map product photos
	var fotoProdukResponses []models.FotoProdukResponse
	for _, foto := range kb.Product.FotoProduk {
		var createdAt, updatedAt time.Time
		if foto.CreatedAt != nil {
			createdAt = *foto.CreatedAt
		}
		if foto.UpdatedAt != nil {
			updatedAt = *foto.UpdatedAt
		}

		fotoProdukResponses = append(fotoProdukResponses, models.FotoProdukResponse{
			ID:        foto.ID,
			Photo:     foto.Photo,
			URL:       foto.PhotoURL,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}
	response.Product.FotoProduk = fotoProdukResponses

	response.Product.CreatedAt = kb.Product.CreatedAt
	response.Product.UpdatedAt = kb.Product.UpdatedAt

	return response
}

func mapStorePhotoToResponse(foto entities.StorePhoto) models.FotoTokoResponse {
	var createdAt, updatedAt time.Time
	if foto.CreatedAt != nil {
		createdAt = *foto.CreatedAt
	}
	if foto.UpdatedAt != nil {
		updatedAt = *foto.UpdatedAt
	}

	return models.FotoTokoResponse{
		ID:        foto.ID,
		URL:       foto.URL,
		Photo:     foto.Photo,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func (service *keranjangBelanjaServiceImpl) createFotoToko(fotos []entities.FotoToko) []models.FotoTokoResponse {
	var fotoTokoResponse []models.FotoTokoResponse
	for _, foto := range fotos {
		fotoTokoResponse = append(fotoTokoResponse, models.FotoTokoResponse{
			ID:        foto.ID,
			URL:       foto.URL,
			Photo:     foto.Photo,
			CreatedAt: foto.CreatedAt,
			UpdatedAt: foto.UpdatedAt,
		})
	}
	return fotoTokoResponse
}

func (service *keranjangBelanjaServiceImpl) GetAll() ([]models.KeranjangBelanjaResponse, error) {
	keranjangBelanja, err := service.repository.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []models.KeranjangBelanjaResponse
	for _, kb := range keranjangBelanja {
		responses = append(responses, service.toResponse(kb))
	}

	return responses, nil
}

func (service *keranjangBelanjaServiceImpl) GetById(id uint) (models.KeranjangBelanjaResponse, error) {
	keranjangBelanja, err := service.repository.FindById(id)
	if err != nil {
		return models.KeranjangBelanjaResponse{}, err
	}

	return service.toResponse(keranjangBelanja), nil
}

func (service *keranjangBelanjaServiceImpl) Create(input models.KeranjangBelanjaRequest) (models.KeranjangBelanjaResponse, error) {
	keranjangBelanja, err := service.repository.Create(input)
	if err != nil {
		return models.KeranjangBelanjaResponse{}, err
	}

	return service.GetById(keranjangBelanja.ID)
}

func (service *keranjangBelanjaServiceImpl) Update(id uint, input models.KeranjangBelanjaRequest) (models.KeranjangBelanjaResponse, error) {
	keranjangBelanja, err := service.repository.Update(id, input)
	if err != nil {
		return models.KeranjangBelanjaResponse{}, err
	}

	return service.GetById(keranjangBelanja.ID)
}

func (service *keranjangBelanjaServiceImpl) Delete(id uint) (models.KeranjangBelanjaResponse, error) {
	// First get the complete data before deletion
	response, err := service.GetById(id)
	if err != nil {
		return models.KeranjangBelanjaResponse{}, err
	}

	// Then delete the data
	_, err = service.repository.Delete(id)
	if err != nil {
		return models.KeranjangBelanjaResponse{}, err
	}

	// Return the complete data that was just deleted
	return response, nil
}

func (service *keranjangBelanjaServiceImpl) ClearAll() ([]models.KeranjangBelanjaResponse, error) {
	// Get all items before deletion
	items, err := service.GetAll()
	if err != nil {
		return nil, err
	}

	// Clear all items
	err = service.repository.ClearAll()
	if err != nil {
		return nil, err
	}

	return items, nil
}
