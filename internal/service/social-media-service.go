package service

import (
	"final-project-02/internal/model"
	"final-project-02/internal/utils"
	"final-project-02/internal/repository"
)

type socialMediaServiceRepo interface {
	CreateSocialMedia(*model.SocialMedia) (*model.SocialMedia, utils.Error)
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