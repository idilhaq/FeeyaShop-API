package models

import "time"

type Product struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Price     string    `json:"price"`
	Rate      uint      `json:"rate"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
