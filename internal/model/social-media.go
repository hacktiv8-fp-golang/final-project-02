package model

import (
	"final-project-02/internal/utils"
	"time"

	"github.com/asaskevich/govalidator"
)

type SocialMedia struct {
	ID             uint      `json:"id,omitempty" gorm:"primaryKey" `
	Name           string    `json:"name" gorm:"not null" valid:"required~Name is required"`
	SocialMediaURL string    `json:"social_media_url" gorm:"not null" valid:"required~Social media url is required"`
	UserID         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (socialMedia *SocialMedia) Validate() utils.Error {
	_, err := govalidator.ValidateStruct(socialMedia)

	if err != nil {
		return utils.BadRequest(err.Error())
	}

	return nil
}
