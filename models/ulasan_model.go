package models

import "time"

type ProductReviewRequest struct {
	Ulasan   string `json:"ulasan" form:"ulasan"`
	Rating   int    `json:"rating" form:"rating"`
	IDToko   uint   `json:"id_toko" form:"id_toko"`
	IDProduk uint   `json:"id_produk" form:"id_produk"`
}

type ProductReviewResponse struct {
	ID        uint          `json:"id"`
	IDToko    uint          `json:"id_toko"`
	IDProduk  uint          `json:"id_produk"`
	Ulasan    string        `json:"ulasan"`
	Rating    int           `json:"rating"`
	Store     StoreResponse `json:"toko"`
	Product   ProductDetail `json:"produk"`
	CreatedAt *time.Time    `json:"created_at"`
	UpdatedAt *time.Time    `json:"updated_at"`
}

type ProductDetail struct {
	ID            uint                 `json:"id"`
	NamaProduk    string               `json:"nama_produk"`
	Slug          string               `json:"slug"`
	HargaReseller string               `json:"harga_reseler"`
	HargaKonsumen string               `json:"harga_konsumen"`
	Stok          int                  `json:"stok"`
	Deskripsi     string               `json:"deskripsi"`
	FotoProduk    []FotoProdukResponse `json:"foto_produk"` // Add this field
	CreatedAt     *time.Time           `json:"created_at"`
	UpdatedAt     *time.Time           `json:"updated_at"`
}
