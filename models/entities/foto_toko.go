package entities

import "time"

type FotoToko struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	IdToko    uint      `json:"id_toko"`
	URL       string    `json:"url"`
	Photo     string    `json:"photo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName specifies the table name for GORM
func (FotoToko) TableName() string {
	return "foto_toko"
}
