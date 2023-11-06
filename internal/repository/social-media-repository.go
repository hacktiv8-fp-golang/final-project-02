package repository

import (
	"final-project-02/internal/database"
	"final-project-02/internal/model"
	"final-project-02/internal/utils"
)

type socialMediaModelRepo interface {
	CreateSocialMedia(*model.SocialMedia) (*model.SocialMedia, utils.Error)
	GetAllSocialMedias(uint) ([]*model.SocialMedia, utils.Error)
	UpdateSocialMedia(*model.SocialMedia, uint) (*model.SocialMedia, utils.Error)
	DeleteSocialMedia(uint) utils.Error
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

func (s *socialMediaRepo) GetAllSocialMedias(userId uint) ([]*model.SocialMedia, utils.Error) {
	var socialMedia []*model.SocialMedia

	db := database.GetDB()

	err := db.Preload("User").Where("user_id", userId).Find(&socialMedia).Error

	if err != nil {
		return nil, utils.ParseError(err)
	}

	if len(socialMedia) == 0 {
		return nil, utils.NotFound("Social media data is still empty. Please add data to continue.")
	}

	return socialMedia, nil
}

func (s *socialMediaRepo) UpdateSocialMedia(socialMediaUpdated *model.SocialMedia, socialMediaId uint) (*model.SocialMedia, utils.Error) {
	var socialMedia model.SocialMedia

	db := database.GetDB()

	err := db.Model(&socialMedia).Where("id = ?", socialMediaId).Updates(socialMediaUpdated).Error

	if err != nil {
		return nil, utils.ParseError(err)
	}

	return &socialMedia, nil
}

func (s *socialMediaRepo) DeleteSocialMedia(socialMediaId uint) utils.Error {
	db := database.GetDB()
	var socialMedia model.SocialMedia
	
	err := db.Where("id = ?", socialMediaId).Delete(&socialMedia).Error

	if err != nil {
		return utils.ParseError(err)
	}

	return nil
}
