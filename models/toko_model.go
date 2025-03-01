package models

import (
	"mime/multipart"
	"time"
)

// Response
type StoreResponse struct {
	ID            uint               `json:"id"`
	IDUser        uint               `json:"id_user"` // Add this line
	NamaToko      string             `json:"nama_toko"`
	DeskripsiToko string             `json:"deskripsi_toko"`
	FotoToko      []FotoTokoResponse `json:"foto_toko"`
	CreatedAt     *time.Time         `json:"created_at"`
	UpdatedAt     *time.Time         `json:"updated_at"`
}

type StoreUpdate struct {
	NamaToko string
	UrlFoto  string
	IdFoto   string
}

type StoreProcess struct {
	ID            uint
	UserID        uint
	NamaToko      string
	DeskripsiToko string
	URL           string
	Photo         string
	IdFoto        uint `json:"id_foto"`
	PhotoURLs     []interface{}
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}

type File struct {
	File multipart.File `json:"file,omitempty"`
}

type StorePaginationResponse struct {
	Limit      int                   `json:"limit"`
	Page       int                   `json:"page"`
	TotalRows  int                   `json:"total_rows"`
	TotalPages int                   `json:"total_pages"`
	Rows       []StoreDetailResponse `json:"rows"`
	Keyword    string                `json:"keyword"`
}

// StoreDetailResponse is used for pagination response
type StoreDetailResponse struct {
	ID            uint             `json:"id"`
	IDUser        uint             `json:"id_user"` // Add this line
	NamaToko      string           `json:"nama_toko"`
	DeskripsiToko string           `json:"deskripsi_toko"`
	FotoToko      []StorePhotoData `json:"foto_toko"`
	CreatedAt     *time.Time       `json:"created_at"`
	UpdatedAt     *time.Time       `json:"updated_at"`
}
