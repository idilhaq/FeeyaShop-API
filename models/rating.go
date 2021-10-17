package models

import "time"

type Rating struct {
	ID        uint      `json:"id"`
	Rating    uint      `json:"rating"`
	UserID    uint      `json:"user_id"`
	ProductID uint      `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Product   Product   `json:"-"`
	User      User      `json:"-"`
}
