package models

import "time"

type Rating struct {
	ID        uint      `json:"id"`
	Rating    string    `json:"rating"`
	UserID    string    `json:"user_id"`
	ProductID string    `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
