package models

import "time"

type ProductLogProcess struct {
	ProductID     uint    `json:"product_id"`
	NamaProduk    string  `json:"nama_produk"`
	Slug          string  `json:"slug"`
	HargaReseller string  `json:"harga_reseller"` // Changed from float64 to string
	HargaKonsumen string  `json:"harga_konsumen"` // Changed from float64 to string
	Deskripsi     string  `json:"deskripsi"`
	StoreID       uint    `json:"store_id"`
	CategoryID    uint    `json:"category_id"`
	Kuantitas     int     `json:"kuantitas"`
	HargaTotal    float64 `json:"harga_total"`
}

type ProductLogDetailResponse struct {
	ID            uint                 `json:"id"`
	NamaProduk    string               `json:"nama_produk"`
	Slug          string               `json:"slug"`
	HargaReseller string               `json:"harga_reseler"`
	HargaKonsumen string               `json:"harga_konsumen"`
	Stok          int                  `json:"stok"`
	Deskripsi     string               `json:"deskripsi"`
	FotoProduk    []FotoProdukResponse `json:"foto_produk"`
	CreatedAt     *time.Time           `json:"created_at"`
	UpdatedAt     *time.Time           `json:"updated_at"`
}

type ProductLogResponse struct {
	ID            uint                     `json:"id"`
	StoreID       uint                     `json:"store_id"`
	ProductID     uint                     `json:"product_id"`
	CategoryID    uint                     `json:"category_id"`
	NamaProduk    string                   `json:"nama_produk"`
	Slug          string                   `json:"slug"`
	HargaReseller float64                  `json:"harga_reseller"`
	HargaKonsumen float64                  `json:"harga_konsumen"`
	Deskripsi     string                   `json:"deskripsi"`
	Store         StoreResponse            `json:"store"`
	Produk        ProductLogDetailResponse `json:"produk"` // Use the new response type
	Category      CategoryResponse         `json:"category"`
	CreatedAt     time.Time                `json:"created_at"`
	UpdatedAt     time.Time                `json:"updated_at"`
}
