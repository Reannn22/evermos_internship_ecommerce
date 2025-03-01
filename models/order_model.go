package models

import "time"

type OrderRequest struct {
	TransactionDetailID uint   `json:"transaction_detail_id"`
	ProductStatus       string `json:"product_status"` // Changed back to string
}

type OrderResponse struct {
	ID           uint      `json:"id"`
	TrxDetailID  uint      `json:"id_trx_detail"`
	StatusProduk string    `json:"status_produk"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
