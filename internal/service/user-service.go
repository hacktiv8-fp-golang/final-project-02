package service

import (
	"final-project-02/internal/helper"
	"final-project-02/internal/model"
	"final-project-02/internal/repository"

	"github.com/asaskevich/govalidator"
)

type userServiceRepo interface {
	Register(*model.User) (*model.User, helper.Error)
	Login(*model.LoginCredential) (string, helper.Error)
	UpdateUser(userID uint, update *model.UserUpdate) (*model.User, helper.Error)
	DeleteUser(userID uint) (*model.User, helper.Error)
}

type userService struct{}

var UserService userServiceRepo = &userService{}

func (t *userService) Register(user *model.User) (*model.User, helper.Error) {
	if _, err := govalidator.ValidateStruct(user); err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	password, err := helper.HashPass(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = password

	result, err := repository.UserModel.Register(user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *userService) Login(login *model.LoginCredential) (string, helper.Error) {
	if _, err := govalidator.ValidateStruct(login); err != nil {
		return "", helper.BadRequest(err.Error())
	}

	user, err := repository.UserModel.Login(login)
	if err != nil {
		return "", err
	}

	if isPasswordCorrect := helper.ComparePass(user.Password, login.Password); !isPasswordCorrect {
		return "", helper.Unautorized("Invalid email/password")
	}

	token, err := helper.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (t *userService) UpdateUser(userID uint, update *model.UserUpdate) (*model.User, helper.Error) {
	if _, err := govalidator.ValidateStruct(update); err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	updatedUser, err := repository.UserModel.UpdateUser(userID, update)

	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (t *userService) DeleteUser(userID uint) (*model.User, helper.Error) {
	Result, err := repository.UserModel.DeleteUser(userID)

	if err != nil {
		return nil, err
	}

	return Result, nil
}

