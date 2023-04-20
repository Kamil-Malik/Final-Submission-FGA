package router

import (
	"Proyek-Akhir-Golang/control"
	"Proyek-Akhir-Golang/docs"
	"Proyek-Akhir-Golang/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title           MyGram API
// @version         1.0
// @description     FGA Digitalent Final Submission
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func StartServer() *gin.Engine {
	router := gin.Default()
	docs.SwaggerInfo.Title = "MyGram Final Submission"

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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Change Port
	router.Run(":8080")

	return router
}
