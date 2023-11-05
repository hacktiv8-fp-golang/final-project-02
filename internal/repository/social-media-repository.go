package repository

import (
	"final-project-02/internal/database"
	"final-project-02/internal/model"
	"final-project-02/internal/utils"
)

type socialMediaModelRepo interface {
	CreateSocialMedia(*model.SocialMedia) (*model.SocialMedia, utils.Error)
}

type socialMediaRepo struct{}

var SocialMediaRepo socialMediaModelRepo = &socialMediaRepo{}

func (s *socialMediaRepo) CreateSocialMedia(socialMedia *model.SocialMedia) (*model.SocialMedia, utils.Error) {
	db := database.GetDB()

	err := db.Create(&socialMedia).Error

	if err != nil {
		return nil, utils.ParseError(err)
	}

	return socialMedia, nil
}
