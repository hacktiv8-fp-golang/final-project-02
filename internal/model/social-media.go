package model

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type SocialMedia struct {
	ID             uint      `json:"id" gorm:"primaryKey" `
	Name           string    `json:"name" gorm:"not null" valid:"required"`
	SocialMediaURL string    `json:"social_media_url" gorm:"not null" valid:"required"`
	UserID         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (socialMedia *SocialMedia) Validate() error {
	_, err := govalidator.ValidateStruct(socialMedia)

	if err != nil {
		return err
	}

	return nil
}
