package entities

import "time"

type Notification struct {
	ID        uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Pesan     string     `json:"pesan" gorm:"type:text;not null"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (Notification) TableName() string {
	return "notifikasi"
}
