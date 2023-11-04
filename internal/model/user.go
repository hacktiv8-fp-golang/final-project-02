package model

import (
	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Email        string    `gorm:"not null;unique" json:"email" valid:"required,email"`
	Username     string    `gorm:"not null;unique" json:"username" valid:"required"`
	Password     string    `gorm:"not null" json:"password" valid:"required,minstringlength(6)"`
	Age          int       `gorm:"not null" json:"age" valid:"required,range(8|150)"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Photos       []Photo
	Comments     []Comment
	SocialMedias []SocialMedia
}

type UserUpdate struct {
	Email    string `gorm:"not null;unique" json:"email" valid:"email"`
	Username string `gorm:"not null;unique" json:"username"`
	Age      int    `gorm:"not null" json:"age" valid:"minstringlength(8|150)"`
}

type LoginCredential struct {
	Email    string `gorm:"not null;unique" json:"email" valid:"required,email"`
	Password string `gorm:"not null" json:"password" valid:"required,minstringlength(6)"`
}
