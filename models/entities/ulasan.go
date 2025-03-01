package entities

import "time"

type ProductReview struct {
	ID        uint64     `json:"id" gorm:"primaryKey;autoIncrement;type:bigint unsigned"`
	IDToko    uint64     `json:"id_toko" gorm:"column:id_toko;not null;type:bigint unsigned"`
	IDProduk  uint       `json:"id_produk" gorm:"column:id_produk;not null;type:int unsigned"`
	Ulasan    string     `json:"ulasan" gorm:"column:ulasan;type:text;not null"`
	Rating    int        `json:"rating" gorm:"column:rating;type:int;not null"`
	Store     Store      `json:"store" gorm:"foreignKey:IDToko;references:ID"`
	Product   Product    `json:"product" gorm:"foreignKey:IDProduk;references:ID"`
	CreatedAt *time.Time `json:"created_at" gorm:"<-:create"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"<-:create;autoUpdateTime"`
}

func (ProductReview) TableName() string {
	return "product_reviews"
}
