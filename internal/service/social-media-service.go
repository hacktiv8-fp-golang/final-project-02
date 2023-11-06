package service

import (
	"final-project-02/internal/model"
	"final-project-02/internal/repository"
	"final-project-02/internal/utils"
)

type socialMediaServiceRepo interface {
	CreateSocialMedia(*model.SocialMedia) (*model.SocialMedia, utils.Error)
	GetAllSocialMedias(uint) ([]*model.SocialMedia, utils.Error)
	UpdateSocialMedia(*model.SocialMedia, uint) (*model.SocialMedia, utils.Error)
	DeleteSocialMedia(uint) utils.Error
}

type socialMediaService struct{}

var SocialMediaService socialMediaServiceRepo = &socialMediaService{}

func (s *socialMediaService) CreateSocialMedia(socialMedia *model.SocialMedia) (*model.SocialMedia, utils.Error) {
	err := socialMedia.Validate()

	if err != nil {
		return nil, err
	}

	socialMediaResponse, err := repository.SocialMediaRepo.CreateSocialMedia(socialMedia)

	if err != nil {
		return nil, err
	}

	return socialMediaResponse, nil
}

func (s *socialMediaService) GetAllSocialMedias(userId uint) ([]*model.SocialMedia, utils.Error) {
	socialMedias, err := repository.SocialMediaRepo.GetAllSocialMedias(userId)

	if err != nil {
		return nil, err
	}

	return socialMedias, nil
}

func (s *socialMediaService) UpdateSocialMedia(socialMediaUpdated *model.SocialMedia, socialMediaId uint) (*model.SocialMedia, utils.Error) {
	err := socialMediaUpdated.Validate()

	if err != nil {
		return nil, err
	}

	socialMedia, err := repository.SocialMediaRepo.UpdateSocialMedia(socialMediaUpdated, socialMediaId)

	if err != nil {
		return nil, err
	}

	return socialMedia, nil
}

func (s *socialMediaService) DeleteSocialMedia(socialMediaId uint) utils.Error {
	err := repository.SocialMediaRepo.DeleteSocialMedia(socialMediaId)

	if err != nil {
		return err
	}

	return nil
}