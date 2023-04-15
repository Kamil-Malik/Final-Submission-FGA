package controller

import (
	"Proyek-Akhir-Golang/dto"
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {

	var user dto.UserDTO

	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			dto.GenericResponseDTO{
				Code:    http.StatusBadRequest,
				Message: "Please provide a valid data",
			})
		return
	}

	if _, err := valid.ValidateStruct(user); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			dto.GenericResponseDTO{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		return
	}

	if user.Age <= 8 {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			dto.GenericResponseDTO{
				Code:    http.StatusBadRequest,
				Message: "You should be at least 9 years old",
			})
		return
	}
}

func Login(ctx *gin.Context) {

}
