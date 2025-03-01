package models

import "time"

type ProductCouponRequest struct {
	ProductID uint      `json:"product_id"`
	Code      string    `json:"code"`
	Discount  float64   `json:"discount"`
	ValidFrom time.Time `json:"valid_from"`
	ValidTo   time.Time `json:"valid_to"`
}

type ProductCouponResponse struct {
	ID        uint      `json:"id"`
	ProductID uint      `json:"product_id"`
	Code      string    `json:"code"`
	Discount  float64   `json:"discount"`
	ValidFrom time.Time `json:"valid_from"`
	ValidTo   time.Time `json:"valid_to"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
