package middleware

import (
	"final-project-02/internal/database"
	"final-project-02/internal/model"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		productID, err := strconv.Atoi(context.Param("productId"))
		if err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}

		userData := context.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		db := database.GetDB()
		photo := model.Photo{}

		err = db.Select("user_id").First(&photo, uint(productID)).Error
		if err != nil {
			context.JSON(http.StatusNotFound, err.Error())
			return
		}

		if photo.UserID != userID {
			context.JSON(http.StatusUnauthorized, err.Error())
			return
		}

		context.Next()
	}
}
