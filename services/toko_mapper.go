package services

import (
	"mini-project-evermos/models"
	"mini-project-evermos/models/entities"
	"time"
)

// Common mapping functions used across services
func mapStorePhotosToResponse(photos []entities.StorePhoto) []models.FotoTokoResponse {
	var responses []models.FotoTokoResponse
	for _, photo := range photos {
		var createdAt, updatedAt time.Time
		if photo.CreatedAt != nil {
			createdAt = *photo.CreatedAt
		}
		if photo.UpdatedAt != nil {
			updatedAt = *photo.UpdatedAt
		}
		responses = append(responses, models.FotoTokoResponse{
			ID:        photo.ID,
			URL:       photo.URL,
			Photo:     photo.Photo,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}
	return responses
}

func mapStoreToResponse(store entities.Store) models.StoreResponse {
	return models.StoreResponse{
		ID:        store.ID,
		NamaToko:  store.NamaToko,
		FotoToko:  mapStorePhotosToResponse(store.FotoToko),
		CreatedAt: store.CreatedAt,
		UpdatedAt: store.UpdatedAt,
	}
}

func ToFotoTokoResponse(photo entities.StorePhoto) models.FotoTokoResponse {
	var createdAt, updatedAt time.Time
	if photo.CreatedAt != nil {
		createdAt = *photo.CreatedAt
	}
	if photo.UpdatedAt != nil {
		updatedAt = *photo.UpdatedAt
	}

	return models.FotoTokoResponse{
		ID:        photo.ID,
		URL:       photo.URL,
		Photo:     photo.Photo,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func convertFotoTokoToResponse(foto entities.FotoToko) models.FotoTokoResponse {
	return models.FotoTokoResponse{
		ID:        foto.ID,
		URL:       foto.URL,
		Photo:     foto.Photo,
		CreatedAt: foto.CreatedAt,
		UpdatedAt: foto.UpdatedAt,
	}
}

func createStorePhoto(store entities.Store) []models.FotoTokoResponse {
	var storePhotoResponse []models.FotoTokoResponse
	for _, foto := range store.FotoToko {
		var createdAt, updatedAt time.Time
		if foto.CreatedAt != nil {
			createdAt = *foto.CreatedAt
		}
		if foto.UpdatedAt != nil {
			updatedAt = *foto.UpdatedAt
		}

		storePhotoResponse = append(storePhotoResponse, models.FotoTokoResponse{
			ID:        foto.ID,
			URL:       foto.URL,
			Photo:     foto.Photo,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}
	return storePhotoResponse
}
