package entities

import "time"

type TrxDetail struct {
	ID            uint       `json:"id" gorm:"primaryKey"`
	IDTrx         uint       `json:"id_transaksi" gorm:"column:id_trx"`
	IDLogProduk   uint       `json:"id_log_produk"`
	IDToko        uint       `json:"id_toko"`
	Kuantitas     int        `json:"kuantitas"`
	HargaTotal    float64    `json:"harga_total"`
	ProductStatus string     `json:"product_status"` // Make sure this matches your DB column
	Store         Store      `json:"store" gorm:"foreignKey:IDToko"`
	ProductLog    ProductLog `json:"product_log" gorm:"foreignKey:IDLogProduk"`
	Transaction   Trx        `json:"transaction" gorm:"foreignKey:IDTrx"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
}

func (TrxDetail) TableName() string {
	return "trx_detail"
}

type TransactionDetail struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	IDTrx       uint       `gorm:"not null" json:"id_transaksi"`
	IDLogProduk uint       `gorm:"not null" json:"id_log_produk"`
	IDToko      uint       `gorm:"not null" json:"id_toko"`
	Kuantitas   int        `gorm:"not null" json:"kuantitas"`
	HargaTotal  string     `gorm:"size:255;not null" json:"harga_total"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	Store       Store      `gorm:"foreignKey:IDToko"`
	ProductLog  ProductLog `gorm:"foreignKey:IDLogProduk"`
}

func (TransactionDetail) TableName() string {
	return "detail_transaksi"
}
