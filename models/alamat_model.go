package models

import "time"

// Request
type AddressRequest struct {
	IDUser       uint   `json:"id_user"`
	JudulAlamat  string `json:"judul_alamat" binding:"required"`
	NamaPenerima string `json:"nama_penerima" binding:"required"`
	NoTelp       string `json:"no_telp" binding:"required"`
	DetailAlamat string `json:"detail_alamat" binding:"required"`
	IDProvinsi   string `json:"id_provinsi" binding:"required"`
	IDKota       string `json:"id_kota" binding:"required"`
}

// Response
type AddressResponse struct {
	ID           uint             `json:"id"`
	IDUser       uint             `json:"id_user"`
	JudulAlamat  string           `json:"judul_alamat"`
	NamaPenerima string           `json:"nama_penerima"`
	NoTelp       string           `json:"no_telp"`
	DetailAlamat string           `json:"detail_alamat"`
	Province     LocationResponse `json:"id_provinsi"`
	City         CityResponse     `json:"id_kota"`
	CreatedAt    time.Time        `json:"created_at"`
	UpdatedAt    time.Time        `json:"updated_at"`
}
