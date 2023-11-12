package repository

import (
	"final-project-02/internal/database"
	"final-project-02/internal/model"
	"final-project-02/internal/helper"
)

type photoModelRepo interface {
	CreatePhoto(*model.Photo) (*model.Photo, helper.Error)
	UpdatePhoto(*model.PhotoUpdate, uint) (*model.Photo, helper.Error)
	GetAllPhotos(uint) ([]*model.Photo, helper.Error)
	DeletePhoto(uint) (helper.Error)
}

type photoModel struct{}

var PhotoModel photoModelRepo = &photoModel{}

func (t *photoModel) CreatePhoto(photo *model.Photo) (*model.Photo, helper.Error) {
	db := database.GetDB()

	err := db.Create(&photo).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	return photo, nil
}

func (t *photoModel) UpdatePhoto(input *model.PhotoUpdate, photoID uint) (*model.Photo, helper.Error) {
	db := database.GetDB()

	var photo model.Photo
	err := db.First(&photo, photoID).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	db.Model(&photo).Updates(input)

	return &photo, nil
}

func (p *photoModel) GetAllPhotos(userId uint) ([]*model.Photo, helper.Error) {
	db := database.GetDB()
	var photos []*model.Photo

	err := db.Preload("User").Where("user_id = ?", userId).Find(&photos).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	return photos, nil
}

func (p *photoModel) DeletePhoto(photoId uint) helper.Error {
	db := database.GetDB()

	var photo model.Photo

	err := db.Where("id = ?", photoId).Delete(&photo).Error
	
	if err != nil {
		return helper.ParseError(err)
	}

	return nil
}