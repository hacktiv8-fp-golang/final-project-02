package controller

import (
	"final-project-02/internal/model"
	"final-project-02/internal/service"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreatePhoto(context *gin.Context) {
	var photo model.Photo

	if err := context.ShouldBindJSON(&photo); err != nil {
		context.JSON(http.StatusUnprocessableEntity, "Invalid JSON body")
		return
	}

	userData := context.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	result, err := service.PhotoService.CreatePhoto(&photo, userID)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusCreated, result)
}

func UpdatePhoto(context *gin.Context) {
	var photo model.PhotoUpdate

	if err := context.ShouldBindJSON(&photo); err != nil {
		context.JSON(http.StatusUnprocessableEntity, "Invalid JSON body")
		return
	}

	productIDInt, _ := strconv.Atoi(context.Param("productId"))
	productID := uint(productIDInt)

	result, err := service.PhotoService.UpdatePhoto(&photo, productID)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, result)
}