package models

import "time"

type Capsule struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Message   string    `gorm:"not null" json:"message"`
	DeliverAt time.Time `gorm:"not null" json:"deliver_at"`
	Delivered bool      `gorm:"default:false" json:"delivered"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
