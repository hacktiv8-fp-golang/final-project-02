package controller

import (
	"final-project-02/internal/model"
	"final-project-02/internal/service"
	"final-project-02/internal/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateSocialMedia(context *gin.Context) {
	var socialMedia model.SocialMedia

	if err := context.ShouldBindJSON(&socialMedia); err != nil {
		err := utils.UnprocessibleEntity("Invalid JSON body")
		context.JSON(err.Status(), err)
		return
	}

	userData := context.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	socialMedia.UserID = userId

	socialMediaResponse, err := service.SocialMediaService.CreateSocialMedia(&socialMedia)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id": socialMediaResponse.ID,
		"name": socialMediaResponse.Name,
		"social_media_url": socialMediaResponse.SocialMediaURL,
		"user_id": socialMediaResponse.UserID,
		"created_at": socialMediaResponse.CreatedAt,
	})
}