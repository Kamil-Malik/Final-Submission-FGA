package router

import (
	"Proyek-Akhir-Golang/control"
	"Proyek-Akhir-Golang/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	// User
	user := router.Group("/user")
	{
		user.POST("/register", control.Register)
		user.POST("/login", control.Login)
	}

	// Photo
	photo := router.Group("/photo")
	{
		photo.Use(middleware.Authentication())
		photo.GET("/", control.GetAllPhoto)
		photo.GET("/:id", control.GetPhotoByID)
		photo.POST("/", control.CreatePhoto)
		photo.PUT("/:id", control.UpdatePhotoByID)
		photo.DELETE("/:id", control.DeletePhotoByID)
	}

	// Comment
	comment := router.Group("/comment")
	{
		comment.Use(middleware.Authentication())
		comment.GET("/", control.GetAllComments)
		comment.GET("/:id", control.GetCommentByID)
		comment.POST("/", control.CreateComment)
		comment.PUT("/:id", control.UpdateCommentByID)
		comment.DELETE("/:id", control.DeleteCommentByID)
	}

	// Social Media
	socialMedia := router.Group("/social")
	{
		socialMedia.Use(middleware.Authentication())
		socialMedia.GET("/", control.GetAllSocialMedia)
		socialMedia.GET("/:id", control.GetSocialMediaByID)
		socialMedia.POST("/", control.CreateSocialMedia)
		socialMedia.PUT("/:id", control.UpdateSocialMediaByID)
		socialMedia.DELETE("/:id", control.DeleteSocialMediaByID)
	}

	// Change Port
	router.Run(":8080")

	return router
}
