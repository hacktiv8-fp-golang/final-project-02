package service

import (
	"final-project-02/internal/helper"
	"final-project-02/internal/model"
	"final-project-02/internal/repository"

	"github.com/asaskevich/govalidator"
)

type userServiceRepo interface {
	Register(*model.User) (*model.User, error)
	Login(*model.LoginCredential) (string, error)
}

type userService struct{}

var UserService userServiceRepo = &userService{}

func (t *userService) Register(user *model.User) (*model.User, error) {
	if _, err := govalidator.ValidateStruct(user); err != nil {
		return nil, err
	}

	var err error
	user.Password, err = helper.HashPass(user.Password)
	if err != nil {
		return nil, err
	}

	result, err := repository.UserModel.Register(user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *userService) Login(login *model.LoginCredential) (string, error) {
	if _, err := govalidator.ValidateStruct(login); err != nil {
		return "", err
	}

	user, err := repository.UserModel.Login(login)
	if err != nil {
		return "", err
	}

	if isPasswordCorrect := helper.ComparePass(user.Password, login.Password); !isPasswordCorrect {
		return "", err
	}

	token, err := helper.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
