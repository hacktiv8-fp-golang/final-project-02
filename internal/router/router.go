package router

import (
	"github.com/hacktiv8-fp-golang/final-project-02/internal/controller"
	"github.com/hacktiv8-fp-golang/final-project-02/internal/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controller.Register)
		userRouter.POST("/login", controller.Login)
		userRouter.PUT("/",middleware.Authentication(), controller.UpdateUser)
		userRouter.DELETE("/",middleware.Authentication(), controller.DeleteUser)
	}

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.POST("/", controller.CreatePhoto)
		photoRouter.GET("/", controller.GetAllPhotos)
		photoRouter.PUT("/:photoId", middleware.PhotoAuthorization(), controller.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middleware.PhotoAuthorization(), controller.DeletePhoto)
	}

	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middleware.Authentication())
		commentRouter.POST("/",controller.CreateComment)
		commentRouter.GET("/",controller.GetComment)
		commentRouter.PUT("/:commentId", middleware.CommentAuthorization(),controller.UpdateComment)
		commentRouter.DELETE("/:commentId",middleware.CommentAuthorization(),controller.DeleteComment)
	}

	socialMediaRouter := router.Group("/socialmedias")
	{
		socialMediaRouter.Use(middleware.Authentication())
		socialMediaRouter.POST("/", controller.CreateSocialMedia)
		socialMediaRouter.GET("/", controller.GetAllSocialMedias)
		socialMediaRouter.PUT("/:socialMediaId",middleware.SocialMediaAuthorization(), controller.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", middleware.SocialMediaAuthorization(), controller.DeleteSocialMedia)
	}

	var PORT = os.Getenv("PORT")

	router.Run(":" +PORT)
}
