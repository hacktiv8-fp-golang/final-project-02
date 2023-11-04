package repository

import (
	"final-project-02/internal/database"
	"final-project-02/internal/model"
)

type photoModelRepo interface {
	CreatePhoto(*model.Photo) (*model.Photo, error)
	UpdatePhoto(*model.PhotoUpdate, uint) (*model.Photo, error)
	GetAllPhotos() ([]*model.Photo, error)
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

func (p *photoModel) GetAllPhotos() ([]*model.Photo, error) {
	db := database.GetDB()
	var photos []*model.Photo

	err := db.Preload("User").Find(&photos).Error

	if err != nil {
		return nil, err
	}

	return photos, nil
}