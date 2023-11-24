package repository

import (
	"github.com/hacktiv8-fp-golang/final-project-02/internal/database"
	"github.com/hacktiv8-fp-golang/final-project-02/internal/model"
	"github.com/hacktiv8-fp-golang/final-project-02/internal/helper"
)

type socialMediaModelRepo interface {
	CreateSocialMedia(*model.SocialMedia) (*model.SocialMedia, helper.Error)
	GetAllSocialMedias(uint) ([]*model.SocialMedia, helper.Error)
	UpdateSocialMedia(*model.SocialMediaUpdate, uint) (*model.SocialMedia, helper.Error)
	DeleteSocialMedia(uint) helper.Error
}

type socialMediaRepo struct{}

var SocialMediaRepo socialMediaModelRepo = &socialMediaRepo{}

func (s *socialMediaRepo) CreateSocialMedia(socialMedia *model.SocialMedia) (*model.SocialMedia, helper.Error) {
	db := database.GetDB()

	err := db.Create(&socialMedia).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	return socialMedia, nil
}

func (s *socialMediaRepo) GetAllSocialMedias(userId uint) ([]*model.SocialMedia, helper.Error) {
	db := database.GetDB()
	var socialMedia []*model.SocialMedia

	err := db.Preload("User").Where("user_id", userId).Find(&socialMedia).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	if len(socialMedia) == 0 {
		return nil, helper.NotFound("Social media data is still empty. Please add data to continue.")
	}

	return socialMedia, nil
}

func (s *socialMediaRepo) UpdateSocialMedia(socialMediaUpdated *model.SocialMediaUpdate, socialMediaId uint) (*model.SocialMedia, helper.Error) {
	db := database.GetDB()
	var socialMedia model.SocialMedia

	err := db.First(&socialMedia, socialMediaId).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	db.Model(&socialMedia).Updates(socialMediaUpdated)

	return &socialMedia, nil
}

func (s *socialMediaRepo) DeleteSocialMedia(socialMediaId uint) helper.Error {
	db := database.GetDB()
	var socialMedia model.SocialMedia
	
	err := db.Where("id = ?", socialMediaId).Delete(&socialMedia).Error

	if err != nil {
		return helper.ParseError(err)
	}

	return nil
}
