package router

import (
	"final-project-02/internal/controller"
	"final-project-02/internal/middleware"

	"github.com/gin-gonic/gin"
)

var PORT = ":8080"

func StartServer() {
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controller.Register)
		userRouter.POST("/login", controller.Login)
		userRouter.PUT("/")
		userRouter.DELETE("/")
	}

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.POST("/", controller.CreatePhoto)
		photoRouter.GET("/", controller.GetAllPhotos)
		photoRouter.PUT("/:photoId", middleware.PhotoAuthorization(), controller.UpdatePhoto)
		photoRouter.DELETE("/:photoId")
	}

	commentRouter := router.Group("/comments")
	{
		commentRouter.POST("/")
		commentRouter.GET("/")
		commentRouter.PUT("/:commentId")
		commentRouter.DELETE("/:commentId")
	}

	socialMediaRouter := router.Group("/socialmedias")
	{
		socialMediaRouter.POST("/")
		socialMediaRouter.GET("/")
		socialMediaRouter.PUT("/:socialMediaId")
		socialMediaRouter.DELETE("/:socialMediaId")
	}

	router.Run(PORT)
}
