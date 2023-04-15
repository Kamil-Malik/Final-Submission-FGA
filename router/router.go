package router

import (
	"Proyek-Akhir-Golang/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	// User
	router.POST("/user/register", controller.Register)
	router.POST("/user/login")

	// Photo
	router.GET("/photos")
	router.GET("/photo/:id")
	router.POST("/photo")
	router.PUT("/photo")
	router.DELETE("/photo")

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
