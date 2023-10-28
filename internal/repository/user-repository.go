package repository

import (
	"final-project-02/internal/database"
	"final-project-02/internal/model"
)

type userModelRepo interface {
	Register(*model.User) (*model.User, error)
	Login(*model.LoginCredential) (*model.User, error)
}

type userModel struct{}

var UserModel userModelRepo = &userModel{}

func (t *userModel) Register(user *model.User) (*model.User, error) {
	db := database.GetDB()

	err := db.Create(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (t *userModel) Login(login *model.LoginCredential) (*model.User, error) {
	db := database.GetDB()

	var user model.User
	err := db.Where("username = ?", login.Username).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
