package models

import "time"

type Role struct {
	ID        uint      `json:"id"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Users     []User    `json:"-"`
}
