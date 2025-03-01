package models

import "time"

type WishlistRequest struct {
	ProductID uint `json:"product_id" form:"product_id"`
	StoreID   uint `json:"store_id" form:"store_id"`
}

type WishlistResponse struct {
	ID       uint          `json:"id"`
	IDToko   uint          `json:"id_toko"`
	IDProduk uint          `json:"id_produk"`
	Store    StoreResponse `json:"toko"`
	Product  struct {
		ID            uint                 `json:"id"`
		NamaProduk    string               `json:"nama_produk"`
		Slug          string               `json:"slug"`
		HargaReseller string               `json:"harga_reseler"`
		HargaKonsumen string               `json:"harga_konsumen"`
		Stok          int                  `json:"stok"`
		Deskripsi     *string              `json:"deskripsi"`
		FotoProduk    []FotoProdukResponse `json:"foto_produk"`
		CreatedAt     *time.Time           `json:"created_at"`
		UpdatedAt     *time.Time           `json:"updated_at"`
	} `json:"produk"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
