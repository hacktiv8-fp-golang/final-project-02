package repository

import (
	"final-project-02/internal/database"
	"final-project-02/internal/helper"
	"final-project-02/internal/model"
)

type userModelRepo interface {
	Register(*model.User) (*model.User, helper.Error)
	Login(*model.LoginCredential) (*model.User, helper.Error)
	UpdateUser(userID uint, update *model.UserUpdate) (*model.User, helper.Error)
	DeleteUser(userID uint) (*model.User, helper.Error)
}

type userModel struct{}

var UserModel userModelRepo = &userModel{}

func (t *userModel) Register(user *model.User) (*model.User, helper.Error) {
	db := database.GetDB()

	err := db.Create(&user).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	return user, nil
}

func (t *userModel) Login(login *model.LoginCredential) (*model.User, helper.Error) {
	db := database.GetDB()

	var user model.User
	err := db.Where("email = ?", login.Email).First(&user).Error

	if err != nil {
		return nil, helper.Unautorized("Invalid email/password")
	}

	return &user, nil
}

func (t *userModel) UpdateUser(userID uint, update *model.UserUpdate) (*model.User, helper.Error){
	db := database.GetDB()

	var user model.User
	err := db.First(&user, userID).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	if update.Email != "" {
		user.Email = update.Email
	}
	if update.Username != "" {
		user.Username = update.Username
	}

	err = db.Save(&user).Error

  if err != nil {
		return nil, helper.ParseError(err)
	}

	return &user, nil
}

func (t *userModel) DeleteUser(userID uint) (*model.User, helper.Error) {
	db := database.GetDB()

	var user model.User
	err := db.First(&user, userID).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	db.Delete(&user)

	return &user, nil
}