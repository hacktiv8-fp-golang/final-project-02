package controller

import (
	"github.com/hacktiv8-fp-golang/final-project-02/internal/helper"
	"github.com/hacktiv8-fp-golang/final-project-02/internal/model"
	"github.com/hacktiv8-fp-golang/final-project-02/internal/service"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateComment(context *gin.Context) {
	var comment model.Comment

	if err := context.ShouldBindJSON(&comment); err != nil {
		errorHandler := helper.UnprocessibleEntity("Invalid JSON body")
		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	userData := context.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	createComment, err := service.CommentService.CreateComment(&comment, userID)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"id":         createComment.ID,
		"message":    createComment.Message,
		"photo_id":   createComment.PhotoID,
		"user_id":    createComment.UserID,
		"created_at": createComment.CreatedAt,
	})
}

func GetComment(context *gin.Context) {
	userData := context.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	results, err := service.CommentService.GetComment(userId)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	comments := make([]gin.H, 0, len(results))

	for _, result := range results {
		comment := gin.H{
			"id":         result.ID,
			"message":    result.Message,
			"photo_id":   result.PhotoID,
			"user_id":    result.UserID,
			"updated_at": result.UpdatedAt,
			"created_at": result.CreatedAt,
			"User": gin.H{
				"id":       result.User.ID,
				"email":    result.User.Email,
				"username": result.User.Username,
			},
			"Photo": gin.H{
				"id":        result.Photo.ID,
				"title":     result.Photo.Title,
				"caption":   result.Photo.Caption,
				"photo_url": result.Photo.PhotoURL,
				"user_id":   result.Photo.UserID,
			},
		}

		comments = append(comments, comment)
	}

	context.JSON(http.StatusOK, comments)
}

func UpdateComment(context *gin.Context) {
	var message model.CommentUpdate

	if err := context.ShouldBindJSON(&message); err != nil {
		errorHandler := helper.UnprocessibleEntity("Invalid JSON body")
		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	ID, _ := helper.GetIdParam(context, "commentId")

	commentID := uint(ID)

	Update, err := service.CommentService.UpdateComment(&message, commentID)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":         Update.ID,
		"message":    Update.Message,
		"user_id":    Update.UserID,
		"updated_at": Update.CreatedAt,
	})
}

func DeleteComment(context *gin.Context) {
	ID, _ := helper.GetIdParam(context, "commentId")
	commentID := uint(ID)

	_, err := service.CommentService.DeleteComment(commentID)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
