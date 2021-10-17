package models

import "time"

type Cart struct {
	ID        uint      `json:"id"`
	Amount    uint      `json:"amount"`
	UserID    uint      `json:"user_id"`
	ProductID uint      `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Product   Product   `json:"-"`
	User      User      `json:"-"`
}
