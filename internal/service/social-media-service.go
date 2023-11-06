package service

import (
	"final-project-02/internal/model"
	"final-project-02/internal/repository"
	"final-project-02/internal/helper"

	"github.com/asaskevich/govalidator"
)

type socialMediaServiceRepo interface {
	CreateSocialMedia(*model.SocialMedia) (*model.SocialMedia, helper.Error)
	GetAllSocialMedias(uint) ([]*model.SocialMedia, helper.Error)
	UpdateSocialMedia(*model.SocialMediaUpdate, uint) (*model.SocialMedia, helper.Error)
	DeleteSocialMedia(uint) helper.Error
}

type socialMediaService struct{}

var SocialMediaService socialMediaServiceRepo = &socialMediaService{}

func (s *socialMediaService) CreateSocialMedia(socialMedia *model.SocialMedia) (*model.SocialMedia, helper.Error) {
	_, err := govalidator.ValidateStruct(socialMedia)

	if err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	socialMediaResponse, errorMessage := repository.SocialMediaRepo.CreateSocialMedia(socialMedia)

	if errorMessage != nil {
		return nil, errorMessage
	}

	return socialMediaResponse, nil
}

func (s *socialMediaService) GetAllSocialMedias(userId uint) ([]*model.SocialMedia, helper.Error) {
	socialMedias, err := repository.SocialMediaRepo.GetAllSocialMedias(userId)

	if err != nil {
		return nil, err
	}

	return socialMedias, nil
}

func (s *socialMediaService) UpdateSocialMedia(socialMediaUpdated *model.SocialMediaUpdate, socialMediaId uint) (*model.SocialMedia, helper.Error) {
	_, err := govalidator.ValidateStruct(socialMediaUpdated)

	if err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	socialMedia, errorMessage := repository.SocialMediaRepo.UpdateSocialMedia(socialMediaUpdated, socialMediaId)

	if errorMessage != nil {
		return nil, errorMessage
	}

	return socialMedia, nil
}

func (s *socialMediaService) DeleteSocialMedia(socialMediaId uint) helper.Error {
	err := repository.SocialMediaRepo.DeleteSocialMedia(socialMediaId)

	if err != nil {
		return err
	}

	return nil
}