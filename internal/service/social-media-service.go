package service

import (
	"github.com/hacktiv8-fp-golang/final-project-02/internal/model"
	"github.com/hacktiv8-fp-golang/final-project-02/internal/repository"
	"github.com/hacktiv8-fp-golang/final-project-02/internal/helper"

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
	if _, err := govalidator.ValidateStruct(socialMedia); err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	socialMediaResponse, err := repository.SocialMediaRepo.CreateSocialMedia(socialMedia)

	if err != nil {
		return nil, err
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
	if _, err := govalidator.ValidateStruct(socialMediaUpdated); err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	socialMedia, err := repository.SocialMediaRepo.UpdateSocialMedia(socialMediaUpdated, socialMediaId)

	if err != nil {
		return nil, err
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