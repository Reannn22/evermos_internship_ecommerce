package services

import (
	"math"
	"mime/multipart"
	"mini-project-evermos/models"
	"mini-project-evermos/models/entities"
	"mini-project-evermos/models/responder"
	"mini-project-evermos/repositories"
	"strconv"
	"time"
)

// Add this function at the beginning of the file
func convertToFotoTokoResponses(photoData []models.StorePhotoData) []models.FotoTokoResponse {
	responses := make([]models.FotoTokoResponse, len(photoData))
	for i, photo := range photoData {
		responses[i] = photo.ToFotoTokoResponse()
	}
	return responses
}

// Add this helper function near the top
func convertStorePhotoToFotoToko(photos []entities.StorePhoto) []entities.FotoToko {
	result := make([]entities.FotoToko, len(photos))
	for i, p := range photos {
		createdAt := time.Now()
		updatedAt := time.Now()
		if p.CreatedAt != nil {
			createdAt = *p.CreatedAt
		}
		if p.UpdatedAt != nil {
			updatedAt = *p.UpdatedAt
		}

		result[i] = entities.FotoToko{
			ID:        p.ID,
			IdToko:    p.IdToko,
			URL:       p.URL,
			Photo:     p.Photo,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
	}
	return result
}

// Contract
type StoreService interface {
	GetAll(limit int, page int, keyword string) (models.StorePaginationResponse, error)
	GetByUserId(id uint) (models.StoreResponse, error)
	GetById(id uint, user_id uint) (models.StoreResponse, error)
	Create(input models.StoreProcess) (models.StoreResponse, error)
	Edit(input models.StoreProcess) (models.StoreResponse, error)                   // Changed return type
	Delete(id uint, user_id uint) (models.StoreResponse, error)                     // Add this line
	CreatePhoto(input models.StorePhotoRequest) (*models.StorePhotoResponse, error) // Add this line
}

type storeServiceImpl struct {
	repository           repositories.StoreRepository
	storePhotoRepository repositories.StorePhotoRepository // Changed type
}

// Update constructor to accept correct types
func NewStoreService(storeRepository *repositories.StoreRepository, storePhotoRepository *repositories.StorePhotoRepository) StoreService {
	return &storeServiceImpl{
		repository:           *storeRepository,
		storePhotoRepository: *storePhotoRepository,
	}
}

func (service *storeServiceImpl) GetAll(limit, page int, keyword string) (models.StorePaginationResponse, error) {
	request := responder.Pagination{}
	request.Limit = limit
	request.Page = page
	request.Keyword = keyword

	// Get stores and total count from repository
	stores, total, err := service.repository.FindAllPagination(request)
	if err != nil {
		return models.StorePaginationResponse{}, err
	}

	var responses []models.StoreDetailResponse
	for _, store := range stores {
		fotoTokos := convertStorePhotoToFotoToko(store.FotoToko)
		photoResponses := service.convertPhotosToResponse(fotoTokos)
		photoData := convertToStorePhotoData(photoResponses)

		responses = append(responses, models.StoreDetailResponse{
			ID:            store.ID,
			IDUser:        store.IDUser, // Add this line
			NamaToko:      store.NamaToko,
			DeskripsiToko: store.DeskripsiToko,
			FotoToko:      photoData, // Changed this line
			CreatedAt:     store.CreatedAt,
			UpdatedAt:     store.UpdatedAt,
		})
	}

	return models.StorePaginationResponse{
		Limit:      limit,
		Page:       page,
		TotalRows:  total,
		TotalPages: int(math.Ceil(float64(total) / float64(limit))),
		Rows:       responses,
		Keyword:    keyword,
	}, nil
}

func (service *storeServiceImpl) GetByUserId(user_id uint) (models.StoreResponse, error) {
	store, err := service.repository.FindByUserId(user_id)
	if err != nil {
		return models.StoreResponse{}, err
	}

	response := models.StoreResponse{
		ID:            store.ID,
		NamaToko:      store.NamaToko,
		DeskripsiToko: store.DeskripsiToko, // Add this field
		CreatedAt:     store.CreatedAt,
		UpdatedAt:     store.UpdatedAt,
	}

	return response, nil
}

// Update the GetById method to include IDUser
func (service *storeServiceImpl) GetById(id uint, user_id uint) (models.StoreResponse, error) {
	// Get store with its photos through the relationship
	store, photos, err := service.repository.FindById(id)
	if err != nil {
		return models.StoreResponse{}, err
	}

	// Convert photos directly from the store's photos
	fotoTokos := convertStorePhotoToFotoToko(photos)
	photoResponses := service.convertPhotosToResponse(fotoTokos)
	photoData := convertToStorePhotoData(photoResponses)
	fotoResponses := convertToFotoTokoResponses(photoData)

	response := models.StoreResponse{
		ID:            store.ID,
		IDUser:        store.IDUser, // Add this line to include the IDUser
		NamaToko:      store.NamaToko,
		DeskripsiToko: store.DeskripsiToko,
		FotoToko:      fotoResponses,
		CreatedAt:     store.CreatedAt,
		UpdatedAt:     store.UpdatedAt,
	}

	return response, nil
}

// Update Create method to include DeskripsiToko in the response
func (service *storeServiceImpl) Create(input models.StoreProcess) (models.StoreResponse, error) {
	store := entities.Store{
		IDUser:        input.UserID,
		NamaToko:      input.NamaToko,
		DeskripsiToko: input.DeskripsiToko, // Add this field
	}

	// Create photo URLs array with the correct id_foto
	photoURLs := []interface{}{}
	if len(input.PhotoURLs) > 0 {
		// Use the photoURLs from input directly
		photoURLs = input.PhotoURLs
	} else if input.URL != "" {
		photoURLs = append(photoURLs, map[string]string{
			"url":          input.URL,
			"originalName": input.Photo,
			"id_foto":      strconv.FormatUint(uint64(input.IdFoto), 10),
		})
	}

	result, err := service.repository.Insert(store, photoURLs)
	if err != nil {
		return models.StoreResponse{}, err
	}

	// Get latest photos from store photo repository
	latestPhotos, err := service.storePhotoRepository.FindByToko(result.ID)
	if err != nil {
		return models.StoreResponse{}, err
	}

	photoResponses := service.convertPhotosToResponse(latestPhotos)
	photoData := convertToStorePhotoData(photoResponses)
	fotoResponses := convertToFotoTokoResponses(photoData)

	return models.StoreResponse{
		ID:            result.ID,
		IDUser:        result.IDUser, // Add this line
		NamaToko:      result.NamaToko,
		DeskripsiToko: result.DeskripsiToko,
		FotoToko:      fotoResponses,
		CreatedAt:     result.CreatedAt,
		UpdatedAt:     result.UpdatedAt,
	}, nil
}

// Update Edit method to include IDUser in response and check permissions properly
func (service *storeServiceImpl) Edit(input models.StoreProcess) (models.StoreResponse, error) {
	// Only get the store at first, we don't need photos yet
	store, _, err := service.repository.FindById(input.ID)
	if err != nil {
		return models.StoreResponse{}, err
	}

	// Use store.IDUser if needed
	_ = store.IDUser // This acknowledges we received the value even if we don't use it

	date_now := time.Now()
	string_date := date_now.Format("2006_01_02_15_04_05")
	filename := string_date + "-" + input.URL

	req := entities.Store{
		NamaToko:      input.NamaToko,
		DeskripsiToko: input.DeskripsiToko,
		UrlFoto:       filename,
	}

	photoURLs := []interface{}{}
	if input.URL != "" {
		photoURLs = append(photoURLs, map[string]string{
			"url":          input.URL,
			"originalName": input.Photo,
		})
	}

	success, err := service.repository.Update(input.ID, req, photoURLs)
	if err != nil || !success {
		return models.StoreResponse{}, err
	}

	// Get updated store with photos
	updated_store, updated_photos, err := service.repository.FindById(input.ID)
	if err != nil {
		return models.StoreResponse{}, err
	}

	updated_fotoTokos := convertStorePhotoToFotoToko(updated_photos)
	photoResponses := service.convertPhotosToResponse(updated_fotoTokos)
	photoData := convertToStorePhotoData(photoResponses)
	fotoResponses := convertToFotoTokoResponses(photoData)

	return models.StoreResponse{
		ID:            updated_store.ID,
		IDUser:        updated_store.IDUser, // Add this line to include IDUser
		NamaToko:      updated_store.NamaToko,
		DeskripsiToko: updated_store.DeskripsiToko,
		FotoToko:      fotoResponses,
		CreatedAt:     updated_store.CreatedAt,
		UpdatedAt:     updated_store.UpdatedAt,
	}, nil
}

// Update convertPhotosToResponse to handle StorePhoto input
func (service *storeServiceImpl) convertPhotosToResponse(photos []entities.FotoToko) []models.StorePhotoResponse {
	var responses []models.StorePhotoResponse
	for _, photo := range photos {
		response := models.StorePhotoResponse{
			ID:        photo.ID,
			URL:       photo.URL,
			Photo:     photo.Photo,
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
		}
		responses = append(responses, response)
	}
	return responses
}

func (service *storeServiceImpl) CreatePhoto(input models.StorePhotoRequest) (*models.StorePhotoResponse, error) {
	// Create photo in repository
	photo, err := service.storePhotoRepository.Create(input)
	if err != nil {
		return nil, err
	}

	// Convert to response
	response := &models.StorePhotoResponse{
		ID:        photo.ID,
		IdToko:    photo.IdToko,
		URL:       photo.URL,
		Photo:     photo.Photo,
		CreatedAt: photo.CreatedAt,
		UpdatedAt: photo.UpdatedAt,
	}

	return response, nil
}

func (service *storeServiceImpl) Delete(id uint, user_id uint) (models.StoreResponse, error) {
	// Get store and photos before deletion
	store, photos, err := service.repository.FindById(id)
	if err != nil {
		return models.StoreResponse{}, err
	}

	fotoTokos := convertStorePhotoToFotoToko(photos)
	photoResponses := service.convertPhotosToResponse(fotoTokos)
	photoData := convertToStorePhotoData(photoResponses)
	fotoResponses := convertToFotoTokoResponses(photoData)

	// Delete the store
	success, err := service.repository.Delete(id)
	if err != nil || !success {
		return models.StoreResponse{}, err
	}

	return models.StoreResponse{
		ID:            store.ID,
		IDUser:        store.IDUser, // Add this line to include IDUser
		NamaToko:      store.NamaToko,
		DeskripsiToko: store.DeskripsiToko,
		FotoToko:      fotoResponses,
		CreatedAt:     store.CreatedAt,
		UpdatedAt:     store.UpdatedAt,
	}, nil
}

func GenerateFilename(file *multipart.FileHeader) string {
	if file == nil {
		return ""
	}
	return file.Filename
}

func convertPhotosToFotoTokoResponses(photos []models.StorePhotoData) []models.FotoTokoResponse {
	responses := make([]models.FotoTokoResponse, len(photos))
	for i, photo := range photos {
		responses[i] = photo.ToFotoTokoResponse()
	}
	return responses
}

// Add conversion function for StorePhotoResponse to StorePhotoData
func convertToStorePhotoData(responses []models.StorePhotoResponse) []models.StorePhotoData {
	result := make([]models.StorePhotoData, len(responses))
	for i, r := range responses {
		result[i] = models.StorePhotoData{
			ID:        r.ID,
			URL:       r.URL,
			Photo:     r.Photo,
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		}
	}
	return result
}
