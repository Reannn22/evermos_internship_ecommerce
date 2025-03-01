package entities

import (
	"time"

	"gorm.io/gorm"
)

type ProductLog struct {
	gorm.Model
	ID            uint       `json:"id" gorm:"primaryKey;column:id"`
	IDProduk      uint       `json:"id_produk" gorm:"column:id_produk;not null"`
	NamaProduk    string     `json:"nama_produk" gorm:"column:nama_produk;size:255;not null"`
	Slug          string     `json:"slug" gorm:"column:slug;size:255;not null"`
	HargaReseller string     `json:"harga_reseller" gorm:"column:harga_reseller;size:255;not null"`
	HargaKonsumen string     `json:"harga_konsumen" gorm:"column:harga_konsumen;size:255;not null"`
	Deskripsi     *string    `json:"deskripsi" gorm:"column:deskripsi;type:text;default:null"`
	IDToko        uint       `json:"id_toko" gorm:"column:id_toko;not null"`
	IDCategory    uint       `json:"id_category" gorm:"column:id_category;not null"`
	Kuantitas     int        `json:"kuantitas"`
	HargaTotal    float64    `json:"harga_total"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	Product       Product    `json:"product" gorm:"foreignKey:IDProduk"`
	Store         Store      `json:"store" gorm:"foreignKey:IDToko"`
	Category      Category   `json:"category" gorm:"foreignKey:IDCategory"`
}

func (ProductLog) TableName() string {
	return "log_produk"
}
