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

	context.JSON(http.StatusCreated, gin.H{
		"id":         result.ID,
		"title":      result.Title,
		"caption":    result.Caption,
		"photo_url":  result.PhotoURL,
		"user_id":    result.UserID,
		"created_at": result.CreatedAt,
	})
}

func UpdatePhoto(context *gin.Context) {
	var photo model.PhotoUpdate

	if err := context.ShouldBindJSON(&photo); err != nil {
		context.JSON(http.StatusUnprocessableEntity, "Invalid JSON body")
		return
	}

	photoIDInt, _ := strconv.Atoi(context.Param("photoId"))
	photoID := uint(photoIDInt)

	result, err := service.PhotoService.UpdatePhoto(&photo, photoID)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":         result.ID,
		"title":      result.Title,
		"caption":    result.Caption,
		"photo_url":  result.PhotoURL,
		"user_id":    result.UserID,
		"updated_at": result.CreatedAt,
	})
}

func GetAllPhotos(context *gin.Context) {
	photos, err := service.PhotoService.GetAllPhotos()

	if err != nil {
		context.JSON(http.StatusNotFound, err.Error())
		return
	}

	var photoMaps []map[string]interface{}

	for _, photo := range photos {
		photoMap := map[string]interface{}{
			"id":         photo.ID,
			"title":      photo.Title,
			"caption":    photo.Caption,
			"photo_url":  photo.PhotoURL,
			"user_id":    photo.UserID,
			"created_at": photo.CreatedAt,
			"updated_at": photo.UpdatedAt,
			"User": 			map[string]interface{}{
				"email":    photo.User.Email,
				"username": photo.User.Username,
			},
		}

		photoMaps = append(photoMaps, photoMap)
	}

	context.JSON(http.StatusOK, photoMaps)
}

func DeletePhoto(context *gin.Context) {
	photoId, err := strconv.Atoi(context.Param("photoId"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Photo id must be an integer",})
		return
	}

	photoIdParsed := uint(photoId)

	err = service.PhotoService.DeletePhoto(photoIdParsed)

	if err != nil {
		context.JSON(http.StatusNotFound, err.Error())
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}