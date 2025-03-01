package entities

import "time"

type ProductCoupon struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	IDProduk  uint       `json:"id_produk" gorm:"column:id_produk;index:idx_product_code,unique"`
	Product   Product    `json:"product" gorm:"foreignKey:IDProduk"`
	Code      string     `json:"code" gorm:"column:kode_kupon;index:idx_product_code,unique"` // Changed to composite unique index
	Discount  float64    `json:"discount" gorm:"column:diskon"`
	ValidFrom time.Time  `json:"valid_from" gorm:"column:berlaku_dari"`
	ValidTo   time.Time  `json:"valid_to" gorm:"column:berlaku_sampai"`
	IsActive  bool       `json:"is_active" gorm:"column:aktif;default:true"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// TableName specifies the table name for GORM
func (ProductCoupon) TableName() string {
	return "kupon_produk"
}
