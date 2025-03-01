package entities

import "time"

type FotoProduk struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	IDProduk  uint       `gorm:"not null" json:"id_produk"`
	PhotoURL  string     `gorm:"size:255" json:"photo_url"` // For storing the URL
	URL       string     `gorm:"size:255" json:"url"`       // Also store URL
	Photo     string     `gorm:"size:255" json:"photo"`     // For storing the file path
	FileName  string     `gorm:"size:255"`
	FileType  string     `gorm:"size:50"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (FotoProduk) TableName() string {
	return "foto_produk"
}
