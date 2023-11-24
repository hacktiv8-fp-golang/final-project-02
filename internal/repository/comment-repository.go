package repository

import (
	"github.com/hacktiv8-fp-golang/final-project-02/internal/database"
	"github.com/hacktiv8-fp-golang/final-project-02/internal/helper"
	"github.com/hacktiv8-fp-golang/final-project-02/internal/model"
)

type commentModelRepo interface {
	CreateComment(comment *model.Comment) (*model.Comment, helper.Error)
	GetComment(userID uint) ([]*model.Comment, helper.Error)
	UpdateComment(*model.CommentUpdate, uint) (*model.Comment, helper.Error)
	DeleteComment(uint) (*model.Comment, helper.Error)
}

type commentModel struct{}

var CommentModel commentModelRepo = &commentModel{}

func (t *commentModel) CreateComment(comment *model.Comment) (*model.Comment, helper.Error) {
	db := database.GetDB()

	err:= db.Create(&comment).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	return comment, nil
}

func (t *commentModel) GetComment(userID uint) ([]*model.Comment, helper.Error) {
	db := database.GetDB()

	var comment []*model.Comment

	err := db.Preload("User").Preload("Photo").Where("user_id = ?", userID).Find(&comment).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	return comment, nil
}

func (t *commentModel) UpdateComment(message *model.CommentUpdate, commentID uint) (*model.Comment, helper.Error) {
	db := database.GetDB()

	var comment model.Comment
	err := db.First(&comment, commentID).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	db.Model(&comment).Updates(message)

	return &comment, nil
}

func (p *commentModel) DeleteComment(commentID uint) (*model.Comment, helper.Error) {
	db := database.GetDB()

	var comment model.Comment

	err := db.Where("id = ?", commentID).Delete(&comment).Error
	
	if err != nil {
		return nil, helper.ParseError(err)
	}

	return &comment, nil
}