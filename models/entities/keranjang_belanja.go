package entities

import (
	"time"
)

type KeranjangBelanja struct {
	ID           uint      `gorm:"primaryKey;column:id"`
	IDToko       uint      `gorm:"column:id_toko;not null"`
	IDProduk     uint      `gorm:"column:id_produk;not null"`
	JumlahProduk int       `gorm:"column:jumlah_produk;not null;default:1"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
	Store        Store     `gorm:"foreignKey:IDToko"`
	Product      Product   `gorm:"foreignKey:IDProduk"`
}

func (KeranjangBelanja) TableName() string {
	return "keranjang_belanja"
}
