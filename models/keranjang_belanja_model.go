package models

import "time"

type KeranjangBelanjaRequest struct {
	IDToko       uint `json:"id_toko" form:"id_toko"`
	IDProduk     uint `json:"id_produk" form:"id_produk"`
	JumlahProduk int  `json:"jumlah_produk" form:"jumlah_produk"`
}

type KeranjangBelanjaResponse struct {
	ID           uint `json:"id"`
	IDToko       uint `json:"id_toko"`
	IDProduk     uint `json:"id_produk"`
	JumlahProduk int  `json:"jumlah_produk"`
	Store        struct {
		ID            uint       `json:"id"`
		IDUser        uint       `json:"id_user"` // Add this line
		NamaToko      string     `json:"nama_toko"`
		DeskripsiToko string     `json:"deskripsi_toko"`
		FotoToko      []FotoToko `json:"foto_toko"`
		CreatedAt     *time.Time `json:"created_at"`
		UpdatedAt     *time.Time `json:"updated_at"`
	} `json:"toko"`
	Product struct {
		ID            uint                 `json:"id"`
		NamaProduk    string               `json:"nama_produk"`
		Slug          string               `json:"slug"`
		HargaReseller string               `json:"harga_reseler"`
		HargaKonsumen string               `json:"harga_konsumen"`
		Stok          int                  `json:"stok"`
		Deskripsi     *string              `json:"deskripsi"`
		FotoProduk    []FotoProdukResponse `json:"foto_produk"` // Add this field
		CreatedAt     *time.Time           `json:"created_at"`
		UpdatedAt     *time.Time           `json:"updated_at"`
	} `json:"produk"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FotoToko struct {
	ID        uint       `json:"id"`
	IdFoto    uint       `json:"id_foto"`
	URL       string     `json:"url"`
	Foto      string     `json:"foto"`
	CreatedAt *time.Time `json:"created_at"` // Changed back to pointer type
	UpdatedAt *time.Time `json:"updated_at"` // Changed back to pointer type
}
