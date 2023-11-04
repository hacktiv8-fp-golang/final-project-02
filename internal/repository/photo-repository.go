package repository

import (
	"final-project-02/internal/database"
	"final-project-02/internal/model"
	"final-project-02/internal/utils"
)

type photoModelRepo interface {
	CreatePhoto(*model.Photo) (*model.Photo, error)
	UpdatePhoto(*model.PhotoUpdate, uint) (*model.Photo, error)
	GetAllPhotos(uint) ([]*model.Photo, utils.Error)
	DeletePhoto(uint) (error)
}

type photoModel struct{}

var PhotoModel photoModelRepo = &photoModel{}

func (t *photoModel) CreatePhoto(photo *model.Photo) (*model.Photo, error) {
	db := database.GetDB()

	err := db.Create(&photo).Error

	if err != nil {
		return nil, err
	}

	return photo, nil
}

func (t *photoModel) UpdatePhoto(input *model.PhotoUpdate, photoID uint) (*model.Photo, error) {
	db := database.GetDB()

	var photo model.Photo
	err := db.First(&photo, photoID).Error

	if err != nil {
		return nil, err
	}

	db.Model(&photo).Updates(input)

	return &photo, nil
}

func (p *photoModel) GetAllPhotos(userId uint) ([]*model.Photo, utils.Error) {
	db := database.GetDB()
	var photos []*model.Photo

	err := db.Preload("User").Where("user_id = ?", userId).Find(&photos).Error

	if err != nil {
		return nil, utils.ParseError(err)
	}

	return photos, nil
}

func (p *photoModel) DeletePhoto(photoId uint) error {
	db := database.GetDB()

	var photo model.Photo

	err := db.First(&photo, photoId).Error

	if err != nil {
		return err
	}

	err = db.Delete(&photo).Error

	if err != nil {
		return err
	}

	return nil
}