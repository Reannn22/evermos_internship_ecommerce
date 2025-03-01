package entities

import "time"

type ProductPromo struct {
	ID        uint64     `json:"id" gorm:"primaryKey;autoIncrement;type:bigint unsigned"`
	IDToko    uint64     `json:"id_toko" gorm:"column:id_toko;not null;type:bigint unsigned"`
	IDProduk  uint       `json:"id_produk" gorm:"column:id_produk;not null;type:int unsigned"`
	Promo     string     `json:"promo" gorm:"column:promo;type:text;not null"`
	Store     Store      `json:"store" gorm:"foreignKey:IDToko;references:ID"`
	Product   Product    `json:"product" gorm:"foreignKey:IDProduk;references:ID"`
	CreatedAt *time.Time `json:"created_at" gorm:"<-:create"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"<-:create;autoUpdateTime"`
}

func (ProductPromo) TableName() string {
	return "product_promos"
}
