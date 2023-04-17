package control

import (
	"Proyek-Akhir-Golang/dto"
	"Proyek-Akhir-Golang/helper"
	"Proyek-Akhir-Golang/mapper"
	"Proyek-Akhir-Golang/service"
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var user dto.UserDTO
	contentType := helper.GetContentType(ctx)
	if contentType == "application/json" {
		if err := ctx.ShouldBindJSON(&user); err != nil {
			helper.AbortGenericResponse(
				ctx,
				http.StatusBadRequest,
				"Please provide a valid data",
			)
			return
		}
	} else {
		if err := ctx.ShouldBind(&user); err != nil {
			helper.AbortGenericResponse(
				ctx,
				http.StatusBadRequest,
				"Please provide a valid data",
			)
			return
		}
	}

	if _, err := valid.ValidateStruct(user); err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	if user.Age <= 8 {
		helper.AbortGenericResponse(
			ctx,
			http.StatusBadRequest,
			"You should be at least 9 years old",
		)
		return
	}

	user.Password = helper.HashPass(user.Password)
	if err := service.InsertUser(mapper.DtoToEntity(user)); err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	helper.SendGenericResponse(ctx)
}

func Login(ctx *gin.Context) {
	var user dto.UserDTO
	contentType := helper.GetContentType(ctx)
	if contentType == "application/json" {
		if err := ctx.ShouldBindJSON(&user); err != nil {
			helper.AbortGenericResponse(
				ctx,
				http.StatusBadRequest,
				"Please provide a valid data",
			)
			return
		}
	} else {
		if err := ctx.ShouldBind(&user); err != nil {
			helper.AbortGenericResponse(
				ctx,
				http.StatusBadRequest,
				"Please provide a valid data",
			)
			return
		}
	}

	localUser, err := service.GetUserByEmail(user.Email)
	if err != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusUnauthorized,
			"Email is not registered",
		)
		return
	}

	validationError := helper.ComparePassword(user.Password, localUser.Password)
	token := helper.GenerateToken(localUser.ID, localUser.Email)
	if validationError != nil {
		helper.AbortGenericResponse(
			ctx,
			http.StatusUnauthorized,
			"Email/Password is invalid",
		)
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}
