package entities

import "time"

type Trx struct {
	ID               uint        `json:"id" gorm:"primaryKey"`
	IDUser           uint        `json:"user_id" gorm:"column:id_user"` // Updated json tag
	AlamatPengiriman uint        `json:"alamat_pengiriman"`
	HargaTotal       float64     `json:"harga_total"`
	KodeInvoice      string      `json:"kode_invoice"`
	MethodBayar      string      `json:"method_bayar"`
	Address          Address     `json:"address" gorm:"foreignKey:AlamatPengiriman"`
	TrxDetail        []TrxDetail `json:"trx_detail" gorm:"foreignKey:IDTrx"`
	CreatedAt        *time.Time  `json:"created_at"`
	UpdatedAt        *time.Time  `json:"updated_at"`
}

func (Trx) TableName() string {
	return "trx"
}
