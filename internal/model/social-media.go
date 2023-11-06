package model

import "time"

type SocialMedia struct {
	ID             uint      `json:"id,omitempty" gorm:"primaryKey" `
	Name           string    `json:"name" gorm:"not null" valid:"required~Name is required"`
	SocialMediaURL string    `json:"social_media_url" gorm:"not null" valid:"required~Social media url is required"`
	UserID         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User *User
}

type SocialMediaCreate struct {
	Name           string    `json:"name" gorm:"not null" valid:"required~Name is required"`
	SocialMediaURL string    `json:"social_media_url" gorm:"not null" valid:"required~Social media url is required"`
}

type SocialMediaUpdate struct {
	Name           string    `json:"name" gorm:"not null" valid:"required~Name is required"`
	SocialMediaURL string    `json:"social_media_url" gorm:"not null" valid:"required~Social media url is required"`
}
