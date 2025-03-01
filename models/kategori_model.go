package models

import "time"

// Request
type CategoryRequest struct {
	NamaCategory string `json:"nama_category" binding:"required"`
}

// Response
type CategoryResponse struct {
	ID           uint       `json:"id"`
	NamaCategory string     `json:"nama_category"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}
