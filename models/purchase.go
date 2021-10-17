package models

import "time"

type Purchase struct {
	ID        uint      `json:"id"`
	Price     uint      `json:"price"`
	Amount    uint      `json:"amount"`
	UserID    uint      `json:"user_id"`
	ProductID uint      `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
