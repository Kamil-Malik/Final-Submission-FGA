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
		photo.PUT("/")
		photo.DELETE("/photo")
	}

	// Comment
	router.GET("/comments")
	router.GET("/comment/:id")
	router.POST("/comment")
	router.PUT("/comment")
	router.DELETE("/comment")

	// Social Media
	router.GET("/sosmed")
	router.GET("/sosmed/:id")
	router.POST("/sosmed")
	router.PUT("/sosmed")
	router.DELETE("/sosmed")

	// Change Port
	router.Run(":8080")

	return router
}
