package models

import "time"

// Request
type RegisterRequest struct {
	Nama         string `json:"nama" binding:"required"`
	KataSandi    string `json:"kata_sandi" binding:"required"`
	NoTelp       string `json:"no_telp" binding:"required"`
	TanggalLahir string `json:"tanggal_lahir" binding:"required"`
	JenisKelamin string `json:"jenis_kelamin" binding:"required"`
	Tentang      string `json:"tentang"` // Optional field
	Pekerjaan    string `json:"pekerjaan" binding:"required"`
	Email        string `json:"email" binding:"required"`
	IDProvinsi   string `json:"id_provinsi" binding:"required"`
	IDKota       string `json:"id_kota" binding:"required"`
	IsAdmin      bool   `json:"is_admin"` // Optional field
}

type LoginRequest struct {
	Email     string `json:"email" binding:"required"`
	KataSandi string `json:"kata_sandi" binding:"required"`
}

// Remove these duplicated types
// type ProvinceDetail struct {...}
// type CityDetail struct {...}

// Use the shared types from location_model.go

// Response
type LoginResponse struct {
	ID           uint           `json:"id"`
	Nama         string         `json:"nama"`
	KataSandi    string         `json:"kata_sandi"`
	NoTelp       string         `json:"no_telp"`
	TanggalLahir string         `json:"tanggal_lahir"`
	JenisKelamin string         `json:"jenis_kelamin"`
	Tentang      *string        `json:"tentang"`
	Pekerjaan    string         `json:"pekerjaan"`
	Email        string         `json:"email"`
	IDProvinsi   ProvinceDetail `json:"id_provinsi"`
	IDKota       CityDetail     `json:"id_kota"`
	IsAdmin      bool           `json:"is_admin"`
	CreatedAt    string         `json:"created_at"`
	UpdatedAt    string         `json:"updated_at"`
	Token        string         `json:"token"`
}

type RegisterResponse struct {
	ID           uint           `json:"id"`
	Nama         string         `json:"nama"`
	KataSandi    string         `json:"kata_sandi"`
	NoTelp       string         `json:"no_telp"`       // Changed from notelp
	TanggalLahir string         `json:"tanggal_lahir"` // Changed from time.Time
	JenisKelamin string         `json:"jenis_kelamin"`
	Tentang      *string        `json:"tentang"`
	Pekerjaan    string         `json:"pekerjaan"`
	Email        string         `json:"email"`
	IDProvinsi   ProvinceDetail `json:"id_provinsi"` // Changed to ProvinceDetail
	IDKota       CityDetail     `json:"id_kota"`     // Changed to CityDetail
	IsAdmin      bool           `json:"is_admin"`
	CreatedAt    string         `json:"created_at"` // Changed from time.Time
	UpdatedAt    string         `json:"updated_at"` // Changed from time.Time
}

type LogoutResponse struct {
	ID           uint           `json:"id"`
	Nama         string         `json:"nama"`
	KataSandi    string         `json:"kata_sandi"`
	NoTelp       string         `json:"no_telp"`
	TanggalLahir string         `json:"tanggal_lahir"`
	JenisKelamin string         `json:"jenis_kelamin"`
	Tentang      *string        `json:"tentang"`
	Pekerjaan    string         `json:"pekerjaan"`
	Email        string         `json:"email"`
	IDProvinsi   ProvinceDetail `json:"id_provinsi"`
	IDKota       CityDetail     `json:"id_kota"`
	IsAdmin      bool           `json:"is_admin"`
	CreatedAt    string         `json:"created_at"`
	UpdatedAt    string         `json:"updated_at"`
}

// process mapping
type RegisterProcess struct {
	Nama         string
	NoTelp       string
	Email        string
	KataSandi    string
	TanggalLahir time.Time
	Pekerjaan    string
	IDProvinsi   string
	IDKota       string
}
