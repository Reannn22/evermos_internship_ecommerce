package models

import (
	"mime/multipart"
	"time"
)

// Request
type ProductRequest struct {
	NamaProduk    string                  `json:"nama_produk" form:"nama_produk"`
	CategoryID    uint                    `json:"category_id" form:"category_id"`
	StoreID       uint                    `json:"store_id"`
	HargaReseller string                  `json:"harga_reseller" form:"harga_reseller"`
	HargaKonsumen string                  `json:"harga_konsumen" form:"harga_konsumen"`
	Stok          int                     `json:"stok" form:"stok"`
	Deskripsi     string                  `json:"deskripsi" form:"deskripsi"`
	PhotoURLs     []interface{}           `json:"photo_urls" form:"photo_urls"` // Changed type to interface{}
	PhotoIDs      []uint                  `json:"photo_ids" form:"photo_ids"`
	PhotoFiles    []*multipart.FileHeader `form:"photo_files"`
}

// Response
type ProductResponse struct {
	ID            uint                          `json:"id"`
	NamaProduk    string                        `json:"nama_produk"`
	Slug          string                        `json:"slug"`
	HargaReseller string                        `json:"harga_reseler"`
	HargaKonsumen string                        `json:"harga_konsumen"`
	Stok          int                           `json:"stok"`
	Deskripsi     *string                       `json:"deskripsi"`
	Store         StoreResponse                 `json:"toko"`
	Category      CategoryResponse              `json:"category"`
	FotoProduk    []FotoProdukResponse          `json:"foto_produk"`
	Reviews       []SimpleProductReviewResponse `json:"reviews"`
	Promos        []ProductPromoResponse        `json:"promos"`
	Coupons       []ProductCouponResponse       `json:"coupons"` // Add this line
	CreatedAt     *time.Time                    `json:"created_at"`
	UpdatedAt     *time.Time                    `json:"updated_at"`
}

type SimpleProductReviewResponse struct {
	ID        uint       `json:"id"`
	IDToko    uint       `json:"id_toko"`
	IDProduk  uint       `json:"id_produk"`
	Ulasan    string     `json:"ulasan"`
	Rating    int        `json:"rating"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
