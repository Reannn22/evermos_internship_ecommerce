package entities

import "time"

type Order struct {
	ID                  uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	TransactionDetailID uint      `json:"transaction_detail_id"`
	StatusProduk        string    `json:"status_produk"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
