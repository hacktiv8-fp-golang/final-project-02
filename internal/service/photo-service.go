package service

import (
	"final-project-02/internal/model"
	"final-project-02/internal/repository"

	"github.com/asaskevich/govalidator"
)

type photoServiceRepo interface {
	CreatePhoto(*model.Photo, uint) (*model.Photo, error)
	UpdatePhoto(*model.PhotoUpdate, uint) (*model.Photo, error)
	GetAllPhotos() ([]*model.Photo, error)
}

type photoService struct{}

var PhotoService photoServiceRepo = &photoService{}

func (t *photoService) CreatePhoto(photo *model.Photo, userID uint) (*model.Photo, error) {
	photo.UserID = userID

	if _, err := govalidator.ValidateStruct(photo); err != nil {
		return nil, err
	}

	result, err := repository.PhotoModel.CreatePhoto(photo)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *photoService) UpdatePhoto(photo *model.PhotoUpdate, photoID uint) (*model.Photo, error) {
	if _, err := govalidator.ValidateStruct(photo); err != nil {
		return nil, err
	}

	result, err := repository.PhotoModel.UpdatePhoto(photo, photoID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *photoService) GetAllPhotos() ([]*model.Photo, error) {
	photos, err := repository.PhotoModel.GetAllPhotos()

	if err != nil {
		return nil, err
	}

	return photos, nil
}
