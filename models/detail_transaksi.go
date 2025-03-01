package models

import "time"

type TransactionDetailProcess struct {
	TrxID        uint    `json:"trx_id"`         // Changed from "TrxID"
	LogProductID uint    `json:"log_product_id"` // Changed from "LogProductID"
	StoreID      uint    `json:"store_id"`       // Changed from "StoreID"
	Kuantitas    int     `json:"kuantitas"`
	HargaTotal   float64 `json:"harga_total"`
}

type TransactionDetailRequest struct {
	LogProductID uint    `json:"log_product_id"`
	StoreID      uint    `json:"store_id"`
	Kuantitas    int     `json:"kuantitas"`
	HargaTotal   float64 `json:"harga_total"`
}

type TransactionDetailResponse struct {
	ID            uint                  `json:"id"`
	IDTransaksi   uint                  `json:"id_transaksi"`
	IDLogProduk   uint                  `json:"id_log_produk"`
	IDToko        uint                  `json:"id_toko"`
	Kuantitas     int                   `json:"kuantitas"`
	HargaTotal    float64               `json:"harga_total"`
	ProductStatus string                `json:"product_status"` // Add this field
	Store         StoreResponse         `json:"store"`
	Produk        SimpleProductResponse `json:"produk"` // Changed from Product to Produk and using simpler structure
	CreatedAt     time.Time             `json:"created_at"`
	UpdatedAt     time.Time             `json:"updated_at"`
}

// Add new type for simpler product response
type SimpleProductResponse struct {
	ID            uint                          `json:"id"`
	NamaProduk    string                        `json:"nama_produk"`
	Slug          string                        `json:"slug"`
	HargaReseller string                        `json:"harga_reseler"`
	HargaKonsumen string                        `json:"harga_konsumen"`
	Stok          int                           `json:"stok"`
	Deskripsi     string                        `json:"deskripsi"`
	FotoProduk    []FotoProdukResponse          `json:"foto_produk"` // Add this field
	Reviews       []SimpleProductReviewResponse `json:"reviews"`     // Change type from interface{} to slice
	Promos        []ProductPromoResponse        `json:"promos"`      // Change type from interface{} to slice
	CreatedAt     *time.Time                    `json:"created_at"`
	UpdatedAt     *time.Time                    `json:"updated_at"`
}
