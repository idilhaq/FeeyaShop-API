package models

import "time"

type Cart struct {
	ID        uint      `json:"id"`
	Price     string    `json:"rating"`
	Amount    string    `json:"amount"`
	UserID    string    `json:"user_id"`
	ProductID string    `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}