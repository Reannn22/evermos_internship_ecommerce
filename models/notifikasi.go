package models

import "time"

type NotificationRequest struct {
	Pesan string `json:"pesan" binding:"required"`
}

type NotificationResponse struct {
	ID        uint       `json:"id"`
	Pesan     string     `json:"pesan"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
