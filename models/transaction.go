package models

import (
	"time"
)

type Transaction struct {
	UserID           uint   `json:"user_id"`
	AlamatPengiriman uint   `json:"alamat_pengiriman"`
	HargaTotal       int    `json:"harga_total"`
	KodeInvoice      string `json:"kode_invoice"`
	MethodBayar      string `json:"method_bayar"`
}

type TransactionRequest struct {
	UserID           uint                 `json:"user_id"`
	AlamatPengiriman uint                 `json:"alamat_pengiriman"`
	HargaTotal       float64              `json:"harga_total"`
	MethodBayar      string               `json:"method_bayar"`
	Products         []TransactionProduct `json:"products"`
}

type TransactionProduct struct {
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type TransactionUpdateRequest struct {
	AlamatPengiriman uint   `json:"alamat_pengiriman"`
	MethodBayar      string `json:"method_bayar"`
}

type TransactionResponse struct {
	ID                 uint                `json:"id"`
	UserID             uint                `json:"user_id"`
	HargaTotal         float64             `json:"harga_total"`
	KodeInvoice        string              `json:"kode_invoice"`
	MethodBayar        string              `json:"method_bayar"`
	Address            AddressResponse     `json:"address"`
	TransactionDetails []TransactionDetail `json:"transaction_details,omitempty"`
	CreatedAt          time.Time           `json:"created_at"`
	UpdatedAt          time.Time           `json:"updated_at"`
}

type TransactionProcessData struct {
	Transaction Transaction         `json:"transaction"`
	LogProduct  []ProductLogProcess `json:"log_product"`
}

type TransactionDetail struct {
	ID          uint      `json:"id"`
	IDTrx       uint      `json:"id_trx"`
	IDLogProduk uint      `json:"id_log_produk"`
	IDToko      uint      `json:"id_toko"`
	Kuantitas   int       `json:"kuantitas"`
	HargaTotal  float64   `json:"harga_total"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
