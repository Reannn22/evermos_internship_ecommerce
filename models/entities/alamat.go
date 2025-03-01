package entities

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	ID           uint           `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	IDUser       uint           `json:"id_user" gorm:"column:id_user;not null"`
	JudulAlamat  string         `json:"judul_alamat" gorm:"column:judul_alamat"`
	NamaPenerima string         `json:"nama_penerima" gorm:"column:nama_penerima"`
	NoTelp       string         `json:"no_telp" gorm:"column:no_telp"`
	DetailAlamat string         `json:"detail_alamat" gorm:"column:detail_alamat"`
	IDProvinsi   string         `json:"id_provinsi" gorm:"column:id_provinsi"`
	IDKota       string         `json:"id_kota" gorm:"column:id_kota"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

func (Address) TableName() string {
	return "alamat"
}
