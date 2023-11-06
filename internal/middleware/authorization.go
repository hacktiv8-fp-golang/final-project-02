package middleware

import (
	"final-project-02/internal/database"
	"final-project-02/internal/model"
	"final-project-02/internal/helper"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		photoId, err := helper.GetIdParam(context, "photoId")

		if err != nil {
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		userData := context.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		db := database.GetDB()
		photo := model.Photo{}

		errMsg := db.Select("user_id").First(&photo, photoId).Error
		if errMsg != nil {
			err := helper.NotFound("Data not found")
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if photo.UserID != userID {
			err := helper.Unautorized("You are not allowed to access this data")
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		context.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		socialMediaId, err := helper.GetIdParam(context, "socialMediaId")

		if err != nil {
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		userData := context.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		db := database.GetDB()
		socialMedia := model.SocialMedia{}

		errMsg := db.Select("user_id").First(&socialMedia, socialMediaId).Error
		
		if errMsg != nil {
			err := helper.NotFound("Data not found")
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if socialMedia.UserID != userID {
			err := helper.Unautorized("You are not allowed to access this data")
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		context.Next()
	}
}
