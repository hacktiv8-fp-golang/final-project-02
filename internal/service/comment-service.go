package service

import (
	"github.com/hacktiv8-fp-golang/final-project-02/internal/helper"
	"github.com/hacktiv8-fp-golang/final-project-02/internal/model"
	"github.com/hacktiv8-fp-golang/final-project-02/internal/repository"

	"github.com/asaskevich/govalidator"
)

type commentServiceRepo interface {
	CreateComment(comment *model.Comment, userID uint) (*model.Comment, helper.Error)
	GetComment(userID uint) ([]*model.Comment, helper.Error)
	UpdateComment(*model.CommentUpdate, uint) (*model.Comment, helper.Error)
	DeleteComment(uint) (*model.Comment, helper.Error)
}

type commentService struct{}

var CommentService commentServiceRepo = &commentService{}

func (t *commentService) CreateComment(comment *model.Comment, userID uint) (*model.Comment, helper.Error) {
	comment.UserID = userID

	if _, err := govalidator.ValidateStruct(comment); err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	result, err := repository.CommentModel.CreateComment(comment)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *commentService) GetComment(userID uint) ([]*model.Comment, helper.Error) {

	comment, err := repository.CommentModel.GetComment(userID)

	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (t *commentService) UpdateComment(message *model.CommentUpdate, commentID uint) (*model.Comment, helper.Error) {
	if _, err := govalidator.ValidateStruct(message); err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	result, err := repository.CommentModel.UpdateComment(message, commentID)
	
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *commentService) DeleteComment(commentID uint) (*model.Comment, helper.Error) {
	Result, err := repository.CommentModel.DeleteComment(commentID)

	if err != nil {
		return nil, err
	}

	return Result,nil
}