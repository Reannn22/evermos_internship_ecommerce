package models

import "time"

// StorePhotoData is used for internal data handling
type StorePhotoData struct {
	ID        uint      `json:"id"`
	URL       string    `json:"url"`
	Photo     string    `json:"photo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s StorePhotoData) ToFotoTokoResponse() FotoTokoResponse {
	return FotoTokoResponse{
		ID:        s.ID,
		URL:       s.URL,
		Photo:     s.Photo,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

// StorePhotoRequest is used for creating new photos
type StorePhotoRequest struct {
	IdToko uint   `json:"id_toko"`
	IdFoto string `json:"id_foto"`
	URL    string `json:"url"`
	Photo  string `json:"photo"`
}

// StorePhotoResponse is used for API responses
type StorePhotoResponse struct {
	ID        uint      `json:"id"`
	IdToko    uint      `json:"-"` // Hide from JSON response
	URL       string    `json:"url"`
	Photo     string    `json:"photo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
