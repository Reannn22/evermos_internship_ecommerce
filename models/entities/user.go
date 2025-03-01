package entities

import (
	"time"
)

type User struct {
	ID           uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Nama         string     `json:"nama" gorm:"type:varchar(255);not null"`
	Username     string     `json:"username" gorm:"type:varchar(255);not null;uniqueIndex"`
	KataSandi    string     `json:"kata_sandi" gorm:"column:kata_sandi;type:varchar(255);not null"`
	Notelp       string     `json:"notelp" gorm:"type:varchar(15);not null;uniqueIndex"`
	Email        string     `json:"email" gorm:"type:varchar(255);not null;uniqueIndex"`
	TanggalLahir time.Time  `json:"tanggal_lahir"`
	JenisKelamin string     `json:"jenis_kelamin"`
	Tentang      *string    `json:"tentang" gorm:"type:text"`
	Pekerjaan    string     `json:"pekerjaan"`
	IDProvinsi   string     `json:"id_provinsi"`
	IDKota       string     `json:"id_kota"`
	IsAdmin      bool       `json:"is_admin" gorm:"default:false"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at" gorm:"index"`
}

func (User) TableName() string {
	return "users"
}
