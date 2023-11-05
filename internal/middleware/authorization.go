package middleware

import (
	"final-project-02/internal/database"
	"final-project-02/internal/model"
	"final-project-02/internal/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		photoId, err := utils.GetIdParam(context, "photoId")

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
			err := utils.NotFound("Data not found")
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if photo.UserID != userID {
			err := utils.Unautorized("You are not allowed to access this data")
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		context.Next()
	}
}
