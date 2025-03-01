package models

import (
	"mime/multipart"
	"time"
)

// Request type for photos
// Update the FotoProdukRequest struct to match the incoming request field names
type FotoProdukRequest struct {
	ProductID uint                  `form:"id_produk"`
	PhotoURL  string                `form:"photo_url"`
	File      *multipart.FileHeader `form:"file"`
}

// Response types for photos
type FotoProdukResponse struct {
	ID        uint      `json:"id"`
	Photo     string    `json:"photo"` // Moved up
	URL       string    `json:"url"`   // Moved down
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FotoTokoResponse struct {
	ID        uint      `json:"id"`
	URL       string    `json:"url"`
	Photo     string    `json:"photo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FileUploadResponse struct {
	URL      string `json:"url"`
	PhotoID  uint   `json:"photo_id"`
	Filename string `json:"filename"`
}
