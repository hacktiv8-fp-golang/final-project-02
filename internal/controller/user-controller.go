package controller

import (
	"github.com/hacktiv8-fp-golang/final-project-02/internal/helper"
	"github.com/hacktiv8-fp-golang/final-project-02/internal/model"
	"github.com/hacktiv8-fp-golang/final-project-02/internal/service"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	var user model.User

	if err := context.ShouldBindJSON(&user); err != nil {
		errorHandler := helper.UnprocessibleEntity("Invalid JSON body")

		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	result, err := service.UserService.Register(&user)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"id":       result.ID,
		"username": result.Username,
		"age":      result.Age,
		"email":    result.Email,
	})
}

func Login(context *gin.Context) {
	var user model.LoginCredential

	if err := context.ShouldBindJSON(&user); err != nil {
		errorHandler := helper.UnprocessibleEntity("Invalid JSON body")

		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	result, err := service.UserService.Login(&user)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": result})
}

func UpdateUser(context *gin.Context) {
	var update model.UserUpdate

	if err := context.ShouldBindJSON(&update); err != nil {
		errorHandler := helper.UnprocessibleEntity("Invalid JSON body")

		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	userData := context.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	updatedUser, err := service.UserService.UpdateUser(userID, &update)
	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":					updatedUser.ID,
		"email":			updatedUser.Email,
		"username":		updatedUser.Username,
		"age":				updatedUser.Age,
		"updated_at":	updatedUser.UpdatedAt,
	})
}

func DeleteUser(context *gin.Context){
	userData := context.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	_, err := service.UserService.DeleteUser(userID)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message" : "Your Account has been successfully deleted",
	})
}