package model

import (
	"time"
)

type Photo struct {
	ID        uint      `json:"id" gorm:"primaryKey" `
	Title     string    `json:"title" gorm:"not null" valid:"required"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url" gorm:"not null" valid:"required"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      *User
	Comments  []Comment
}

type PhotoUpdate struct {
	Title    string `json:"title" gorm:"not null" valid:"required"`
	Caption  string `json:"caption" valid:"required"`
	PhotoURL string `json:"photo_url" gorm:"not null" valid:"required"`
}
