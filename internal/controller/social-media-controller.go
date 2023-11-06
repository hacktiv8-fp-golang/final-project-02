package controller

import (
	"final-project-02/internal/model"
	"final-project-02/internal/service"
	"final-project-02/internal/helper"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateSocialMedia(context *gin.Context) {
	var socialMedia model.SocialMedia

	if err := context.ShouldBindJSON(&socialMedia); err != nil {
		err := helper.UnprocessibleEntity("Invalid JSON body")
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

func GetAllSocialMedias(context *gin.Context) {
	userData := context.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	socialMedias, err := service.SocialMediaService.GetAllSocialMedias(userId)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	var socialMediaMaps []map[string]interface{}

	for _, socialMedia := range socialMedias {
		socialMediaMap := map[string]interface{}{
			"id": socialMedia.ID,
			"name": socialMedia.Name,
			"social_media_url": socialMedia.SocialMediaURL,
			"user_id": socialMedia.UserID,
			"created_at": socialMedia.CreatedAt,
			"updated_at": socialMedia.UpdatedAt,
			"user": map[string]interface{}{
				"id": socialMedia.User.ID,
				"email": socialMedia.User.Email,
				"username": socialMedia.User.Username,
			},
		}

		socialMediaMaps = append(socialMediaMaps, socialMediaMap)
	}

	context.JSON(http.StatusOK, gin.H{
		"social_medias": socialMediaMaps,
	})
}

func UpdateSocialMedia(context *gin.Context) {
	id, _ := helper.GetIdParam(context, "socialMediaId")

	var socialMediaUpdated model.SocialMediaUpdate

	if err := context.ShouldBindJSON(&socialMediaUpdated); err != nil {
		err := helper.UnprocessibleEntity("Invalid JSON body")
		context.JSON(err.Status(), err)
		return
	}

	socialMedia, err := service.SocialMediaService.UpdateSocialMedia(&socialMediaUpdated, id)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id": socialMedia.ID,
		"name": socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaURL,
		"user_id": socialMedia.UserID,
		"updated_at": socialMedia.UpdatedAt,
	})
}

func DeleteSocialMedia(context *gin.Context) {
	socialMediaId, _ := helper.GetIdParam(context, "socialMediaId")

	err := service.SocialMediaService.DeleteSocialMedia(socialMediaId)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}