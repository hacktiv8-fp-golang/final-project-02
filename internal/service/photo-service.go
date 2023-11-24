package service

import (
	"github.com/hacktiv8-fp-golang/final-project-02/internal/model"
	"github.com/hacktiv8-fp-golang/final-project-02/internal/repository"
	"github.com/hacktiv8-fp-golang/final-project-02/internal/helper"

	"github.com/asaskevich/govalidator"
)

type photoServiceRepo interface {
	CreatePhoto(*model.Photo, uint) (*model.Photo, helper.Error)
	UpdatePhoto(*model.PhotoUpdate, uint) (*model.Photo, helper.Error)
	GetAllPhotos(uint) ([]*model.Photo, helper.Error)
	DeletePhoto(uint) helper.Error
}

type photoService struct{}

var PhotoService photoServiceRepo = &photoService{}

func (t *photoService) CreatePhoto(photo *model.Photo, userID uint) (*model.Photo, helper.Error) {
	photo.UserID = userID

	if _, err := govalidator.ValidateStruct(photo); err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	result, err := repository.PhotoModel.CreatePhoto(photo)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *photoService) UpdatePhoto(photo *model.PhotoUpdate, photoID uint) (*model.Photo, helper.Error) {
	if _, err := govalidator.ValidateStruct(photo); err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	result, err := repository.PhotoModel.UpdatePhoto(photo, photoID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *photoService) GetAllPhotos(userId uint) ([]*model.Photo, helper.Error) {
	photos, err := repository.PhotoModel.GetAllPhotos(userId)

	if err != nil {
		return nil, err
	}

	return photos, nil
}

func (p *photoService) DeletePhoto(photoId uint) helper.Error {
	err := repository.PhotoModel.DeletePhoto(photoId)

	if err != nil {
		return err
	}

	return nil
}
